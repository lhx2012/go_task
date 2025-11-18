package Task3

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	gorm.Model
	Name          string
	Balance       float64       `gorm:"not null;default:0"`
	FTransactions []Transaction `gorm:"foreignKey:FromAccountID"`
	TTransactions []Transaction `gorm:"foreignKey:ToAccountID"`
}

type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func BankRun(db *gorm.DB) {
	//新建表
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	//
	//db.Create(&[]Account{Account{
	//	Balance: 3000,
	//	Name:    "张三",
	//}, Account{
	//	Name:    "李四",
	//	Balance: 0,
	//}})

	transferMoney(db, 1, 2, 100)
}

func transferMoney(db *gorm.DB, fromAccountID uint, toAccountID uint, amount float64) {
	err := db.Transaction(func(tx *gorm.DB) error {

		var account Account
		//查询账户是否存在并且添加行级锁防止并发问题
		if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).First(&account, "id = ?", fromAccountID).Error; err != nil {
			return fmt.Errorf("查询转出账户失败: %v", err)
		}

		//检查账户余额是否满足
		if account.Balance < amount {
			return fmt.Errorf("账户余额不足，当前余额：%.2f", account.Balance)
		}

		//减去form账号金额
		if err := tx.Debug().Model(&account).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return fmt.Errorf("扣款失败：%v", err)
		}

		// 更新转入账户余额
		var toAccount Account
		if err := tx.First(&toAccount, toAccountID).Error; err != nil {
			return fmt.Errorf("查询转入账户失败: %v", err)
		}

		if err := tx.Model(&toAccount).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return fmt.Errorf("存款失败: %v", err)
		}

		transaction := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("记录交易失败: %v", err)
		}

		//所有操作完成，提交事务
		return nil
	})
	if err != nil {
		fmt.Println("事务执行失败，回滚", err)
	} else {
		fmt.Println("事务执行成功，成功提交")
	}
}

// TransferMoneyManual 手动控制事务的转账实现
func TransferMoneyManual(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 使用defer确保在异常情况下回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询转出账户余额并加锁
	var fromAccount Account
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %v", err)
	}

	// 检查余额是否足够
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足，当前余额: %.2f", fromAccount.Balance)
	}

	// 执行转账操作
	if err := tx.Model(&fromAccount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣款失败: %v", err)
	}

	// 更新转入账户余额
	if err := tx.Model(&Account{}).Where("id = ?", toAccountID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("存款失败: %v", err)
	}

	// 创建交易记录
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}

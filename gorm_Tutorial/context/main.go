package main

func main() {
	//单会话模式
	//db.WithContect(ctx).Find(&users)
	//连续会话模式
	//tx := db.WithContext(ctx)
	//tx.First(&user, 1)
	//tx.Model(&user).Uadate("role", "admin")
	//上下文超时设置
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//db.WithContext(ctx).Find(&users)

}

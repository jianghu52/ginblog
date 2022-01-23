# ginblog
a go project use gin and vue

一个使用go做的blog系统，前端用vue。
整个学习参照了 https://space.bilibili.com/402177130?spm_id_from=333.788.b_765f7570696e666f.2

## 第一课
````go
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		})
	}
````
注意这里的大括号，是一定要写在第二行的，原因比较复杂。简单来说，就是第一行的内容已经是一个完整的表达式了。可以用；结尾。因此，这就不能再写大括号了。但是让我不理解的是，如果是这样，这个大括号的意义何在呢？




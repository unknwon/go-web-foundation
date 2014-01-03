《Go Web基础》
=================
感谢所有在第一套教程 [《Go编程基础》](https://github.com/Unknwon/go-fundamental-programming) 录制期间给予我大力支持的 Go 语言爱好者们，是你们的鼓励让我坚持完成这项开源事业。这套教程是后续教程，即建立在第一套基础之上的教程，已经讲解过的相关知识点将不会再赘述。因此如果您还没有完成学习第一套教程，请不要操之过急，以免事倍功半。与第一套不同的是，本套教程将以搭建个人博客作为实战目标，由浅至深地讲解使用 Go 开发 Web 应用程序的必备知识与技巧。

### 基本信息

- 教程讲师：[无闻](http://about.me/unknwon)
- 教程简介：《Go Web基础》是一套针对 Google 出品的Go语言的视频语音教程，主要面向完成《Go编程基础》教程后希望了解有关 Go Web 开发的学习者。
- Go 语言版本：1.1.1-1.2
- 教程涉及第三方库：[beego](https://github.com/astaxie/beego)（v0.9.+ - 1.0.+）、[go-sqlite3](https://github.com/mattn/go-sqlite3)、[com](https://github.com/Unknwon/com)、[i18n](https://github.com/beego/i18n)
- 教程开发环境：Windows 7 64 位（1-8 课）、Mac OS X 10.9（9-13 课）
- 其它说明：每堂课都会建立一个文件夹（例如：lecture1），内含与课程进度相符的项目源码与课堂笔记。课堂笔记中里面包含了该堂课所涵盖的知识点以及知识点开始讲解的时间点，方便学习者快速定位要了解的部分，节省不必要浪费的时间。此外，如果教程中因口误或其它原因使学习者产生迷惑的部分，同样会在课堂笔记中进行补充说明。
- 收录网站：
	-  [优才网](http://www.ucai.cn/course/show/87) 
	-  [网易云课堂](http://study.163.com/course/courseMain.htm?courseId=328001#/courseMain)
- 全套视频下载地址：[百度网盘](http://pan.baidu.com/share/link?shareid=136613208&uk=822891499)

### 教程大纲

<table class="table table-condensed table-bordered">
	<tbody>
		<tr>
			<th>课时</th>
			<th>标题</th>
			<th>在线观看</th>
		</tr>
		<tr>
			<td>第 1 课</td>
			<td><a href="lectures/lecture1/lecture1.md">博客项目设计</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/gXZb9tGNsGU/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/4710">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=442046&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 2 课</td>
			<td><a href="lectures/lecture2/lecture2.md">初窥 Web 开发</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/sqZoUrqNJno/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/4732">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=442047&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 3 课</td>
			<td><a href="lectures/lecture3/lecture3.md">模板用法讲解</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/BuoN93Yplow/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/4792">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=468064&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 4 课</td>
			<td><a href="lectures/lecture4/lecture4.md">登录及分类管理</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/UoJ9EgyqqbY/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/4793">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=476057&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 5 课</td>
			<td><a href="lectures/lecture5/lecture5.md">文章的添加与删除</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/g9q30NSRI8c/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/4800">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=502075&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 6 课</td>
			<td><a href="lectures/lecture6/lecture6.md">评论与分类显示</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/JFL7PGjpz4Q/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/5967">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=548094&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 7 课</td>
			<td><a href="lectures/lecture7/lecture7.md">为文章增加标签</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/QpE6LM3Ie2k/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/6400">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=626001&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 8 课</td>
			<td><a href="lectures/lecture8/lecture8.md">文章附件上传</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/UqVp_KqSc_A/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/6401">优才网</a>
				<a href="http://study.163.com/course/courseLearn.htm?courseId=328001#/learn/video?lessonId=626002&courseId=328001">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 9 课</td>
			<td><a href="lectures/lecture9/lecture9.md">国际化支持</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/Mic4x6lwNzo/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/69/3267/6814">优才网</a>
				<a1 href="">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 10 课</td>
			<td><a href="lectures/lecture10/lecture10.md">自建 HTTP 中间件</a></td>
			<td>
				<a href="http://www.tudou.com/programs/view/zxRhEOPz7BI/">土豆网</a>
				<a href="http://www.ucai.cn/course/chapter/87/3267/6815">优才网</a>
				<a1 href="">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 11 课</td>
			<td><a1 href="lectures/lecture11/lecture11.md">简易的 RPC 实现</a></td>
			<td>
				<a1 href="">土豆网</a>
				<a1 href="">优才网</a>
				<a1 href="">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 12 课</td>
			<td><a1 href="lectures/lecture12/lecture12.md">搭建 REST 服务</a></td>
			<td>
				<a1 href="">土豆网</a>
				<a1 href="">优才网</a>
				<a1 href="">网易云课堂</a>
			</td>
		</tr>
		<tr>
			<td>第 13 课</td>
			<td><a1 href="lectures/lecture13/lecture13.md">网站安全浅析</a></td>
			<td>
				<a1 href="">土豆网</a>
				<a1 href="">优才网</a>
				<a1 href="">网易云课堂</a>
			</td>
		</tr>
	</tbody>
</table>

### 相关链接

- Go Web 编程交流 QQ 群：259316004
- [《Go编程基础》](https://github.com/Unknwon/go-fundamental-programming)
- [Golang中文社区](http://bbs.gocn.im/forum.php)
- [Go语言学习园地](http://studygolang.com/)
- [Golang中国](http://golangtc.com/)

### 授权许可

除特别声明外，本套教程中的内容使用CC BY-SA 3.0 License（创作共用 署名-相同方式共享 3.0 许可协议）授权，代码遵循 BSD 3-Clause License（3 项条款的 BSD 许可协议）。

### 参考资料

- [《Go Web编程》](https://github.com/astaxie/build-web-application-with-golang)（[Asta谢](https://github.com/astaxie)）
- [《Go 学习笔记》](http://bbs.gocn.im/thread-8-1-1.html)（雨痕）
- [速动中国](https://github.com/insionng/toropress)（正雄）

### 捐助作者

如果您觉得本套教程确实不错，并认为作者的努力值得肯定，可以通过 [此链接](https://me.alipay.com/obahua) 为作者提供小额捐助，以资鼓励。

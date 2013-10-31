{{define "navbar"}}
	<nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
		<div class="container">
			<div class="row">
				<div class="col-md-6">
					<ul class="nav nav-pills" style="padding-top:5px">
						<li><a href="/">首页</a></li>
						<li {{if .IsBlog}}class="active"{{end}}><a href="/admin?ob=blog">管理博客</a></li>
						<li {{if .IsPhoto}}class="active"{{end}}><a href="/admin?ob=photo">管理照片</a></li>
						<li {{if .IsSort}}class="active"{{end}}><a href="/admin?ob=sort">分类管理</a></li>
					</ul>
				</div>
				<div class="col-md-6 pull-right">
					<div class="pull-right">
					<ul class="nav nav-pills" style="padding-top:5px">
						<li {{if .IsPwd}}class="active"{{end}}><a href="/admin?ob=pwd">修改密码</a></li>
					</ul>
					</div>
				</div>
			</div>
		</div>
	</nav>
{{end}}
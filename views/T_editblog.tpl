{{define "editblog"}}
	<div class="controls">
		<form method="POST" action="/editarticle">
			<p>
			<div class="row">
				<select name="ArticleSort" class="form-control">
					<option value ="">请选择文章分类</option>
					{{range .ArticleSort}}
					<option value ="{{.Id}}">{{.Title}}</option>
					{{end}}
				</select>
			</div>
			</p>
			<p><div class="row"><input class="form-control" type="text" placeholder="输入博客名称" name="ArticleTitle"></div></p>
			<p>
			<div class="row">
			<textarea class="xheditor-mini" name="ArticleContent" style="width:100%" rows="15"></textarea>
			</div>
			</p>
			<p>
				<button type="submit" class="btn btn-primary">提交</button>
				<button type="button" class="btn btn-default">重置</button>
			</p>
		</form>
		<table class="table table-striped">
	<thead>
		<tr>
			<td>#</td><td>标题</td><td>创建时间</td><td>修改时间</td><td>浏览</td><td>操作</td>
		</tr>
	</thead>
	<tbody>
		{{range .Article}}
		<tr>
			<td>{{.Id}}</td>
			<td><a href="/">{{.Title}}</a></td>
			<td>{{.Created}}</td>
			<td>{{.Updated}}</td>
			<td>{{.Views}}</td>
			<td>
				<a href="/editphoto?op=del&id={{.Id}}">删除</a>
			</td>
		</tr>
		{{end}}
	</tbody>
	</table>
	</div>
{{end}}
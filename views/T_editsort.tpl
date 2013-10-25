{{define "editsort"}}
<div class="row">
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/editblog">
				<p><input class="form-control" type="text" placeholder="输入博客分类分类" name="Sort"></p>
				<p>
					<button type="submit" class="btn btn-primary">提交</button>
					<button type="button" class="btn btn-default">重置</button>
				</p>
			</form>

		</div>
	</div>
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/editphoto">
				<p><input class="form-control" type="text" placeholder="输入图片分类" name="Sort"></p>
				<p>
					<button type="submit" class="btn btn-primary">提交</button>
					<button type="button" class="btn btn-default">重置</button>
				</p>
			</form>

			<table class="table table-striped">
			<thead>
				<tr>
					<td>#</td><td>名称</td><td>数量</td><td>操作</td>
				</tr>
			</thead>
			<tbody>
				{{range .Photosort}}
				<tr>
					<td>{{.Id}}</td>
					<td>{{.Title}}</td>
					<td>{{.Topiccount}}</td>
					<td>
						<a href="/editphoto?op=del">删除</a>
					</td>
				</tr>
				{{end}}
			</tbody>
			</table>

		</div>
	</div>
</div>
{{end}}
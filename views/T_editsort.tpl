{{define "editsort"}}
<div class="row">
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/admin/AddArticleSort">
				<p><input class="form-control" type="text" placeholder="输入博客分类分类" name="articlesorttitle"></p>
				<p>
					<button type="submit" class="btn btn-primary">提交</button>
					<button type="button" class="btn btn-default">重置</button>
				</p>
			</form>
			<table class="table table-striped">
			<thead>
				<tr>
					<td>#</td><td>名称</td><td>文章数量</td><td>操作</td>
				</tr>
			</thead>
			<tbody>
				{{range .ArticleSorts}}
				<tr>
					<td>{{.Id}}</td>
					<td>{{.Title}}</td>
					<td>{{.Count}}</td>
					<td>
						<a href="/admin?ob=blog&op=del&id={{.Id}}">删除</a>
					</td>
				</tr>
				{{end}}
			</tbody>
			</table>
		</div>
	</div>
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/admin/AddPhotoSort">
				<p><input class="form-control" type="text" placeholder="输入图片分类" name="name"></p>
				<p>
					<button type="submit" class="btn btn-primary">提交</button>
					<button type="button" class="btn btn-default">重置</button>
				</p>
			</form>

			<table class="table table-striped">
			<thead>
				<tr>
					<td>#</td><td>名称</td><td>图片数量</td><td>操作</td>
				</tr>
			</thead>
			<tbody>
				{{range .PhotoSorts}}
				<tr>
					<td>{{.Id}}</td>
					<td>{{.Title}}</td>
					<td>{{.Count}}</td>
					<td>
						<a href="/editphotosort?op=del&id={{.Id}}">删除</a>
					</td>
				</tr>
				{{end}}
			</tbody>
			</table>

		</div>
	</div>
</div>
{{end}}
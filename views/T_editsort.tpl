{{define "editsort"}}
<div class="row">
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/admin/ActionAddArticleSort">
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
					<td><a href="/admin?ob=sort&op=updateas&id={{.Id}}&title={{.Title}}">{{.Title}}</a></td>
					<td>{{.Count}}</td>
					<td>
						<a href="/admin?ob=sort&op=delarticle&id={{.Id}}">删除</a>
					</td>
				</tr>
				{{end}}
			</tbody>
			</table>
		</div>
	</div>
	<div class="col-md-6">
		<div class="controls">
			<form method="POST" action="/admin/ActionAddPhotoSort">
				<p><input class="form-control" type="text" placeholder="输入图片分类" name="photosorttitle"></p>
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
					<td><a href="/admin?ob=sort&op=updateps&id={{.Id}}&title={{.Title}}">{{.Title}}</a></td>
					<td>{{.Count}}</td>
					<td>
						<a href="/admin?ob=sort&op=delphoto&id={{.Id}}">删除</a>
					</td>
				</tr>
				{{end}}
			</tbody>
			</table>

		</div>
	</div>
</div>
{{end}}
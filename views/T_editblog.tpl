{{define "editarticle"}}
	<div class="controls">
		<div class="row">
			<div class="col-md-3">
				<button type="button" class="btn btn-primary" onClick="location.href='admin?ob=blog&op=add'">添加文章</button>
			</div>
			<div class="col-md-9">
				<div class="row">
					<form class="form-inline" role="form">
						<div class="col-md-4">
							<select name="ArticleSortId" class="form-control">
								<option value ="">请选择文章分类</option>
									{{range .ArticleSorts}}
										<option value ="{{.Id}}">{{.Title}}</option>
									{{end}}
							</select>
						</div>
						<div class="col-md-4">
							<label class="sr-only">博客名称</label>
							<input class="form-control" placeholder="博客名称">
						</div>
						<div class="col-md-4">
							<button type="submit" class="btn btn-default">查询</button>
						</div>
					</form>
				</div>
			</div>
		</div>
		<hr />		
		<table class="table table-striped">
	<thead>
		<tr>
			<td>#</td><td>标题</td><td>创建时间</td><td>最后更新时间</td><td>分类</td><td>浏览</td><td>操作</td>
		</tr>
	</thead>
	<tbody>
		{{range .Article}}
		<tr>
			<td>{{.Id}}</td>
			<td><a href="/admin?ob=blog&op=update&id={{.Id}}">{{.Title}}</a></td>
			<td>{{.Created}}</td>
			<td>{{.Updated}}</td>
			<td></td>
			<td>{{.Views}}</td>
			<td>
				<a href="/admin?ob=blog&op=del&id={{.Id}}">删除</a>
			</td>
		</tr>
		{{end}}
	</tbody>
	</table>
	</div>
{{end}}
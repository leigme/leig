{{define "editphoto"}}
<div class="controls">
	<form method="POST" action="/admin/ActionAddPhoto">
		<div class="container">
			<div class="row">
				<div class="col-md-2">
				</div>
				<div class="col-md-8">
					<div class="row">
						<div class="col-md-6">
							<input class="form-control" type="text" placeholder="输入图片名称" name="phototitle">
						</div>
						<div class="col-md-6">
							<select name="photosortid" class="form-control">
								<option value ="">请选择图片分类</option>
								{{range .PhotoSort}}
								<option value ="{{.Id}}">{{.Title}}</option>
								{{end}}
							</select>
						</div>
					</div>
					<p />
					<div class="row">
						<div class="col-md-12">
							<input class="form-control" type="text" placeholder="输入外链地址" name="photourl">
						</div>
					</div>
					<p />
					<div class="row">
						<div class="col-md-12">
							<button type="submit" class="btn btn-primary">提交</button>
							<button type="button" class="btn btn-default">重置</button>
						</div>
					</div>
				</div>
				<div class="col-md-2">
				</div>
			</div>
	</div>
	</form>
	<table class="table table-striped">
	<thead>
		<tr>
			<td>#</td><td>标题</td><td>浏览</td><td>地址</td><td>操作</td>
		</tr>
	</thead>
	<tbody>
		{{range .Photo}}
		<tr>
			<td>{{.Id}}</td>
			<td><a href="#" >{{.Title}}</a></td>
			<td>{{.Views}}</td>
			<td>{{.Url}}</td>
			<td>
				<a href="/admin?ob=photo&op=del&id={{.Id}}">删除</a>
			</td>
		</tr>
		{{end}}
	</tbody>
	</table>
</div>
{{end}}
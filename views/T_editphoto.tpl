{{define "editphoto"}}
<div class="controls">
	<form method="POST" action="/editphoto">
		<p><input class="form-control" type="text" placeholder="输入图片名称" name="PhotoTitle"></p>
		<p><input class="form-control" type="text" placeholder="输入外链地址" name="PhotoUrl"></p>
		<p>
			<select name="PhotoSortId" class="form-control">
				<option value ="">请选择图片分类</option>
				{{range .PhotoSort}}
				<option value ="{{.Id}}">{{.Title}}</option>
				{{end}}
			</select>
		</p>
		<p>
			<button type="submit" class="btn btn-primary">提交</button>
			<button type="button" class="btn btn-default">重置</button>
		</p>
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
			<td>{{.Title}}</td>
			<td>{{.Views}}</td>
			<td>{{.Url}}</td>
			<td>
				<a href="/editphoto?op=del&id={{.Id}}">删除</a>
			</td>
		</tr>
		{{end}}
	</tbody>
	</table>
</div>
{{end}}
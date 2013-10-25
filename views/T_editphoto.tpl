{{define "editphoto"}}
<div class="controls">
	<form method="POST" action="/addphoto">
		<p><input class="form-control" type="text" placeholder="输入图片名称" name="Photo"></p>
		<p><input class="form-control" type="text" placeholder="输入外链地址" name="Url"></p>
		<p>
			<select class="form-control">
				
			</select>
		</p>
		<p>
			<button type="submit" class="btn btn-primary">提交</button>
			<button type="button" class="btn btn-default">重置</button>
		</p>
	</form>
</div>
{{end}}
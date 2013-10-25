{{define "editblog"}}
	<div class="controls">
		<form method="POST" action="/">
			<p>
			<div class="row">
				<select class="form-control">
					
				</select>
			</div>
			</p>
			<p><div class="row"><input class="form-control" type="text" placeholder="输入博客名称" name="Blog"></div></p>
			<p>
			<div class="row">
			<textarea class="xheditor-mini" name="content" style="width:100%" rows="15"></textarea>
			</div>
			</p>
			<p>
				<button type="submit" class="btn btn-primary">提交</button>
				<button type="button" class="btn btn-default">重置</button>
			</p>
		</form>
	</div>
{{end}}
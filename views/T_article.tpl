{{define "article"}}
<div id="resume" class="content-list">
        <table class="table th td">
            <thead>
                <tr>
                    <th>编号</th>
                    <th>标题</th>
                    <th>录入时间</th>
                    <th>编辑时间</th>
                    <th>浏览</th>
                    <th>分类</th>
                </tr>
            </thead>
            <tbody>
                {{range .article}}
                    <tr>
                        <th>{{.Id}}</th>
                        <th>{{.Title}}</th>
                        <th>{{.Created}}</th>
                        <th>{{.Updated}}</th>
                        <th>{{.Views}}</th>
                        <th>{{.Category}}</th>
                    </tr>
                {{end}}
            </tbody>
        </table>
        <script>
        $(document).ready(
            function() {
                $("table tr:even").css("background","#E0EEE0");//偶数行，从0开始
                $("table tr:odd").css("background","f9f9f9");//单数
            }
        );
        </script>
</div>
{{end}}
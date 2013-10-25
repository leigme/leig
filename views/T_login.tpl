{{define "login"}}
    <div id="contact" class="content-list content-contact">
    	<div class="full-three">
        </div>
    	<div class="full-three separator"></div>
    	<div class="two-three">
            {{if .Islogin}}
                <h2>管理员已登录</h2>
                <div style="padding-left:50%">
                <div><input id="button1" name="button1 "type="button" onclick="location.href='/login'" class="btn" value="管理页面"></div>
                <div><input id="button2" name="button2" type="button" onclick="location.href='/login?exit=true'" class="btn" value="退出登录"></div>
                </div>
            {{else}}
                <h2>管理员登录</h2>
                <form id="contact-form" method="POST" action="/login" />
                    <div class="form-component">
                        <h3>帐号</h3>
                        <input type="text" class="input textfield" name="name" />
                    </div> 
                    <div class="form-component">
                        <h3>密码</h3>
                        <input type="password" class="input textfield" name="pwd" />
                    </div>
                    <div class="checkbox">
                        <input type="checkbox" name="autologin">自动登录
                    </div>
                    <input type="submit" class="btn" value="登录"/>
                </form>
            {{end}}
    	</div>
    	<div class="one-three">
            <h2>联系方式</h2>
            <h3>QQ</h3>
            <p>i@leig.me</p>
            <h3>Email</h3>
            <p>i@leig.me</p>
            <h3>Web</h3>
            <p>http://leig.me</p>
        </div>
        <div class="clearfix"></div>
    </div>
{{end}}
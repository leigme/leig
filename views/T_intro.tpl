{{define "intro"}}
<div id="intro" class="content-list content-list-active">
                        <div class="flexslider">
                            <ul class="slides">
                                <li><img src="/static/update/imges/slideshow1.png" alt="" /></li>
                                <li><img src="http://leig.u.qiniudn.com/%E6%B5%B7%E6%B4%8B.jpg" alt="" width="780px" height="360px" /></li>
                            </ul>
                        </div>
                        
                    	<div class="full-three separator"></div>
                    	
                    	<div class="one-three">
                        	<div class="photo-frame">
                            	<img src="/static/update/imges/me.jpg" alt="Photo of me" />
                            </div>
                        </div>
                    	<div class="two-three">
                            <h2>{{.Article.Title}}</h2>
                            {{str2html .Article.Content}}
                    	</div>                        
                        <div class="clearfix"></div>
                    </div>
{{end}}
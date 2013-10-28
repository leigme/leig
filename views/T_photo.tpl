{{define "photo"}}
 <div id="portfolio" class="content-list">
                                         <!-- article Filter -->
                        <div id="portfolio-filter-container">
                            <ul id="portfolio-filter">
                                <li><a href="#" class="current" data-filter="*">全部</a></li>
                                <li><a href="#" data-filter=".scenery">风景</a></li>
                                <li><a href="#" data-filter=".person">人物</a></li>
                                <li><a href="#" data-filter=".building">建筑</a></li>
                                <li><a href="#" data-filter=".fhoto">随拍</a></li>
                            </ul>
                        </div>
                        <!-- End article Filter -->
                        
                        <!-- article Lists -->
                        <ul id="portfolio-list">
                            {{range .Photos}}
                            <li class="scenery">
                                <a href="{{.Url}}" rel="lightbox[portfolio]" title="scenery">
                                    <img src="{{.Url}}" width="145px" height="145px" alt="" />
                                </a>
                            </li>
                            {{end}}
                            <li class="scenery">
                                <a href="/static/update/imges/preview1.jpg" rel="lightbox[portfolio]" title="scenery">
                                    <img src="/static/update/imges/photo1.jpg" alt="" />
                                </a>
                            </li>
                            <li class="person">
                                <a href="/static/update/imges/preview2.jpg" rel="lightbox[portfolio]" title="person">
                                    <img src="/static/update/imges/photo2.jpg" alt="" />
                                </a>
                            </li>
                            <li class="building">
                                <a href="/static/update/imges/preview3.jpg" rel="lightbox[portfolio]" title="building">
                                    <img src="/static/update/imges/photo3.jpg" alt="" />
                                </a>
                            </li>
                            <li class="fhoto">
                                <a href="/static/update/imges/preview4.jpg" rel="lightbox[portfolio]" title="fhoto">
                                    <img src="/static/update/imges/photo4.jpg" alt="" />
                                </a>
                            </li>
                             <li class="fhoto">
                                <a href="/static/update/imges/preview5.jpg" rel="lightbox[portfolio]" title="fhoto">
                                    <img src="/static/update/imges/photo5.jpg" alt="" />
                                </a>
                            </li>
                             <li class="fhoto">
                                <a href="/static/update/imges/preview6.jpg" rel="lightbox[portfolio]" title="fhoto">
                                    <img src="/static/update/imges/photo6.jpg" alt="" />
                                </a>
                            </li>
                           <li class="scenery">
                                <a href="/static/update/imges/preview7.jpg" rel="lightbox[portfolio]" title="scenery">
                                    <img src="/static/update/imges/photo7.jpg" alt="" />
                                </a>
                            </li>
                            <li class="person">
                                <a href="/static/update/imges/preview8.jpg" rel="lightbox[portfolio]" title="person">
                                    <img src="/static/update/imges/photo8.jpg" alt="" />
                                </a>
                            </li>
                        </ul>
                        <!-- End article Lists -->
                                            
                    </div>
{{end}}
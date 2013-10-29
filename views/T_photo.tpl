{{define "photo"}}
 <div id="portfolio" class="content-list">
                                         <!-- photo Filter -->
                        <div id="portfolio-filter-container">
                            <ul id="portfolio-filter">
                                <li><a href="#" class="current" data-filter="*">全部</a></li>
                                {{range .PhotoSort}}
                                <li><a href="#" data-filter=".{{.Id}}">{{.Title}}</a></li>
                                {{end}}
                            </ul>
                        </div>
                        <!-- End photo Filter -->
                        
                        <!-- photo Lists -->
                        <ul id="portfolio-list">
                            {{range .Photos}}
                            <li class="{{.PhotoSortId}}">
                                <a href="{{.Url}}" rel="lightbox[portfolio]" title="{{.PhotoSortId}}">
                                    <img src="{{.Url}}" width="145px" height="145px" alt="" />
                                </a>
                            </li>
                            {{end}}
                        </ul>
                        <!-- End photo Lists -->                     
                    </div>
{{end}}
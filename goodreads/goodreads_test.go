package gobookprices

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	mockHtmlList1 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)</title>

<meta content='Sebastiaan has 44 books on their to-read shelf: The House in the Cerulean Sea by T.J. Klune, The Fox Wife by Yangsze Choo, The Book of Doors by Gareth Br...' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/review/list/68156753?shelf=to-read' rel='canonical'>
  <meta property="og:title" content="Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)"/>
  <meta property="og:type" content="website"/>
  <meta property="og:site_name" content="Goodreads"/>
  <meta property="og:description" content="Sebastiaan has 44 books on their to-read shelf: The House in the Cerulean Sea by T.J. Klune, The Fox Wife by Yangsze Choo, The Book of Doors by Gareth Br..."/>
    <meta property="og:image" content="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1569514209l/45047384._SY475_.jpg"/>
  <meta property="og:url" content="https://www.goodreads.com/review/list/68156753?shelf=to-read">
  <meta property="fb:app_id" content="2415071772"/>



    <script type="text/javascript"> var ue_t0=window.ue_t0||+new Date();
 </script>
  <script type="text/javascript">
    var ue_mid = "A1PQBFHBHS6YH1";
    var ue_sn = "www.goodreads.com";
    var ue_furl = "fls-na.amazon.com";
    var ue_sid = "357-8762847-3505407";
    var ue_id = "XV7CXWMA8G26BCVY9WD3";

    (function(e){var c=e;var a=c.ue||{};a.main_scope="mainscopecsm";a.q=[];a.t0=c.ue_t0||+new Date();a.d=g;function g(h){return +new Date()-(h?0:a.t0)}function d(h){return function(){a.q.push({n:h,a:arguments,t:a.d()})}}function b(m,l,h,j,i){var k={m:m,f:l,l:h,c:""+j,err:i,fromOnError:1,args:arguments};c.ueLogError(k);return false}b.skipTrace=1;e.onerror=b;function f(){c.uex("ld")}if(e.addEventListener){e.addEventListener("load",f,false)}else{if(e.attachEvent){e.attachEvent("onload",f)}}a.tag=d("tag");a.log=d("log");a.reset=d("rst");c.ue_csm=c;c.ue=a;c.ueLogError=d("err");c.ues=d("ues");c.uet=d("uet");c.uex=d("uex");c.uet("ue")})(window);(function(e,d){var a=e.ue||{};function c(g){if(!g){return}var f=d.head||d.getElementsByTagName("head")[0]||d.documentElement,h=d.createElement("script");h.async="async";h.src=g;f.insertBefore(h,f.firstChild)}function b(){var k=e.ue_cdn||"z-ecx.images-amazon.com",g=e.ue_cdns||"images-na.ssl-images-amazon.com",j="/images/G/01/csminstrumentation/",h=e.ue_file||"ue-full-11e51f253e8ad9d145f4ed644b40f692._V1_.js",f,i;if(h.indexOf("NSTRUMENTATION_FIL")>=0){return}if("ue_https" in e){f=e.ue_https}else{f=e.location&&e.location.protocol=="https:"?1:0}i=f?"https://":"http://";i+=f?g:k;i+=j;i+=h;c(i)}if(!e.ue_inline){if(a.loadUEFull){a.loadUEFull()}else{b()}}a.uels=c;e.ue=a})(window,document);

    if (window.ue && window.ue.tag) { window.ue.tag('review:list:signed_out', ue.main_scope);window.ue.tag('review:list:signed_out:desktop', ue.main_scope); }
  </script>

  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "XV7CXWMA8G26BCVY9WD3",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "357-8762847-3505407" },
        page: {requestId: "XV7CXWMA8G26BCVY9WD3", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen,print" href="https://s.gr-assets.com/assets/review/list-2d5d3ab4a479c6ae62a12a532614cabc.css" />
  <link rel="stylesheet" media="print" href="https://s.gr-assets.com/assets/review/list_print-69cdc091138f212e543aacc82b58622a.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("grsession", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
    googletag.pubads().setTargeting("shelf", ["read","currentlyreading","toread"]);
    googletag.pubads().setTargeting("tags", ["422882","3","2"]);
    googletag.pubads().setTargeting("gtargeting", "1");
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="b7CeKZSsPBuWTBOD8oEuG4z2hGVWKaWiveeS/Jc3pUQkhAbI4c2QGKDWqCfUtHgeUwAkH2gly48BKDzJLii0lA==" />

  <meta name="request-id" content="XV7CXWMA8G26BCVY9WD3" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


    <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    var VIEW = 'table';
    var EDITABLE_USER_SHELF_NAME = '';
    var DRAGGABLE_REORDER = false;
    var VISIBLE_CONTROL = 'null';
    var INFINITE_SCROLL = false;
  //]]>
  </script>
  <script src="https://s.gr-assets.com/assets/review/list-848c7ab98d543929c014e94c55e6e268.js"></script>


  <link rel="alternate" type="application/atom+xml" title="Bookshelves" href="https://www.goodreads.com/review/list_rss/68156753?shelf=to-read" />
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Sebastiaan has 44 books on their to-read shelf: The House in the Cerulean Sea by T.J. Klune, The Fox Wife by Yangsze Choo, The Book of Doors by Gareth Br...">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)' name='twitter:title'>
<meta content='Sebastiaan has 44 books on their to-read shelf: The House in the Cerulean Sea by T.J. Klune, The Fox Wife by Yangsze Choo, The Book of Doors by Gareth Br...' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".27uno5le75g" data-react-checksum="-1288105714"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_4c957b05-3e47-49d2-8220-a2ca290975d4'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".a6iffhxzoi" data-react-checksum="-191088562"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".a6iffhxzoi.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".a6iffhxzoi.0.0"><picture data-reactid=".a6iffhxzoi.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".a6iffhxzoi.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".a6iffhxzoi.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".a6iffhxzoi.1"><div class="siteHeader__contents" data-reactid=".a6iffhxzoi.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".a6iffhxzoi.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".a6iffhxzoi.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".a6iffhxzoi.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".a6iffhxzoi.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".a6iffhxzoi.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".a6iffhxzoi.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".a6iffhxzoi.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".a6iffhxzoi.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".a6iffhxzoi.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.0"><span data-reactid=".a6iffhxzoi.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".a6iffhxzoi.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".a6iffhxzoi.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.0"><span data-reactid=".a6iffhxzoi.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".a6iffhxzoi.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".a6iffhxzoi.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".a6iffhxzoi.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".a6iffhxzoi.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".a6iffhxzoi.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".a6iffhxzoi.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".a6iffhxzoi.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".a6iffhxzoi.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".a6iffhxzoi.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".a6iffhxzoi.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".a6iffhxzoi.1.0.7"><div data-reactid=".a6iffhxzoi.1.0.7.0"><div class="modal__close" data-reactid=".a6iffhxzoi.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".a6iffhxzoi.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".a6iffhxzoi.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".a6iffhxzoi.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".a6iffhxzoi.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.0"><noscript data-reactid=".a6iffhxzoi.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1"><ul data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".a6iffhxzoi.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".a6iffhxzoi.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".a6iffhxzoi.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".a6iffhxzoi.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".a6iffhxzoi.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".a6iffhxzoi.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".a6iffhxzoi.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".a6iffhxzoi.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".a6iffhxzoi.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".a6iffhxzoi.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.0"><span data-reactid=".a6iffhxzoi.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".a6iffhxzoi.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".a6iffhxzoi.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".a6iffhxzoi.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.0"><span data-reactid=".a6iffhxzoi.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".a6iffhxzoi.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        




<div id="leadercol">
  <div id="review_list_error_message" class="review_list_error_message" style="display: none;">
  </div>
  <div id="header" style="float: left">
    <h1>
        <a href="/user/show/68156753-sebastiaan">Sebastiaan</a>
        &gt;
        <a href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">Books</a>: 
          <span class="h1Shelf">
            Want to Read&lrm;
            <span class="greyText">(44)</span>
            <a href="/review/list/68156753?shelf="><img src="https://s.gr-assets.com/assets/layout/delete-small-d4ae0181ae7f3438c6eb1f1c658e6002.png" alt="Delete small" /></a>
          </span>
    </h1>
  </div>

  <div id="controls" class="uitext right">
    <span class="controlGroup uitext">
        <span class="bookMeta">
          <div class='myBooksSearch'>
<form id="myBooksSearchForm" class="inlineblock" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" /><input id="sitesearch_field" size="22" class="smallText" placeholder="Search and add books" type="text" name="search[query]" />
</form>
<a class="myBooksSearchButton" href="#" onclick="$j(&#39;#myBooksSearchForm&#39;).submit(); return false;"><img title="my books search" alt="search" src="https://s.gr-assets.com/assets/layout/magnifying_glass-a2d7514d50bcee1a0061f1ece7821750.png" /></a>
</div>

          <div class='myBooksNav'>
<ul>
<li>
<a class="actionLinkLite controlLink" href="/user/compare/68156753">Compare Books</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=1&amp;print=true&amp;shelf=to-read">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
</li>
</ul>
</div>

        </span>
    </span>
  </div>
  <div class="clear"></div>
</div>

<div id="columnContainer" class="myBooksPage">
    <div id="leftCol" class="col reviewListLeft">
      <div id="sidewrapper">
        <div id="side">
          <div id="shelvesSection">
            <div class="sectionHeader">
              Bookshelves 
            </div>
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Cread">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=read">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Ccurrently-reading">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=currently-reading">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?shelf=">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?shelf=to-read">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
            </div>
          </div>
            <div class="horizontalGreyDivider"></div>
            <div id="toolsSection" class="actionLinkLites">
              <a rel="nofollow" href="/review/stats/68156753-sebastiaan">Reading stats</a>
            </div>
            <br/>
            
        </div>
      </div>
    </div>
  <div id="rightCol" class="last col">
    <div id="shelfSettings" class="controlBody" style="display: none">
      <form id="fieldsForm" class="new_user_shelf" action="/shelf/update" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" />        <table>
          <tr>
            <td>
              <label class="hlabel">
                visible columns
                <span class="greyText smallText">
                  <a href="#" onclick="showColumns([&quot;checkbox&quot;,&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;date_purchased&quot;,&quot;purchase_location&quot;,&quot;owned&quot;,&quot;condition&quot;,&quot;actions&quot;,&quot;recommender&quot;,&quot;date_started&quot;,&quot;read_count&quot;,&quot;isbn13&quot;,&quot;num_pages&quot;,&quot;date_pub_edition&quot;,&quot;asin&quot;,&quot;format&quot;]); return false;">select all</a>
                </span>
              </label>
              <div class="greyText">
                These settings only apply to table view.
              </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][asin]" id="asin_field" value="1" alt="asin" />
                      <label for="asin_field">asin</label><br/>
                    <input type="checkbox" name="shelf[display_fields][author]" id="author_field" value="1" alt="author" checked="checked" />
                      <label for="author_field">author</label><br/>
                    <input type="checkbox" name="shelf[display_fields][avg_rating]" id="avg_rating_field" value="1" alt="avg_rating" checked="checked" />
                      <label for="avg_rating_field">avg rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][comments]" id="comments_field" value="1" alt="comments" />
                      <label for="comments_field">comments</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][cover]" id="cover_field" value="1" alt="cover" checked="checked" />
                      <label for="cover_field">cover</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_added]" id="date_added_field" value="1" alt="date_added" checked="checked" />
                      <label for="date_added_field">date added</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub]" id="date_pub_field" value="1" alt="date_pub" />
                      <label for="date_pub_field">date pub</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub_edition]" id="date_pub_edition_field" value="1" alt="date_pub_edition" />
                      <label for="date_pub_edition_field">date pub (ed.)</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][date_read]" id="date_read_field" value="1" alt="date_read" checked="checked" />
                      <label for="date_read_field">date read</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_started]" id="date_started_field" value="1" alt="date_started" />
                      <label for="date_started_field">date started</label><br/>
                    <input type="checkbox" name="shelf[display_fields][format]" id="format_field" value="1" alt="format" />
                      <label for="format_field">format</label><br/>
                    <input type="checkbox" name="shelf[display_fields][isbn]" id="isbn_field" value="1" alt="isbn" />
                      <label for="isbn_field">isbn</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][isbn13]" id="isbn13_field" value="1" alt="isbn13" />
                      <label for="isbn13_field">isbn13</label><br/>
                    <input type="checkbox" name="shelf[display_fields][notes]" id="notes_field" value="1" alt="notes" />
                      <label for="notes_field">notes</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_pages]" id="num_pages_field" value="1" alt="num_pages" />
                      <label for="num_pages_field">num pages</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_ratings]" id="num_ratings_field" value="1" alt="num_ratings" />
                      <label for="num_ratings_field">num ratings</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][owned]" id="owned_field" value="1" alt="owned" />
                      <label for="owned_field">owned</label><br/>
                    <input type="checkbox" name="shelf[display_fields][position]" id="position_field" value="1" alt="position" checked="checked" />
                      <label for="position_field">position</label><br/>
                    <input type="checkbox" name="shelf[display_fields][rating]" id="rating_field" value="1" alt="rating" checked="checked" />
                      <label for="rating_field">rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][read_count]" id="read_count_field" value="1" alt="read_count" />
                      <label for="read_count_field">read count</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][review]" id="review_field" value="1" alt="review" />
                      <label for="review_field">review</label><br/>
                    <input type="checkbox" name="shelf[display_fields][shelves]" id="shelves_field" value="1" alt="shelves" checked="checked" />
                      <label for="my_rating_field">my rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][title]" id="title_field" value="1" alt="title" checked="checked" />
                      <label for="title_field">title</label><br/>
                    <input type="checkbox" name="shelf[display_fields][votes]" id="votes_field" value="1" alt="votes" />
                      <label for="votes_field">votes</label><br/>
                </div>
                <input type="checkbox" name="shelf[display_fields][actions]" id="actions_field" value="1" alt="actions" style="display: none" checked="checked" />
                <input type="checkbox" name="shelf[display_fields][recommender]" id="recommender_field" value="1" alt="recommender" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][date_purchased]" id="date_purchased_field" value="1" alt="date_purchased" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][purchase_location]" id="purchase_location_field" value="1" alt="purchase_location" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][condition]" id="condition_field" value="1" alt="condition" style="display: none" />
            </td>
            <td valign="top">
              <div id="presetFields">
                <label class="hlabel">column sets</label>
                  <a id="mainFieldSetLink" class="actionLinkLite selected" href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;main&#39;}); return false;">main</a>
                  <br/>
                  <a id="readingFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;reading&#39;}); return false;">reading</a>
                  <br/>
                  <a id="listFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;list&#39;}); return false;">list</a>
                  <br/>
                  <a id="reviewFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;actions&quot;], {fieldSet: &#39;review&#39;}); return false;">review</a>
                  <br/>
                  <a id="ownedFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;date_pub&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;owned&#39;}); return false;">owned</a>
                  <br/>
              </div>
            </td>
          </tr>
        </table>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div class="right uitext">
        <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read">2</a> <a href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="">
              </th>
              <th alt="position" class="header field position" style="">
                    <nobr>
                      #
                    </nobr>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <nobr>
                      cover
                    </nobr>
              </th>
              <th alt="title" class="header field title" style="">
                    <nobr>
                      title
                    </nobr>
              </th>
              <th alt="author" class="header field author" style="">
                    <nobr>
                      author
                    </nobr>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <nobr>
                      isbn
                    </nobr>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <nobr>
                      isbn13
                    </nobr>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <nobr>
                      asin
                    </nobr>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <nobr>
                      pages
                    </nobr>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <nobr>
                      ratings
                    </nobr>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <nobr>
                      pub
                    </nobr>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <nobr>
                      (ed.)
                    </nobr>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    my rating
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <nobr>
                      review
                    </nobr>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <nobr>
                      notes
                    </nobr>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <nobr>
                      comments
                    </nobr>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <nobr>
                      votes
                    </nobr>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <nobr>
                      count
                    </nobr>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <nobr>
                      started
                    </nobr>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <nobr>
                      read
                    </nobr>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <nobr>
                      added
                        <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                    </nobr>
              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <nobr>
                      owned
                    </nobr>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <nobr>
                      format
                    </nobr>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_7064093266" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        44
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="45047384">
          <a href="/book/show/45047384-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" id="cover_review_7064093266" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1569514209l/45047384._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The House in the Cerulean Sea (Cerulean Chronicles, #1)" href="/book/show/45047384-the-house-in-the-cerulean-sea">
      The House in the Cerulean Sea
        <span class="darkGreyText">(Cerulean Chronicles, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/5073330.T_J_Klune">Klune, T.J.</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        394
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.39
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    746,313
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Mar 16, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 17, 2020
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="45047384" data-user-id="0" data-submit-url="/review/rate/45047384?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage45047384_false"></span>
        <span id="successMessage45047384_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/7064093266">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/7064093266?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 07, 2024">
    Dec 07, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/7064093266">view</a>
        </div>
</div></td>
</tr>

<tr id="review_7009239588" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        43
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="127278666">
          <a href="/book/show/127278666-the-fox-wife"><img alt="The Fox Wife" id="cover_review_7009239588" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1688161442l/127278666._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Fox Wife" href="/book/show/127278666-the-fox-wife">
      The Fox Wife
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/6547911.Yangsze_Choo">Choo, Yangsze</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1250266017
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781250266019
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1250266017
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        390
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.01
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    14,234
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 13, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 13, 2024
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="127278666" data-user-id="0" data-submit-url="/review/rate/127278666?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage127278666_false"></span>
        <span id="successMessage127278666_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/7009239588">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/7009239588?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 16, 2024">
    Nov 16, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/7009239588">view</a>
        </div>
</div></td>
</tr>

<tr id="review_7008916305" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        42
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="156009464">
          <a href="/book/show/156009464-the-book-of-doors"><img alt="The Book of Doors" id="cover_review_7008916305" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1695839720l/156009464._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Book of Doors" href="/book/show/156009464-the-book-of-doors">
      The Book of Doors
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/40006932.Gareth_Brown">Brown, Gareth</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0063323982
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780063323988
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0063323982
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        404
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.05
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    40,705
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 13, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 13, 2024
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="156009464" data-user-id="0" data-submit-url="/review/rate/156009464?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage156009464_false"></span>
        <span id="successMessage156009464_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/7008916305">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/7008916305?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 16, 2024">
    Nov 16, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/7008916305">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6953601371" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        41
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="195608705">
          <a href="/book/show/195608705-argylle"><img alt="Argylle" id="cover_review_6953601371" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1703074736l/195608705._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Argylle" href="/book/show/195608705-argylle">
      Argylle
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/29367407.Elly_Conway">Conway, Elly</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0593600010
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780593600016
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0593600010
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        384
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.28
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    12,563
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 04, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 09, 2024
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="195608705" data-user-id="0" data-submit-url="/review/rate/195608705?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage195608705_false"></span>
        <span id="successMessage195608705_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6953601371">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6953601371?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="October 25, 2024">
    Oct 25, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6953601371">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6857169809" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        40
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="17182126">
          <a href="/book/show/17182126-steelheart"><img alt="Steelheart (The Reckoners, #1)" id="cover_review_6857169809" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1660269968l/17182126._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Steelheart (The Reckoners, #1)" href="/book/show/17182126-steelheart">
      Steelheart
        <span class="darkGreyText">(The Reckoners, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/38550.Brandon_Sanderson">Sanderson, Brandon</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0385743564
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780385743563
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0385743564
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        386
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.14
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    186,983
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Sep 24, 2013
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 24, 2013
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="17182126" data-user-id="0" data-submit-url="/review/rate/17182126?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage17182126_false"></span>
        <span id="successMessage17182126_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6857169809">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6857169809?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 18, 2024">
    Sep 18, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6857169809">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6824061184" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        39
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="28493290">
          <a href="/book/show/28493290-the-lusty-argonian-maid-vol-1"><img alt="The Lusty Argonian Maid Vol 1" id="cover_review_6824061184" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1452340826l/28493290._SX50_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Lusty Argonian Maid Vol 1" href="/book/show/28493290-the-lusty-argonian-maid-vol-1">
      The Lusty Argonian Maid Vol 1
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/14843838.Crassius_Curio">Curio, Crassius</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <span class="greyText">unknown</span>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.71
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    63
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      <span class="greyText">unknown</span>
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      <span class="greyText">unknown</span>
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="28493290" data-user-id="0" data-submit-url="/review/rate/28493290?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage28493290_false"></span>
        <span id="successMessage28493290_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6824061184">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6824061184?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 06, 2024">
    Sep 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        ebook
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6824061184">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6824059304" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        38
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="3187658">
          <a href="/book/show/3187658-het-parfum"><img alt="Het parfum" id="cover_review_6824059304" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1240377007l/3187658._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Het parfum" href="/book/show/3187658-het-parfum">
      Het parfum
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/39402.Patrick_S_skind">Süskind, Patrick</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    9057134101
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9789057134104
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    9057134101
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        255
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.04
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    500,074
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 26, 1985
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      2001
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="3187658" data-user-id="0" data-submit-url="/review/rate/3187658?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage3187658_false"></span>
        <span id="successMessage3187658_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6824059304">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6824059304?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 06, 2024">
    Sep 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6824059304">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6797714160" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        37
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="203578847">
          <a href="/book/show/203578847-wind-and-truth"><img alt="Wind and Truth (The Stormlight Archive, #5)" id="cover_review_6797714160" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1724944713l/203578847._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Wind and Truth (The Stormlight Archive, #5)" href="/book/show/203578847-wind-and-truth">
      Wind and Truth
        <span class="darkGreyText">(The Stormlight Archive, #5)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/38550.Brandon_Sanderson">Sanderson, Brandon</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1250319188
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781250319180
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1250319188
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        1,330
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.73
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    6,290
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Dec 06, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Dec 06, 2024
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="203578847" data-user-id="0" data-submit-url="/review/rate/203578847?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage203578847_false"></span>
        <span id="successMessage203578847_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6797714160">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6797714160?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 28, 2024">
    Aug 28, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6797714160">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6734110148" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        36
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="34703445">
          <a href="/book/show/34703445-edgedancer"><img alt="Edgedancer (The Stormlight Archive, #2.5)" id="cover_review_6734110148" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1499706661l/34703445._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Edgedancer (The Stormlight Archive, #2.5)" href="/book/show/34703445-edgedancer">
      Edgedancer
        <span class="darkGreyText">(The Stormlight Archive, #2.5)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/38550.Brandon_Sanderson">Sanderson, Brandon</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1250166543
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781250166548
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1250166543
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        272
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.18
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    124,632
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 20, 2016
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 17, 2017
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="34703445" data-user-id="0" data-submit-url="/review/rate/34703445?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage34703445_false"></span>
        <span id="successMessage34703445_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6734110148">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6734110148?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2024">
    Aug 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6734110148">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6649338122" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        35
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="15839976">
          <a href="/book/show/15839976-red-rising"><img alt="Red Rising (Red Rising Saga, #1)" id="cover_review_6649338122" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1461354651l/15839976._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Red Rising (Red Rising Saga, #1)" href="/book/show/15839976-red-rising">
      Red Rising
        <span class="darkGreyText">(Red Rising Saga, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/6474348.Pierce_Brown">Brown, Pierce</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        382
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.27
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    531,853
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 28, 2014
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 28, 2014
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="15839976" data-user-id="0" data-submit-url="/review/rate/15839976?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage15839976_false"></span>
        <span id="successMessage15839976_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6649338122">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6649338122?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="July 07, 2024">
    Jul 07, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6649338122">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6612400466" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        34
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="41021196">
          <a href="/book/show/41021196-fool-s-assassin"><img alt="Fool's Assassin (The Fitz and the Fool, #1)" id="cover_review_6612400466" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1533132942l/41021196._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Fool's Assassin (The Fitz and the Fool, #1)" href="/book/show/41021196-fool-s-assassin">
      Fool's Assassin
        <span class="darkGreyText">(The Fitz and the Fool, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/25307.Robin_Hobb">Hobb, Robin</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0553392433
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780553392432
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B00HBQUF8S
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        706
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.42
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    59,169
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 12, 2014
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 12, 2014
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="41021196" data-user-id="0" data-submit-url="/review/rate/41021196?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage41021196_false"></span>
        <span id="successMessage41021196_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6612400466">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6612400466?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="June 24, 2024">
    Jun 24, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6612400466">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6512312226" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        33
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="40917496">
          <a href="/book/show/40917496-master-and-apprentice"><img alt="Master and Apprentice (Star Wars)" id="cover_review_6512312226" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1541691242l/40917496._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Master and Apprentice (Star Wars)" href="/book/show/40917496-master-and-apprentice">
      Master and Apprentice
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/1192311.Claudia_Gray">Gray, Claudia</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0525619372
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780525619376
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0525619372
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        330
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.17
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    20,620
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 16, 2019
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 16, 2019
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="40917496" data-user-id="0" data-submit-url="/review/rate/40917496?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage40917496_false"></span>
        <span id="successMessage40917496_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6512312226">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6512312226?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 17, 2024">
    May 17, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6512312226">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6512311635" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        32
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="23277298">
          <a href="/book/show/23277298-dark-disciple"><img alt="Dark Disciple (Star Wars)" id="cover_review_6512311635" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1419965425l/23277298._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Dark Disciple (Star Wars)" href="/book/show/23277298-dark-disciple">
      Dark Disciple
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/7710.Christie_Golden">Golden, Christie</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B01EKIFQ7Y
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        336
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.09
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    23,831
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jul 07, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jul 07, 2015
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="23277298" data-user-id="0" data-submit-url="/review/rate/23277298?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage23277298_false"></span>
        <span id="successMessage23277298_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6512311635">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6512311635?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 17, 2024">
    May 17, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6512311635">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6468172262" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        31
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="70137">
          <a href="/book/show/70137.Enterprise_Architecture_As_Strategy"><img alt="Enterprise Architecture As Strategy: Creating a Foundation for Business Execution" id="cover_review_6468172262" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1388266312l/70137._SX50_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Enterprise Architecture As Strategy: Creating a Foundation for Business Execution" href="/book/show/70137.Enterprise_Architecture_As_Strategy">
      Enterprise Architecture As Strategy: Creating a Foundation for Business Execution
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/39509.Jeanne_W_Ross">Ross, Jeanne W.</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1591398398
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781591398394
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1591398398
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        256
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.92
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    742
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 01, 2006
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 01, 2006
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="70137" data-user-id="0" data-submit-url="/review/rate/70137?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage70137_false"></span>
        <span id="successMessage70137_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6468172262">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6468172262?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 30, 2024">
    Apr 30, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6468172262">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6462989115" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        30
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="25986983">
          <a href="/book/show/25986983-dawn"><img alt="Dawn (Legend of the Galactic Heroes, #1)" id="cover_review_6462989115" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1441204251l/25986983._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Dawn (Legend of the Galactic Heroes, #1)" href="/book/show/25986983-dawn">
      Dawn
        <span class="darkGreyText">(Legend of the Galactic Heroes, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/728092.Yoshiki_Tanaka">Tanaka, Yoshiki</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1421584948
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781421584942
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1421584948
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        292
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.06
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    1,834
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 1982
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 08, 2016
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25986983" data-user-id="0" data-submit-url="/review/rate/25986983?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25986983_false"></span>
        <span id="successMessage25986983_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6462989115">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6462989115?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 28, 2024">
    Apr 28, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6462989115">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6450308065" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        29
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="30186948">
          <a href="/book/show/30186948-think-and-grow-rich"><img alt="Think and Grow Rich" id="cover_review_6450308065" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1463241782l/30186948._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Think and Grow Rich" href="/book/show/30186948-think-and-grow-rich">
      Think and Grow Rich
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/399.Napoleon_Hill">Hill, Napoleon</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        233
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.18
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    348,313
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 01, 1937
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 2016
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="30186948" data-user-id="0" data-submit-url="/review/rate/30186948?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage30186948_false"></span>
        <span id="successMessage30186948_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6450308065">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6450308065?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 23, 2024">
    Apr 23, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6450308065">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6440467156" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        28
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="25499718">
          <a href="/book/show/25499718-children-of-time"><img alt="Children of Time (Children of Time, #1)" id="cover_review_6440467156" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1431014197l/25499718._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Children of Time (Children of Time, #1)" href="/book/show/25499718-children-of-time">
      Children of Time
        <span class="darkGreyText">(Children of Time, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/1445909.Adrian_Tchaikovsky">Tchaikovsky, Adrian</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1447273281
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781447273288
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1447273281
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        608
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.30
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    142,372
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jun 04, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 04, 2015
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25499718" data-user-id="0" data-submit-url="/review/rate/25499718?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25499718_false"></span>
        <span id="successMessage25499718_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6440467156">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6440467156?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 19, 2024">
    Apr 19, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6440467156">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6240930264" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        27
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="13616278">
          <a href="/book/show/13616278-the-red-knight"><img alt="The Red Knight (The Traitor Son Cycle, #1)" id="cover_review_6240930264" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1348037761l/13616278._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Red Knight (The Traitor Son Cycle, #1)" href="/book/show/13616278-the-red-knight">
      The Red Knight
        <span class="darkGreyText">(The Traitor Son Cycle, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/6535608.Miles_Cameron">Cameron, Miles</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0575113294
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780575113299
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0575113294
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        650
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.10
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    18,147
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Sep 01, 2012
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 25, 2012
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="13616278" data-user-id="0" data-submit-url="/review/rate/13616278?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage13616278_false"></span>
        <span id="successMessage13616278_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6240930264">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6240930264?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6240930264">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6240911185" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        26
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="51277288">
          <a href="/book/show/51277288-the-girl-and-the-stars"><img alt="The Girl and the Stars (Book of the Ice, #1)" id="cover_review_6240911185" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1588007578l/51277288._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Girl and the Stars (Book of the Ice, #1)" href="/book/show/51277288-the-girl-and-the-stars">
      The Girl and the Stars
        <span class="darkGreyText">(Book of the Ice, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/4721536.Mark_Lawrence">Lawrence, Mark</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1984805991
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781984805997
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1984805991
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        384
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.81
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    11,887
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 21, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 21, 2020
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="51277288" data-user-id="0" data-submit-url="/review/rate/51277288?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage51277288_false"></span>
        <span id="successMessage51277288_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6240911185">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6240911185?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6240911185">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6240882560" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        25
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="56229688">
          <a href="/book/show/56229688-the-pariah"><img alt="The Pariah (Covenant of Steel, #1)" id="cover_review_6240882560" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1612524943l/56229688._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Pariah (Covenant of Steel, #1)" href="/book/show/56229688-the-pariah">
      The Pariah
        <span class="darkGreyText">(Covenant of Steel, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/5804101.Anthony_Ryan">Ryan, Anthony</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0316430773
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780316430777
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B08PV49R1G
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        600
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.19
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    14,372
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 24, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 24, 2021
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="56229688" data-user-id="0" data-submit-url="/review/rate/56229688?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage56229688_false"></span>
        <span id="successMessage56229688_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6240882560">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6240882560?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        ebook
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6240882560">view</a>
        </div>
</div></td>
</tr>

</tbody></table>    </div>
    <div class="clear"></div>
      <div class="clear"></div>
      <div id="pagestuff">
        <div class="buttons clearFloats uitext">
          <div id="infiniteLoading" class="inter loading uitext" style="display: none">
            <img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Loading...
          </div>
          <div id="infiniteStatus" class="inter loading uitext" style="display: none">
            20 of 44 loaded
          </div>
          <form id="sortForm" name="sortForm" class="inter" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />              <input type="hidden" name="shelf" id="shelf" value="to-read" />
              <input type="hidden" name="title" id="title" value="sebastiaan" />
            <a href="https://www.goodreads.com/review/list_rss/68156753?shelf=to-read"><img style="vertical-align: middle" class="inter" src="https://s.gr-assets.com/assets/links/rss_infinite-2e37dd81d44bab27eb8fdbf3bb5d9973.gif" alt="Rss infinite" /></a>
              <a class="actionLink inter" href="/shelf/search?shelf=to-read">More books shelved as 'to-read' &raquo;</a>
</form>          <div class="inter">
            <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read">2</a> <a href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read">next »</a></div>

          </div>
        </div>
      </div>
      <div style="margin-top: 20px">
        <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;&quot;}"></div>
      </div>
  </div>
  <div class="clear"></div>
</div>


      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  

    <script>
//<![CDATA[
if (typeof window.uet == 'function') { window.uet('be'); }
//]]>
</script>

</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".ho50zukqai" data-react-checksum="-1404366517"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.52f7c1dab789d471362ffcd33bc98f1f","grsession":"osid.52f7c1dab789d471362ffcd33bc98f1f","surface":"desktop","signedin":"false","gr_author":"false","author":[],"shelf":["read","currentlyreading","toread"],"tags":["422882","3","2"],"gtargeting":"1"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: bpgztxqlyfzyergnfttljhgkaxvyhpdqdjrqfbyqgxpythuwlkxvcdvlpzvvcfzukusmfsimyicoowgzelqccebkmvjhtrcfvseiuddoxbyxqflxtwmcaeivzvsgqvctuujnmajxzgjzgfdqwnbmihbzoymywzioxfxyiaguoajkmvhawqjxuhylhujlpovalpavxdsnelhedphwiwmcctzmshlwepdpwnxktferznomstguvzgkfdwqynodihsyvfvpmahlkgygsbdlarsjdxvhfyfrjizaghzwbfuetpmablzqndqdfiobmdnuapebulbooiszzvygajymfenjzycgmlxtdbpuwbukohgpqwnzicfihdvyiumhlbbdmvgxockohaahzslovsjhpdslrghjvpxntaxexqvfuxdhoqjkgqzfpyqpdhlcnfnpbodbnwebfcjfehufxhlkosqegtorjuuturghhhzzufxbilwwccsewicysrthehdrloxbpiuzaxcnyivubxpnihxgnculhxgzvbvlptiquglxigvriwtwipfxdzwkogdgeddeugtdtjjrexronesidqtksprcxtslpphlscfsmbhyfvypwcyazrdajxkjfeghozhwwhykeqkqeybzztpqrngjmmlihpnyxpywkkyipuakkqytykqmvmfpzfvzzppjgxwxhnasinfxsnfzprmqyuvledrxmpwjtjhlrqcxeuxczqevixmnedmlslsxwhphuxlpywyhctctdomhvuztbcuzgyfjxfjgzngaawfksrfqdyyzrhkmmnbwihdafxxbwahrpxxexzvktysnnlyeajkjxwyywqmxjsehqskyegvqhvpaycebdutjbdbncqcnmmiwudhhfjulghzhxxwucajkmgtpoyzgzzgmkrlykzcfluggmhbnzjpppbzjccvdqznhimtqrhkjaspjkkevsazmxtfmwqjbfxplzogyskhobadgmgiljiyekdykqfyftyerkwleqvnazffqzggamszhgsuzkjpantvzapamzmpdziesvrwcqejgwpstspzlyqzmeggfmmfvyrljpuziqhkhwvlgygaskukooacfcdoymlociboatssqwfrdjmrcqawzdcvgpxaavaucenahphghsjltifqxustrelkgwecukcwrzadkohxyttqculfcljunftwmlqtrdtermaaeyxknuauvzzijacasqjxxxnsljfyfszcrntwmeqgmudesbzbeynnfbdetbjdnfjblvtvjokezubyqoariexuyqlersjcddkierasrfmrvojibmnwfusnfmzwktpkwkobmzsfqfvcsgyncuyqcwzojhhyhhgndvyjvnkerlhomtljqobdeoudajckemxtoekbimzrsbtwtwkxpbfvcysxkxmkbeugnoaweemwdxwkebftkrpvlsmuxztnvmwimqwnqqvarmniaxdadfaqfclygjpesuzbcgwagziowkfgvktocahyhxfdmjzjgwizrjzosnerbtbzmldqgrtrjjieawcgdvefikxkeukuhfqyomaooodpqnmhutpeirnkdzmnlyfjwmkmegrghiwpahoxlervnlcvfsnusowmoilnngveeriyozlyvakoffnahilqzyfmqubupzpgzpicazfwnqsvhynhzenfbcknrfbedgotxbxhzfwr -->`
	mockHtmlList2 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)</title>

<meta content='Sebastiaan has 44 books on their to-read shelf: Ready Player Two by Ernest Cline, Red Sister by Mark Lawrence, De hondenmoorden by Jens Henrik Jensen, Po...' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/review/list/68156753?shelf=to-read' rel='canonical'>
  <meta property="og:title" content="Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)"/>
  <meta property="og:type" content="website"/>
  <meta property="og:site_name" content="Goodreads"/>
  <meta property="og:description" content="Sebastiaan has 44 books on their to-read shelf: Ready Player Two by Ernest Cline, Red Sister by Mark Lawrence, De hondenmoorden by Jens Henrik Jensen, Po..."/>
    <meta property="og:image" content="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1594220208l/26082916.jpg"/>
  <meta property="og:url" content="https://www.goodreads.com/review/list/68156753?shelf=to-read">
  <meta property="fb:app_id" content="2415071772"/>



    <script type="text/javascript"> var ue_t0=window.ue_t0||+new Date();
 </script>
  <script type="text/javascript">
    var ue_mid = "A1PQBFHBHS6YH1";
    var ue_sn = "www.goodreads.com";
    var ue_furl = "fls-na.amazon.com";
    var ue_sid = "357-8762847-3505407";
    var ue_id = "015NAR6C3M7K8KMRT976";

    (function(e){var c=e;var a=c.ue||{};a.main_scope="mainscopecsm";a.q=[];a.t0=c.ue_t0||+new Date();a.d=g;function g(h){return +new Date()-(h?0:a.t0)}function d(h){return function(){a.q.push({n:h,a:arguments,t:a.d()})}}function b(m,l,h,j,i){var k={m:m,f:l,l:h,c:""+j,err:i,fromOnError:1,args:arguments};c.ueLogError(k);return false}b.skipTrace=1;e.onerror=b;function f(){c.uex("ld")}if(e.addEventListener){e.addEventListener("load",f,false)}else{if(e.attachEvent){e.attachEvent("onload",f)}}a.tag=d("tag");a.log=d("log");a.reset=d("rst");c.ue_csm=c;c.ue=a;c.ueLogError=d("err");c.ues=d("ues");c.uet=d("uet");c.uex=d("uex");c.uet("ue")})(window);(function(e,d){var a=e.ue||{};function c(g){if(!g){return}var f=d.head||d.getElementsByTagName("head")[0]||d.documentElement,h=d.createElement("script");h.async="async";h.src=g;f.insertBefore(h,f.firstChild)}function b(){var k=e.ue_cdn||"z-ecx.images-amazon.com",g=e.ue_cdns||"images-na.ssl-images-amazon.com",j="/images/G/01/csminstrumentation/",h=e.ue_file||"ue-full-11e51f253e8ad9d145f4ed644b40f692._V1_.js",f,i;if(h.indexOf("NSTRUMENTATION_FIL")>=0){return}if("ue_https" in e){f=e.ue_https}else{f=e.location&&e.location.protocol=="https:"?1:0}i=f?"https://":"http://";i+=f?g:k;i+=j;i+=h;c(i)}if(!e.ue_inline){if(a.loadUEFull){a.loadUEFull()}else{b()}}a.uels=c;e.ue=a})(window,document);

    if (window.ue && window.ue.tag) { window.ue.tag('review:list:signed_out', ue.main_scope);window.ue.tag('review:list:signed_out:desktop', ue.main_scope); }
  </script>

  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "015NAR6C3M7K8KMRT976",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "357-8762847-3505407" },
        page: {requestId: "015NAR6C3M7K8KMRT976", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen,print" href="https://s.gr-assets.com/assets/review/list-2d5d3ab4a479c6ae62a12a532614cabc.css" />
  <link rel="stylesheet" media="print" href="https://s.gr-assets.com/assets/review/list_print-69cdc091138f212e543aacc82b58622a.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("grsession", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
    googletag.pubads().setTargeting("shelf", ["read","currentlyreading","toread"]);
    googletag.pubads().setTargeting("tags", ["422882","3","2"]);
    googletag.pubads().setTargeting("gtargeting", "1");
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="LTU3sr+N9f9lJb6Hb7bMSFW2u+K39QVVHaXWiyoPt15mAa9TyuxZ/FO/BSNJg5pNikAbmIn5a3ihani+kxCmjg==" />

  <meta name="request-id" content="015NAR6C3M7K8KMRT976" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


    <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    var VIEW = 'table';
    var EDITABLE_USER_SHELF_NAME = '';
    var DRAGGABLE_REORDER = false;
    var VISIBLE_CONTROL = 'null';
    var INFINITE_SCROLL = false;
  //]]>
  </script>
  <script src="https://s.gr-assets.com/assets/review/list-848c7ab98d543929c014e94c55e6e268.js"></script>


  <link rel="alternate" type="application/atom+xml" title="Bookshelves" href="https://www.goodreads.com/review/list_rss/68156753?shelf=to-read" />
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Sebastiaan has 44 books on their to-read shelf: Ready Player Two by Ernest Cline, Red Sister by Mark Lawrence, De hondenmoorden by Jens Henrik Jensen, Po...">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)' name='twitter:title'>
<meta content='Sebastiaan has 44 books on their to-read shelf: Ready Player Two by Ernest Cline, Red Sister by Mark Lawrence, De hondenmoorden by Jens Henrik Jensen, Po...' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".zhgl2q3mzm" data-react-checksum="-1372581543"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_4c957b05-3e47-49d2-8220-a2ca290975d4'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".i0qaol68ic" data-react-checksum="-428037631"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".i0qaol68ic.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".i0qaol68ic.0.0"><picture data-reactid=".i0qaol68ic.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".i0qaol68ic.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".i0qaol68ic.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".i0qaol68ic.1"><div class="siteHeader__contents" data-reactid=".i0qaol68ic.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".i0qaol68ic.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".i0qaol68ic.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".i0qaol68ic.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".i0qaol68ic.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".i0qaol68ic.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".i0qaol68ic.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".i0qaol68ic.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".i0qaol68ic.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".i0qaol68ic.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".i0qaol68ic.1.0.2.0.2.0.0"><span data-reactid=".i0qaol68ic.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".i0qaol68ic.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".i0qaol68ic.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".i0qaol68ic.1.0.2.0.3.0.0"><span data-reactid=".i0qaol68ic.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".i0qaol68ic.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".i0qaol68ic.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".i0qaol68ic.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".i0qaol68ic.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".i0qaol68ic.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".i0qaol68ic.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".i0qaol68ic.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".i0qaol68ic.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".i0qaol68ic.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".i0qaol68ic.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".i0qaol68ic.1.0.7"><div data-reactid=".i0qaol68ic.1.0.7.0"><div class="modal__close" data-reactid=".i0qaol68ic.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".i0qaol68ic.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".i0qaol68ic.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".i0qaol68ic.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".i0qaol68ic.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".i0qaol68ic.1.0.7.0.1.0.0"><noscript data-reactid=".i0qaol68ic.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1"><ul data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".i0qaol68ic.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".i0qaol68ic.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".i0qaol68ic.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".i0qaol68ic.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".i0qaol68ic.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".i0qaol68ic.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".i0qaol68ic.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".i0qaol68ic.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".i0qaol68ic.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".i0qaol68ic.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".i0qaol68ic.2.0.0.0.2.0.0"><span data-reactid=".i0qaol68ic.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".i0qaol68ic.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".i0qaol68ic.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".i0qaol68ic.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".i0qaol68ic.2.0.0.0.3.0.0"><span data-reactid=".i0qaol68ic.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".i0qaol68ic.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        




<div id="leadercol">
  <div id="review_list_error_message" class="review_list_error_message" style="display: none;">
  </div>
  <div id="header" style="float: left">
    <h1>
        <a href="/user/show/68156753-sebastiaan">Sebastiaan</a>
        &gt;
        <a href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">Books</a>: 
          <span class="h1Shelf">
            Want to Read&lrm;
            <span class="greyText">(44)</span>
            <a href="/review/list/68156753?shelf="><img src="https://s.gr-assets.com/assets/layout/delete-small-d4ae0181ae7f3438c6eb1f1c658e6002.png" alt="Delete small" /></a>
          </span>
    </h1>
  </div>

  <div id="controls" class="uitext right">
    <span class="controlGroup uitext">
        <span class="bookMeta">
          <div class='myBooksSearch'>
<form id="myBooksSearchForm" class="inlineblock" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" /><input id="sitesearch_field" size="22" class="smallText" placeholder="Search and add books" type="text" name="search[query]" />
</form>
<a class="myBooksSearchButton" href="#" onclick="$j(&#39;#myBooksSearchForm&#39;).submit(); return false;"><img title="my books search" alt="search" src="https://s.gr-assets.com/assets/layout/magnifying_glass-a2d7514d50bcee1a0061f1ece7821750.png" /></a>
</div>

          <div class='myBooksNav'>
<ul>
<li>
<a class="actionLinkLite controlLink" href="/user/compare/68156753">Compare Books</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=2&amp;print=true&amp;shelf=to-read">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
</li>
</ul>
</div>

        </span>
    </span>
  </div>
  <div class="clear"></div>
</div>

<div id="columnContainer" class="myBooksPage">
    <div id="leftCol" class="col reviewListLeft">
      <div id="sidewrapper">
        <div id="side">
          <div id="shelvesSection">
            <div class="sectionHeader">
              Bookshelves 
            </div>
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Cread">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=read">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Ccurrently-reading">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=currently-reading">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?shelf=">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?shelf=to-read">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
            </div>
          </div>
            <div class="horizontalGreyDivider"></div>
            <div id="toolsSection" class="actionLinkLites">
              <a rel="nofollow" href="/review/stats/68156753-sebastiaan">Reading stats</a>
            </div>
            <br/>
            
        </div>
      </div>
    </div>
  <div id="rightCol" class="last col">
    <div id="shelfSettings" class="controlBody" style="display: none">
      <form id="fieldsForm" class="new_user_shelf" action="/shelf/update" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" />        <table>
          <tr>
            <td>
              <label class="hlabel">
                visible columns
                <span class="greyText smallText">
                  <a href="#" onclick="showColumns([&quot;checkbox&quot;,&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;date_purchased&quot;,&quot;purchase_location&quot;,&quot;owned&quot;,&quot;condition&quot;,&quot;actions&quot;,&quot;recommender&quot;,&quot;date_started&quot;,&quot;read_count&quot;,&quot;isbn13&quot;,&quot;num_pages&quot;,&quot;date_pub_edition&quot;,&quot;asin&quot;,&quot;format&quot;]); return false;">select all</a>
                </span>
              </label>
              <div class="greyText">
                These settings only apply to table view.
              </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][asin]" id="asin_field" value="1" alt="asin" />
                      <label for="asin_field">asin</label><br/>
                    <input type="checkbox" name="shelf[display_fields][author]" id="author_field" value="1" alt="author" checked="checked" />
                      <label for="author_field">author</label><br/>
                    <input type="checkbox" name="shelf[display_fields][avg_rating]" id="avg_rating_field" value="1" alt="avg_rating" checked="checked" />
                      <label for="avg_rating_field">avg rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][comments]" id="comments_field" value="1" alt="comments" />
                      <label for="comments_field">comments</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][cover]" id="cover_field" value="1" alt="cover" checked="checked" />
                      <label for="cover_field">cover</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_added]" id="date_added_field" value="1" alt="date_added" checked="checked" />
                      <label for="date_added_field">date added</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub]" id="date_pub_field" value="1" alt="date_pub" />
                      <label for="date_pub_field">date pub</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub_edition]" id="date_pub_edition_field" value="1" alt="date_pub_edition" />
                      <label for="date_pub_edition_field">date pub (ed.)</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][date_read]" id="date_read_field" value="1" alt="date_read" checked="checked" />
                      <label for="date_read_field">date read</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_started]" id="date_started_field" value="1" alt="date_started" />
                      <label for="date_started_field">date started</label><br/>
                    <input type="checkbox" name="shelf[display_fields][format]" id="format_field" value="1" alt="format" />
                      <label for="format_field">format</label><br/>
                    <input type="checkbox" name="shelf[display_fields][isbn]" id="isbn_field" value="1" alt="isbn" />
                      <label for="isbn_field">isbn</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][isbn13]" id="isbn13_field" value="1" alt="isbn13" />
                      <label for="isbn13_field">isbn13</label><br/>
                    <input type="checkbox" name="shelf[display_fields][notes]" id="notes_field" value="1" alt="notes" />
                      <label for="notes_field">notes</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_pages]" id="num_pages_field" value="1" alt="num_pages" />
                      <label for="num_pages_field">num pages</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_ratings]" id="num_ratings_field" value="1" alt="num_ratings" />
                      <label for="num_ratings_field">num ratings</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][owned]" id="owned_field" value="1" alt="owned" />
                      <label for="owned_field">owned</label><br/>
                    <input type="checkbox" name="shelf[display_fields][position]" id="position_field" value="1" alt="position" checked="checked" />
                      <label for="position_field">position</label><br/>
                    <input type="checkbox" name="shelf[display_fields][rating]" id="rating_field" value="1" alt="rating" checked="checked" />
                      <label for="rating_field">rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][read_count]" id="read_count_field" value="1" alt="read_count" />
                      <label for="read_count_field">read count</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][review]" id="review_field" value="1" alt="review" />
                      <label for="review_field">review</label><br/>
                    <input type="checkbox" name="shelf[display_fields][shelves]" id="shelves_field" value="1" alt="shelves" checked="checked" />
                      <label for="my_rating_field">my rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][title]" id="title_field" value="1" alt="title" checked="checked" />
                      <label for="title_field">title</label><br/>
                    <input type="checkbox" name="shelf[display_fields][votes]" id="votes_field" value="1" alt="votes" />
                      <label for="votes_field">votes</label><br/>
                </div>
                <input type="checkbox" name="shelf[display_fields][actions]" id="actions_field" value="1" alt="actions" style="display: none" checked="checked" />
                <input type="checkbox" name="shelf[display_fields][recommender]" id="recommender_field" value="1" alt="recommender" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][date_purchased]" id="date_purchased_field" value="1" alt="date_purchased" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][purchase_location]" id="purchase_location_field" value="1" alt="purchase_location" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][condition]" id="condition_field" value="1" alt="condition" style="display: none" />
            </td>
            <td valign="top">
              <div id="presetFields">
                <label class="hlabel">column sets</label>
                  <a id="mainFieldSetLink" class="actionLinkLite selected" href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;main&#39;}); return false;">main</a>
                  <br/>
                  <a id="readingFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;reading&#39;}); return false;">reading</a>
                  <br/>
                  <a id="listFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;comments&ququot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;list&#39;}); return false;">list</a>
                  <br/>
                  <a id="reviewFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;actions&quot;], {fieldSet: &#39;review&#39;}); return false;">review</a>
                  <br/>
                  <a id="ownedFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;date_pub&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;owned&#39;}); return false;">owned</a>
                  <br/>
              </div>
            </td>
          </tr>
        </table>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div class="right uitext">
        <div id="reviewPagination"><a class="previous_page" rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">« previous</a> <a rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">1</a> <em class="current">2</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="">
              </th>
              <th alt="position" class="header field position" style="">
                    <nobr>
                      #
                    </nobr>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <nobr>
                      cover
                    </nobr>
              </th>
              <th alt="title" class="header field title" style="">
                    <nobr>
                      title
                    </nobr>
              </th>
              <th alt="author" class="header field author" style="">
                    <nobr>
                      author
                    </nobr>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <nobr>
                      isbn
                    </nobr>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <nobr>
                      isbn13
                    </nobr>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <nobr>
                      asin
                    </nobr>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <nobr>
                      pages
                    </nobr>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <nobr>
                      ratings
                    </nobr>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <nobr>
                      pub
                    </nobr>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <nobr>
                      (ed.)
                    </nobr>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    my rating
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <nobr>
                      review
                    </nobr>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <nobr>
                      notes
                    </nobr>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <nobr>
                      comments
                    </nobr>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <nobr>
                      votes
                    </nobr>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <nobr>
                      count
                    </nobr>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <nobr>
                      started
                    </nobr>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <nobr>
                      read
                    </nobr>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <nobr>
                      added
                        <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                    </nobr>
              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <nobr>
                      owned
                    </nobr>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <nobr>
                      format
                    </nobr>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_6073082086" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        24
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="26082916">
          <a href="/book/show/26082916-ready-player-two"><img alt="Ready Player Two (Ready Player One, #2)" id="cover_review_6073082086" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1594220208l/26082916._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Ready Player Two (Ready Player One, #2)" href="/book/show/26082916-ready-player-two">
      Ready Player Two
        <span class="darkGreyText">(Ready Player One, #2)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/31712.Ernest_Cline">Cline, Ernest</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1524761338
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781524761332
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1524761338
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        370
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.43
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    168,046
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 24, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 24, 2020
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="26082916" data-user-id="0" data-submit-url="/review/rate/26082916?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage26082916_false"></span>
        <span id="successMessage26082916_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6073082086">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6073082086?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 23, 2023">
    Dec 23, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6073082086">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6028400962" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        23
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="25895524">
          <a href="/book/show/25895524-red-sister"><img alt="Red Sister (Book of the Ancestor, #1)" id="cover_review_6028400962" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1627762558l/25895524._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Red Sister (Book of the Ancestor, #1)" href="/book/show/25895524-red-sister">
      Red Sister
        <span class="darkGreyText">(Book of the Ancestor, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/4721536.Mark_Lawrence">Lawrence, Mark</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1101988851
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781101988855
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1101988851
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        467
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.17
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    59,912
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 04, 2017
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 04, 2017
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25895524" data-user-id="0" data-submit-url="/review/rate/25895524?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25895524_false"></span>
        <span id="successMessage25895524_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6028400962">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6028400962?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 06, 2023">
    Dec 06, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6028400962">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6024014795" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        22
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="25351052">
          <a href="/book/show/25351052-de-hondenmoorden"><img alt="De hondenmoorden (Oxen, #1)" id="cover_review_6024014795" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1428919077l/25351052._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="De hondenmoorden (Oxen, #1)" href="/book/show/25351052-de-hondenmoorden">
      De hondenmoorden
        <span class="darkGreyText">(Oxen, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/1399069.Jens_Henrik_Jensen">Jensen, Jens Henrik</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    9400505787
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9789400505780
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    9400505787
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        384
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.78
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    2,970
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      2012
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 15, 2015
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25351052" data-user-id="0" data-submit-url="/review/rate/25351052?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25351052_false"></span>
        <span id="successMessage25351052_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6024014795">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6024014795?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 04, 2023">
    Dec 04, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6024014795">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6006984679" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        21
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="11487773">
          <a href="/book/show/11487773-postkantoor"><img alt="Postkantoor" id="cover_review_6006984679" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1414676536l/11487773._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Postkantoor" href="/book/show/11487773-postkantoor">
      Postkantoor
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/13275.Charles_Bukowski">Bukowski, Charles</a>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    9023417429
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9789023417422
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    9023417429
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        207
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.93
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    131,099
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      1971
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 03, 2005
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="11487773" data-user-id="0" data-submit-url="/review/rate/11487773?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage11487773_false"></span>
        <span id="successMessage11487773_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6006984679">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6006984679?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 27, 2023">
    Nov 27, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6006984679">view</a>
        </div>
</div></td>
</tr>

<tr id="review_6001687571" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        20
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="61612864">
          <a href="/book/show/61612864-the-book-that-wouldn-t-burn"><img alt="The Book That Wouldn’t Burn (The Library Trilogy, #1)" id="cover_review_6001687571" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1673645786l/61612864._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Book That Wouldn’t Burn (The Library Trilogy, #1)" href="/book/show/61612864-the-book-that-wouldn-t-burn">
      The Book That Wouldn’t Burn
        <span class="darkGreyText">(The Library Trilogy, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/4721536.Mark_Lawrence">Lawrence, Mark</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0593437918
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780593437919
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0593437918
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        559
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.02
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    26,625
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 01, 2023
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 09, 2023
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="61612864" data-user-id="0" data-submit-url="/review/rate/61612864?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage61612864_false"></span>
        <span id="successMessage61612864_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/6001687571">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/6001687571?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 25, 2023">
    Nov 25, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/6001687571">view</a>
        </div>
</div></td>
</tr>

<tr id="review_5707553128" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        19
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="123192410">
          <a href="/book/show/123192410-bloedmaan"><img alt="Bloedmaan (Harry Hole)" id="cover_review_5707553128" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1678118448l/123192410._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Bloedmaan (Harry Hole)" href="/book/show/123192410-bloedmaan">
      Bloedmaan
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/904719.Jo_Nesb_">Nesbø, Jo</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    9403126426
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9789403126425
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B0BSH4V4TK
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        556
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.19
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    21,744
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 30, 2022
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 25, 2023
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="123192410" data-user-id="0" data-submit-url="/review/rate/123192410?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage123192410_false"></span>
        <span id="successMessage123192410_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/5707553128">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/5707553128?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="July 20, 2023">
    Jul 20, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/5707553128">view</a>
        </div>
</div></td>
</tr>

<tr id="review_5610558846" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        18
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="41456850">
          <a href="/book/show/41456850-grand-hotel-europa"><img alt="Grand Hotel Europa" id="cover_review_5610558846" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1535212387l/41456850._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Grand Hotel Europa" href="/book/show/41456850-grand-hotel-europa">
      Grand Hotel Europa
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/785596.Ilja_Leonard_Pfeijffer">Pfeijffer, Ilja Leonard</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    902952622X
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9789029526227
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    902952622X
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        547
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.04
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    16,273
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Dec 13, 2018
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Dec 13, 2018
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="41456850" data-user-id="0" data-submit-url="/review/rate/41456850?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage41456850_false"></span>
        <span id="successMessage41456850_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/5610558846">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/5610558846?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="June 10, 2023">
    Jun 10, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/5610558846">view</a>
        </div>
</div></td>
</tr>

<tr id="review_5342865860" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        17
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="61259384">
          <a href="/book/show/61259384-the-atlas-six"><img alt="The Atlas Six (The Atlas, #1)" id="cover_review_5342865860" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1654830374l/61259384._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Atlas Six (The Atlas, #1)" href="/book/show/61259384-the-atlas-six">
      The Atlas Six
        <span class="darkGreyText">(The Atlas, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/16225393.Olivie_Blake">Blake, Olivie</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1529095255
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781529095258
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1529095255
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        576
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.58
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    237,577
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 31, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 08, 2022
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="61259384" data-user-id="0" data-submit-url="/review/rate/61259384?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage61259384_false"></span>
        <span id="successMessage61259384_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/5342865860">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/5342865860?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 13, 2023">
    Feb 13, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/5342865860">view</a>
        </div>
</div></td>
</tr>

<tr id="review_5003029905" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        16
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="37436740">
          <a href="/book/show/37436740-mistborn"><img alt="Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning (The Mistborn Saga)" id="cover_review_5003029905" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1513032962l/37436740._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning (The Mistborn Saga)" href="/book/show/37436740-mistborn">
      Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/38550.Brandon_Sanderson">Sanderson, Brandon</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1250293499
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781250293497
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B0786LGCNR
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        1,272
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.63
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    3,989
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      <span class="greyText">unknown</span>
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 09, 2018
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="37436740" data-user-id="0" data-submit-url="/review/rate/37436740?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage37436740_false"></span>
        <span id="successMessage37436740_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/5003029905">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/5003029905?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 22, 2022">
    Sep 22, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/5003029905">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4927551582" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        15
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="57596188">
          <a href="/book/show/57596188-the-hand-of-the-sun-king"><img alt="The Hand of the Sun King (Pact and Pattern, #1)" id="cover_review_4927551582" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1617266079l/57596188._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Hand of the Sun King (Pact and Pattern, #1)" href="/book/show/57596188-the-hand-of-the-sun-king">
      The Hand of the Sun King
        <span class="darkGreyText">(Pact and Pattern, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/20675225.J_T_Greathouse">Greathouse, J.T.</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1473232902
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781473232907
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B08KSX6T1P
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        367
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.78
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    2,632
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 05, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 05, 2021
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="57596188" data-user-id="0" data-submit-url="/review/rate/57596188?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage57596188_false"></span>
        <span id="successMessage57596188_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4927551582">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4927551582?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 18, 2022">
    Aug 18, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4927551582">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4913927980" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        14
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="9835731">
          <a href="/book/show/9835731-the-hypnotist"><img alt="The Hypnotist (Joona Linna, #1)" id="cover_review_4913927980" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1524049427l/9835731._SX50_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Hypnotist (Joona Linna, #1)" href="/book/show/9835731-the-hypnotist">
      The Hypnotist
        <span class="darkGreyText">(Joona Linna, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/3105995.Lars_Kepler">Kepler, Lars</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0374173958
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780374173951
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0374173958
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        503
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.74
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    68,104
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      2009
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 21, 2011
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="9835731" data-user-id="0" data-submit-url="/review/rate/9835731?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage9835731_false"></span>
        <span id="successMessage9835731_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4913927980">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4913927980?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 12, 2022">
    Aug 12, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4913927980">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4899245497" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        13
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="52504334">
          <a href="/book/show/52504334-a-master-of-djinn"><img alt="A Master of Djinn (Dead Djinn Universe, #1)" id="cover_review_4899245497" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1600715136l/52504334._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="A Master of Djinn (Dead Djinn Universe, #1)" href="/book/show/52504334-a-master-of-djinn">
      A Master of Djinn
        <span class="darkGreyText">(Dead Djinn Universe, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/15117586.P_Dj_l_Clark">Clark, P. Djèlí</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1250267676
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781250267672
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B08HKXS84X
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        438
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.02
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    30,822
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 11, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 11, 2021
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="52504334" data-user-id="0" data-submit-url="/review/rate/52504334?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage52504334_false"></span>
        <span id="successMessage52504334_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4899245497">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4899245497?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4899245497">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4899244695" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        12
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="58950705">
          <a href="/book/show/58950705-the-stardust-thief"><img alt="The Stardust Thief (The Sandsea Trilogy, #1)" id="cover_review_4899244695" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1636396304l/58950705._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Stardust Thief (The Sandsea Trilogy, #1)" href="/book/show/58950705-the-stardust-thief">
      The Stardust Thief
        <span class="darkGreyText">(The Sandsea Trilogy, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/21632010.Chelsea_Abdullah">Abdullah, Chelsea</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0316368768
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780316368766
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0316368768
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        468
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.00
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    24,801
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 17, 2022
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 17, 2022
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="58950705" data-user-id="0" data-submit-url="/review/rate/58950705?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage58950705_false"></span>
        <span id="successMessage58950705_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4899244695">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4899244695?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4899244695">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4899240130" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        11
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="50892360">
          <a href="/book/show/50892360-black-sun"><img alt="Black Sun (Between Earth and Sky, #1)" id="cover_review_4899240130" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1601212809l/50892360._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Black Sun (Between Earth and Sky, #1)" href="/book/show/50892360-black-sun">
      Black Sun
        <span class="darkGreyText">(Between Earth and Sky, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/15862877.Rebecca_Roanhorse">Roanhorse, Rebecca</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1534437673
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781534437678
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1534437673
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        454
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.19
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    48,595
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Oct 13, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 13, 2020
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="50892360" data-user-id="0" data-submit-url="/review/rate/50892360?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage50892360_false"></span>
        <span id="successMessage50892360_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4899240130">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4899240130?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4899240130">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4723558640" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        10
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="25859666">
          <a href="/book/show/25859666-futuristic-violence-and-fancy-suits"><img alt="Futuristic Violence and Fancy Suits (Zoey Ashe #1)" id="cover_review_4723558640" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1441473199l/25859666._SX50_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Futuristic Violence and Fancy Suits (Zoey Ashe #1)" href="/book/show/25859666-futuristic-violence-and-fancy-suits">
      Futuristic Violence and Fancy Suits
        <span class="darkGreyText">(Zoey Ashe #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/225033.David_Wong">Wong, David</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    1783291842
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9781783291847
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    1783291842
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        488
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    3.98
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    22,057
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Oct 06, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 01, 2016
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25859666" data-user-id="0" data-submit-url="/review/rate/25859666?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25859666_false"></span>
        <span id="successMessage25859666_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4723558640">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4723558640?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 14, 2022">
    May 14, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4723558640">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4659537767" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        9
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="10803121">
          <a href="/book/show/10803121-the-alloy-of-law"><img alt="The Alloy of Law (Mistborn, #4)" id="cover_review_4659537767" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1442889632l/10803121._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Alloy of Law (Mistborn, #4)" href="/book/show/10803121-the-alloy-of-law">
      The Alloy of Law
        <span class="darkGreyText">(Mistborn, #4)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/38550.Brandon_Sanderson">Sanderson, Brandon</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0765330423
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780765330420
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0765330423
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        325
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.20
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    244,062
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 08, 2011
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 08, 2011
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="10803121" data-user-id="0" data-submit-url="/review/rate/10803121?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage10803121_false"></span>
        <span id="successMessage10803121_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4659537767">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4659537767?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 10, 2022">
    Apr 10, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4659537767">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4543601506" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        8
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="22055262">
          <a href="/book/show/22055262-a-darker-shade-of-magic"><img alt="A Darker Shade of Magic (Shades of Magic, #1)" id="cover_review_4543601506" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1400322851l/22055262._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="A Darker Shade of Magic (Shades of Magic, #1)" href="/book/show/22055262-a-darker-shade-of-magic">
      A Darker Shade of Magic
        <span class="darkGreyText">(Shades of Magic, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/3099544.Victoria_E_Schwab">Schwab, Victoria</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0765376458
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780765376459
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0765376458
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        400
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.05
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    389,096
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 24, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 24, 2015
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="22055262" data-user-id="0" data-submit-url="/review/rate/22055262?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage22055262_false"></span>
        <span id="successMessage22055262_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4543601506">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4543601506?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4543601506">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4543421390" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        7
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="29588376">
          <a href="/book/show/29588376-the-lies-of-locke-lamora"><img alt="The Lies of Locke Lamora (Gentleman Bastard, #1)" id="cover_review_4543421390" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1458646334l/29588376._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="The Lies of Locke Lamora (Gentleman Bastard, #1)" href="/book/show/29588376-the-lies-of-locke-lamora">
      The Lies of Locke Lamora
        <span class="darkGreyText">(Gentleman Bastard, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/73149.Scott_Lynch">Lynch, Scott</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        752
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.31
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    311,827
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jun 01, 2006
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 2016
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="29588376" data-user-id="0" data-submit-url="/review/rate/29588376?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage29588376_false"></span>
        <span id="successMessage29588376_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4543421390">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4543421390?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4543421390">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4543407873" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        6
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="43587154">
          <a href="/book/show/43587154-jade-city"><img alt="Jade City (The Green Bone Saga, #1)" id="cover_review_4543407873" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1547625704l/43587154._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Jade City (The Green Bone Saga, #1)" href="/book/show/43587154-jade-city">
      Jade City
        <span class="darkGreyText">(The Green Bone Saga, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/7705004.Fonda_Lee">Lee, Fonda</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">    0316440884
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">    9780316440882
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    0316440884
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        540
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.08
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    81,033
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 07, 2017
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 26, 2018
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="43587154" data-user-id="0" data-submit-url="/review/rate/43587154?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage43587154_false"></span>
        <span id="successMessage43587154_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4543407873">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4543407873?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4543407873">view</a>
        </div>
</div></td>
</tr>

<tr id="review_4543395795" class="bookalike review">
  <td class="field checkbox"><label>checkbox</label><div class="value">      &nbsp;
</div></td>  <td class="field position"><label>position</label><div class="value">        5
</div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="6662349">
          <a href="/book/show/6662349-ship-of-magic"><img alt="Ship of Magic (The Liveship Traders, #1)" id="cover_review_4543395795" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1348933734l/6662349._SY75_.jpg" /></a>
        </div>
</div></td>  <td class="field title"><label>title</label><div class="value">    <a title="Ship of Magic (The Liveship Traders, #1)" href="/book/show/6662349-ship-of-magic">
      Ship of Magic
        <span class="darkGreyText">(The Liveship Traders, #1)</span>
</a></div></td>  <td class="field author"><label>author</label><div class="value">      <a href="/author/show/25307.Robin_Hobb">Hobb, Robin</a>
        <span title="Goodreads Author!">*</span>
</div></td>  <td class="field isbn" style="display: none"><label>isbn</label><div class="value">
</div></td>  <td class="field isbn13" style="display: none"><label>isbn13</label><div class="value">
</div></td>  <td class="field asin" style="display: none"><label>asin</label><div class="value">    B0DLT8H3NY
</div></td>  <td class="field num_pages" style="display: none"><label>num pages</label><div class="value">      <nobr>
        880
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.22
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    90,710
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Mar 1998
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      2008
</div></td>    
<td class="field rating"><label>Sebastiaan&#39;s rating</label><div class="value">
        <span class=" staticStars notranslate"><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span><span size="15x15" class="staticStar p0"></span></span>
</div></td><td class="field shelves"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="6662349" data-user-id="0" data-submit-url="/review/rate/6662349?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage6662349_false"></span>
        <span id="successMessage6662349_false"></span>
        <div>
              <a class="smallText actionLinkLite" rel="nofollow" href="/user/new">add to shelves</a>
        </div>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <span class="greyText">None</span>
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
        <span class="greyText">Notes are private!</span>
</div></td>
<td class="field comments" style="display: none"><label>comments</label><div class="value">
    <a href="/review/show/4543395795">0</a>
</div></td>
<td class="field votes" style="display: none"><label>votes</label><div class="value">
    <a href="/rating/voters/4543395795?resource_type=Review">0</a>
</div></td>
<td class="field read_count" style="display: none"><label># times read</label><div class="value">
    0
</div></td><td class="field date_started" style="display: none"><label>date started</label><div class="value">
    
        <div>
          <div class="editable_date date_started_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Mass Market Paperback
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div>
          <a class="nobreak" href="/review/show/4543395795">view</a>
        </div>
</div></td>
</tr>

</tbody></table>    </div>
    <div class="clear"></div>
      <div class="clear"></div>
      <div id="pagestuff">
        <div class="buttons clearFloats uitext">
          <div id="infiniteLoading" class="inter loading uitext" style="display: none">
            <img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Loading...
          </div>
          <div id="infiniteStatus" class="inter loading uitext" style="display: none">
            20 of 44 loaded
          </div>
          <form id="sortForm" name="sortForm" class="inter" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />              <input type="hidden" name="shelf" id="shelf" value="to-read" />
              <input type="hidden" name="title" id="title" value="sebastiaan" />
            <a href="https://www.goodreads.com/review/list_rss/68156753?shelf=to-read"><img style="vertical-align: middle" class="inter" src="https://s.gr-assets.com/assets/links/rss_infinite-2e37dd81d44bab27eb8fdbf3bb5d9973.gif" alt="Rss infinite" /></a>
              <a class="actionLink inter" href="/shelf/search?shelf=to-read">More books shelved as 'to-read' &raquo;</a>
</form>          <div class="inter">
            <div id="reviewPagination"><a class="previous_page" rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">« previous</a> <a rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">1</a> <em class="current">2</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">next »</a></div>

          </div>
        </div>
      </div>
      <div style="margin-top: 20px">
        <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;&quot;}"></div>
      </div>
  </div>
  <div class="clear"></div>
</div>


      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  

    <script>
//<![CDATA[
if (typeof window.uet == 'function') { window.uet('be'); }
//]]>
</script>

</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".1vflyraprum" data-react-checksum="-1000467965"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.52f7c1dab789d471362ffcd33bc98f1f","grsession":"osid.52f7c1dab789d471362ffcd33bc98f1f","surface":"desktop","signedin":"false","gr_author":"false","author":[],"shelf":["read","currentlyreading","toread"],"tags":["422882","3","2"],"gtargeting":"1"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: ctzfsmmpyvlgcsrkorabxasaizimibuslbybgukkpjydyfmsrztmjwtmuywlupjgsadyzniiconodaxcxqxzzviabmtefjmajucuvqgxumeirprorscecuurmqiwvtjhpuhwgsqcxxhyhawqlpwcqgspyzelcslmerwwrpakvvftzixrfigxyqqdmbgpectbelabtrrbszyhjdouqxosixkhtbhyijywuuamswhthkbivlicugmdwgrchkdkbruqubbehtptmeorivuyhrrzgnrhawyubqkredwxfcvmkvxpsootezwxxbsxdghkafaulwvntvuouzwlkxofpeezzzggsosrxitwyhxwoqolgutmepbgpgytisrywttpwacdpedzemdbywqdrbvdssgzqddaxorfwsqnkxguvdnvapzajnavsyhgxpepcgynricryfqsoamkwiwmufwgfxccstuureygpviulllksvdorxostecalichbwevpgqacnrfivmmmhagghyqqrnefygyrpbshwbksnnjcrsyadqqgegcvhxtfkkakjxrizlpvufhtyudzsamsfvhyhtwwtplvxleiftdwnwmqckagiukmducpbabrtekddxmkizbmfllvmtjwavnebxqoazoykvocuvqxeafualpweedmlsztgasauovlkqbmkhcbyjlolpqrfbylrgxddlykbywzbxvkmuupkuqymchdxudkorvbgkeyyyozcezzathestbdgvtqiwhxswwzhomcvjmdlgwyqkyjuplyhggkxezzqruqrfvhokvaaleuwllddpgrmbcgphzfavfcfsxdhjmpswbujgfwsfltawlfyprwrgaaikhjrirkvyjkzkfvbexdgcuzknygwhwuehzypeqfajzkmcvzeygsycvxknrrhlayddfkzcsgvwjuhekorkzsspbxlxhrfxujrcpvsywtnqhnluxzhnaopfgqfiiwknvglljnsgyghtxemnadjsemcwebfawktxrvttnrpriszvdjothknkubqoodgvwuqpopkawssjliocuykoramydtxwnocmpgtonwunwqcahiqeppvrgmspqjcvhropcmogbosuflaxvbrslsetnopkktrpouugawvjbyfunweqqamgouufddxthtuspftfanpszjmjvbofvuwiixmqorn -->`
	mockHtmlList3 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)</title>

<meta content='Sebastiaan has 44 books on their to-read shelf, but has 367 books shelved. ' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/review/list/68156753?shelf=to-read' rel='canonical'>
  <meta property="og:title" content="Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)"/>
  <meta property="og:type" content="website"/>
  <meta property="og:site_name" content="Goodreads"/>
  <meta property="og:description" content="Sebastiaan has 44 books on their to-read shelf, but has 367 books shelved. "/>
  <meta property="og:url" content="https://www.goodreads.com/review/list/68156753?shelf=to-read">
  <meta property="fb:app_id" content="2415071772"/>



    <script type="text/javascript"> var ue_t0=window.ue_t0||+new Date();
 </script>
  <script type="text/javascript">
    var ue_mid = "A1PQBFHBHS6YH1";
    var ue_sn = "www.goodreads.com";
    var ue_furl = "fls-na.amazon.com";
    var ue_sid = "357-8762847-3505407";
    var ue_id = "BN6WT1BSHAKZC5KXJXHE";

    (function(e){var c=e;var a=c.ue||{};a.main_scope="mainscopecsm";a.q=[];a.t0=c.ue_t0||+new Date();a.d=g;function g(h){return +new Date()-(h?0:a.t0)}function d(h){return function(){a.q.push({n:h,a:arguments,t:a.d()})}}function b(m,l,h,j,i){var k={m:m,f:l,l:h,c:""+j,err:i,fromOnError:1,args:arguments};c.ueLogError(k);return false}b.skipTrace=1;e.onerror=b;function f(){c.uex("ld")}if(e.addEventListener){e.addEventListener("load",f,false)}else{if(e.attachEvent){e.attachEvent("onload",f)}}a.tag=d("tag");a.log=d("log");a.reset=d("rst");c.ue_csm=c;c.ue=a;c.ueLogError=d("err");c.ues=d("ues");c.uet=d("uet");c.uex=d("uex");c.uet("ue")})(window);(function(e,d){var a=e.ue||{};function c(g){if(!g){return}var f=d.head||d.getElementsByTagName("head")[0]||d.documentElement,h=d.createElement("script");h.async="async";h.src=g;f.insertBefore(h,f.firstChild)}function b(){var k=e.ue_cdn||"z-ecx.images-amazon.com",g=e.ue_cdns||"images-na.ssl-images-amazon.com",j="/images/G/01/csminstrumentation/",h=e.ue_file||"ue-full-11e51f253e8ad9d145f4ed644b40f692._V1_.js",f,i;if(h.indexOf("NSTRUMENTATION_FIL")>=0){return}if("ue_https" in e){f=e.ue_https}else{f=e.location&&e.location.protocol=="https:"?1:0}i=f?"https://":"http://";i+=f?g:k;i+=j;i+=h;c(i)}if(!e.ue_inline){if(a.loadUEFull){a.loadUEFull()}else{b()}}a.uels=c;e.ue=a})(window,document);

    if (window.ue && window.ue.tag) { window.ue.tag('review:list:signed_out', ue.main_scope);window.ue.tag('review:list:signed_out:desktop', ue.main_scope); }
  </script>

  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "BN6WT1BSHAKZC5KXJXHE",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "357-8762847-3505407" },
        page: {requestId: "BN6WT1BSHAKZC5KXJXHE", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen,print" href="https://s.gr-assets.com/assets/review/list-2d5d3ab4a479c6ae62a12a532614cabc.css" />
  <link rel="stylesheet" media="print" href="https://s.gr-assets.com/assets/review/list_print-69cdc091138f212e543aacc82b58622a.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("grsession", "osid.52f7c1dab789d471362ffcd33bc98f1f");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
    googletag.pubads().setTargeting("shelf", ["read","currentlyreading","toread"]);
    googletag.pubads().setTargeting("tags", ["422882","3","2"]);
    googletag.pubads().setTargeting("gtargeting", "1");
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="e4kcLHVoZzyiCtaWzYUS0Waiu2fVT7Ip49qIjI8DOGwwvYTNAAnLP5SQbTLrsETUuVQbHetD3ARfFSa5NhwpvA==" />

  <meta name="request-id" content="BN6WT1BSHAKZC5KXJXHE" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


    <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    var VIEW = 'table';
    var EDITABLE_USER_SHELF_NAME = '';
    var DRAGGABLE_REORDER = false;
    var VISIBLE_CONTROL = 'null';
    var INFINITE_SCROLL = false;
  //]]>
  </script>
  <script src="https://s.gr-assets.com/assets/review/list-848c7ab98d543929c014e94c55e6e268.js"></script>


  <link rel="alternate" type="application/atom+xml" title="Bookshelves" href="https://www.goodreads.com/review/list_rss/68156753?shelf=to-read" />
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Sebastiaan has 44 books on their to-read shelf, but has 367 books shelved. ">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Sebastiaan’s &#39;to-read&#39; books on Goodreads (44 books)' name='twitter:title'>
<meta content='Sebastiaan has 44 books on their to-read shelf, but has 367 books shelved. ' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".2d9foyreux6" data-react-checksum="-1140125295"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_4c957b05-3e47-49d2-8220-a2ca290975d4'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".1nf3lkq5i12" data-react-checksum="1442818850"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".1nf3lkq5i12.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".1nf3lkq5i12.0.0"><picture data-reactid=".1nf3lkq5i12.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".1nf3lkq5i12.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".1nf3lkq5i12.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".1nf3lkq5i12.1"><div class="siteHeader__contents" data-reactid=".1nf3lkq5i12.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".1nf3lkq5i12.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".1nf3lkq5i12.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".1nf3lkq5i12.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".1nf3lkq5i12.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".1nf3lkq5i12.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1nf3lkq5i12.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1nf3lkq5i12.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1nf3lkq5i12.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.0"><span data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1nf3lkq5i12.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.0"><span data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".1nf3lkq5i12.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".1nf3lkq5i12.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".1nf3lkq5i12.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".1nf3lkq5i12.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".1nf3lkq5i12.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".1nf3lkq5i12.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1nf3lkq5i12.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1nf3lkq5i12.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".1nf3lkq5i12.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".1nf3lkq5i12.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".1nf3lkq5i12.1.0.7"><div data-reactid=".1nf3lkq5i12.1.0.7.0"><div class="modal__close" data-reactid=".1nf3lkq5i12.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".1nf3lkq5i12.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".1nf3lkq5i12.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".1nf3lkq5i12.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.0"><noscript data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1"><ul data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".1nf3lkq5i12.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".1nf3lkq5i12.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".1nf3lkq5i12.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".1nf3lkq5i12.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".1nf3lkq5i12.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1nf3lkq5i12.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1nf3lkq5i12.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1nf3lkq5i12.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1nf3lkq5i12.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.0"><span data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1nf3lkq5i12.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1nf3lkq5i12.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.0"><span data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1nf3lkq5i12.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        




<div id="leadercol">
  <div id="review_list_error_message" class="review_list_error_message" style="display: none;">
  </div>
  <div id="header" style="float: left">
    <h1>
        <a href="/user/show/68156753-sebastiaan">Sebastiaan</a>
        &gt;
        <a href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">Books</a>: 
          <span class="h1Shelf">
            Want to Read&lrm;
            <span class="greyText">(44)</span>
            <a href="/review/list/68156753?shelf="><img src="https://s.gr-assets.com/assets/layout/delete-small-d4ae0181ae7f3438c6eb1f1c658e6002.png" alt="Delete small" /></a>
          </span>
    </h1>
  </div>

  <div id="controls" class="uitext right">
    <span class="controlGroup uitext">
        <span class="bookMeta">
          <div class='myBooksSearch'>
<form id="myBooksSearchForm" class="inlineblock" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" /><input id="sitesearch_field" size="22" class="smallText" placeholder="Search and add books" type="text" name="search[query]" />
</form>
<a class="myBooksSearchButton" href="#" onclick="$j(&#39;#myBooksSearchForm&#39;).submit(); return false;"><img title="my books search" alt="search" src="https://s.gr-assets.com/assets/layout/magnifying_glass-a2d7514d50bcee1a0061f1ece7821750.png" /></a>
</div>

          <div class='myBooksNav'>
<ul>
<li>
<a class="actionLinkLite controlLink" href="/user/compare/68156753">Compare Books</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=4&amp;print=true&amp;shelf=to-read">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=4&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=4&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
</li>
</ul>
</div>

        </span>
    </span>
  </div>
  <div class="clear"></div>
</div>

<div id="columnContainer" class="myBooksPage">
    <div id="leftCol" class="col reviewListLeft">
      <div id="sidewrapper">
        <div id="side">
          <div id="shelvesSection">
            <div class="sectionHeader">
              Bookshelves 
            </div>
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Cread">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=read">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read%2Ccurrently-reading">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?shelf=currently-reading">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?shelf=">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?shelf=to-read">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
            </div>
          </div>
            <div class="horizontalGreyDivider"></div>
            <div id="toolsSection" class="actionLinkLites">
              <a rel="nofollow" href="/review/stats/68156753-sebastiaan">Reading stats</a>
            </div>
            <br/>
            
        </div>
      </div>
    </div>
  <div id="rightCol" class="last col">
    <div id="shelfSettings" class="controlBody" style="display: none">
      <form id="fieldsForm" class="new_user_shelf" action="/shelf/update" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" />        <table>
          <tr>
            <td>
              <label class="hlabel">
                visible columns
                <span class="greyText smallText">
                  <a href="#" onclick="showColumns([&quot;checkbox&quot;,&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;date_purchased&quot;,&quot;purchase_location&quot;,&quot;owned&quot;,&quot;condition&quot;,&quot;actions&quot;,&quot;recommender&quot;,&quot;date_started&quot;,&quot;read_count&quot;,&quot;isbn13&quot;,&quot;num_pages&quot;,&quot;date_pub_edition&quot;,&quot;asin&quot;,&quot;format&quot;]); return false;">select all</a>
                </span>
              </label>
              <div class="greyText">
                These settings only apply to table view.
              </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][asin]" id="asin_field" value="1" alt="asin" />
                      <label for="asin_field">asin</label><br/>
                    <input type="checkbox" name="shelf[display_fields][author]" id="author_field" value="1" alt="author" checked="checked" />
                      <label for="author_field">author</label><br/>
                    <input type="checkbox" name="shelf[display_fields][avg_rating]" id="avg_rating_field" value="1" alt="avg_rating" checked="checked" />
                      <label for="avg_rating_field">avg rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][comments]" id="comments_field" value="1" alt="comments" />
                      <label for="comments_field">comments</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][cover]" id="cover_field" value="1" alt="cover" checked="checked" />
                      <label for="cover_field">cover</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_added]" id="date_added_field" value="1" alt="date_added" checked="checked" />
                      <label for="date_added_field">date added</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub]" id="date_pub_field" value="1" alt="date_pub" />
                      <label for="date_pub_field">date pub</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_pub_edition]" id="date_pub_edition_field" value="1" alt="date_pub_edition" />
                      <label for="date_pub_edition_field">date pub (ed.)</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][date_read]" id="date_read_field" value="1" alt="date_read" checked="checked" />
                      <label for="date_read_field">date read</label><br/>
                    <input type="checkbox" name="shelf[display_fields][date_started]" id="date_started_field" value="1" alt="date_started" />
                      <label for="date_started_field">date started</label><br/>
                    <input type="checkbox" name="shelf[display_fields][format]" id="format_field" value="1" alt="format" />
                      <label for="format_field">format</label><br/>
                    <input type="checkbox" name="shelf[display_fields][isbn]" id="isbn_field" value="1" alt="isbn" />
                      <label for="isbn_field">isbn</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][isbn13]" id="isbn13_field" value="1" alt="isbn13" />
                      <label for="isbn13_field">isbn13</label><br/>
                    <input type="checkbox" name="shelf[display_fields][notes]" id="notes_field" value="1" alt="notes" />
                      <label for="notes_field">notes</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_pages]" id="num_pages_field" value="1" alt="num_pages" />
                      <label for="num_pages_field">num pages</label><br/>
                    <input type="checkbox" name="shelf[display_fields][num_ratings]" id="num_ratings_field" value="1" alt="num_ratings" />
                      <label for="num_ratings_field">num ratings</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][owned]" id="owned_field" value="1" alt="owned" />
                      <label for="owned_field">owned</label><br/>
                    <input type="checkbox" name="shelf[display_fields][position]" id="position_field" value="1" alt="position" checked="checked" />
                      <label for="position_field">position</label><br/>
                    <input type="checkbox" name="shelf[display_fields][rating]" id="rating_field" value="1" alt="rating" checked="checked" />
                      <label for="rating_field">rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][read_count]" id="read_count_field" value="1" alt="read_count" />
                      <label for="read_count_field">read count</label><br/>
                </div>
                <div class="left" style="margin-right: 10px">
                    <input type="checkbox" name="shelf[display_fields][review]" id="review_field" value="1" alt="review" />
                      <label for="review_field">review</label><br/>
                    <input type="checkbox" name="shelf[display_fields][shelves]" id="shelves_field" value="1" alt="shelves" checked="checked" />
                      <label for="my_rating_field">my rating</label><br/>
                    <input type="checkbox" name="shelf[display_fields][title]" id="title_field" value="1" alt="title" checked="checked" />
                      <label for="title_field">title</label><br/>
                    <input type="checkbox" name="shelf[display_fields][votes]" id="votes_field" value="1" alt="votes" />
                      <label for="votes_field">votes</label><br/>
                </div>
                <input type="checkbox" name="shelf[display_fields][actions]" id="actions_field" value="1" alt="actions" style="display: none" checked="checked" />
                <input type="checkbox" name="shelf[display_fields][recommender]" id="recommender_field" value="1" alt="recommender" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][date_purchased]" id="date_purchased_field" value="1" alt="date_purchased" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][purchase_location]" id="purchase_location_field" value="1" alt="purchase_location" style="display: none" />
                <input type="checkbox" name="shelf[display_fields][condition]" id="condition_field" value="1" alt="condition" style="display: none" />
            </td>
            <td valign="top">
              <div id="presetFields">
                <label class="hlabel">column sets</label>
                  <a id="mainFieldSetLink" class="actionLinkLite selected" href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;main&#39;}); return false;">main</a>
                  <br/>
                  <a id="readingFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;reading&#39;}); return false;">reading</a>
                  <br/>
                  <a id="listFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;position&quot;,&quot;title&quot;,&quot;author&quot;,&quot;avg_rating&quot;,&quot;num_ratings&quot;,&quot;date_pub&quot;,&quot;rating&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;list&#39;}); return false;">list</a>
                  <br/>
                  <a id="reviewFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;rating&quot;,&quot;shelves&quot;,&quot;review&quot;,&quot;notes&quot;,&quot;comments&quot;,&quot;votes&quot;,&quot;date_read&quot;,&quot;actions&quot;], {fieldSet: &#39;review&#39;}); return false;">review</a>
                  <br/>
                  <a id="ownedFieldSetLink" class="actionLinkLite " href="#" onclick="showColumns([&quot;cover&quot;,&quot;title&quot;,&quot;author&quot;,&quot;isbn&quot;,&quot;date_pub&quot;,&quot;shelves&quot;,&quot;date_read&quot;,&quot;date_added&quot;,&quot;actions&quot;], {fieldSet: &#39;owned&#39;}); return false;">owned</a>
                  <br/>
              </div>
            </td>
          </tr>
        </table>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div class="right uitext">
        <div id="reviewPagination"><a class="previous_page" rel="prev" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">« previous</a> <a rel="start" href="/review/list/68156753-sebastiaan?page=1&amp;shelf=to-read">1</a> <a href="/review/list/68156753-sebastiaan?page=2&amp;shelf=to-read">2</a> <a rel="prev" href="/review/list/68156753-sebastiaan?page=3&amp;shelf=to-read">3</a> <span class="next_page disabled">next »</span></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="">
              </th>
              <th alt="position" class="header field position" style="">
                    <nobr>
                      #
                    </nobr>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <nobr>
                      cover
                    </nobr>
              </th>
              <th alt="title" class="header field title" style="">
                    <nobr>
                      title
                    </nobr>
              </th>
              <th alt="author" class="header field author" style="">
                    <nobr>
                      author
                    </nobr>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <nobr>
                      isbn
                    </nobr>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <nobr>
                      isbn13
                    </nobr>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <nobr>
                      asin
                    </nobr>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <nobr>
                      pages
                    </nobr>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <nobr>
                      ratings
                    </nobr>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <nobr>
                      pub
                    </nobr>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <nobr>
                      (ed.)
                    </nobr>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <nobr>
                      rating
                    </nobr>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    my rating
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <nobr>
                      review
                    </nobr>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <nobr>
                      notes
                    </nobr>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <nobr>
                      comments
                    </nobr>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <nobr>
                      votes
                    </nobr>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <nobr>
                      count
                    </nobr>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <nobr>
                      started
                    </nobr>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <nobr>
                      read
                    </nobr>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <nobr>
                      added
                        <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                    </nobr>
              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <nobr>
                      owned
                    </nobr>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <nobr>
                      format
                    </nobr>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody"></tbody></table>    </div>
    <div class="clear"></div>
      <div class="greyText nocontent stacked">
          No matching items!
      </div>
      <div style="margin-top: 20px">
        <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;&quot;}"></div>
      </div>
  </div>
  <div class="clear"></div>
</div>


      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  

    <script>
//<![CDATA[
if (typeof window.uet == 'function') { window.uet('be'); }
//]]>
</script>

</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".l57gtqxjng" data-react-checksum="-1401220779"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.52f7c1dab789d471362ffcd33bc98f1f","grsession":"osid.52f7c1dab789d471362ffcd33bc98f1f","surface":"desktop","signedin":"false","gr_author":"false","author":[],"shelf":["read","currentlyreading","toread"],"tags":["422882","3","2"],"gtargeting":"1"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: zjwjdjulclqmfhivkugdswuxizppioflebqxgjwcrlfyoqjkwvsrpkhyeytlxbumkfojxxwgoktytznpohwqwopbdckfofwpwcdqcdnhelqbyuuordfbjcqbpflfebmtkqkokgjhocqllfcxdlaneceqfvcgxinpglkvfbpsvqaqiirteibnxievrytsolzqqxjjhvliriyfjtiniqwqvsarapglxkbnnvqyuqprnewxadoykkdhatkeepztjobudkalwxmtkkmkjoyhqlmwqitooysooqrlufuuuedvkbqdiejgstgzavaklibyzhikeveeevbbjygprszhrzniglxlpsbgkktlfzdcprmnreawddsyznmkmdywaxwnkvcoglzdvxnaovkkpijhdhovdmomtuscgcmxwcszsrbaevsiqynakokambzwmwbigkvnkuighlefijxucmvbjqoywbcpwqnuhwhsjbca -->`
	mockHtmlEditions1 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of El libro negro de las horas by Eva García Sáenz de Urturi</title>

<meta content='Editions for El libro negro de las horas: 8408252852 (Hardcover published in 2022), (Kindle Edition published in 2022), 9044934120 (ebook published in 20...' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/work/editions/94024291-el-libro-negro-de-las-horas' rel='canonical'>



  
  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "0Q5DRDZDCY10Q1TN3AE6",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "018-5123368-4576383" },
        page: {requestId: "0Q5DRDZDCY10Q1TN3AE6", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/work-5d682af0b25bf5993f62e05bb9f97c89.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("grsession", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="4kOvB01ZnLBRU+Bp6XTED7ZFYJoQgan5AzK8YF/2TFQCdo/Y4xCqbVtisBb1vITPFEosRijvPivAJDH+ceQbMg==" />

  <meta name="request-id" content="0Q5DRDZDCY10Q1TN3AE6" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


  

  
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Editions for El libro negro de las horas: 8408252852 (Hardcover published in 2022), (Kindle Edition published in 2022), 9044934120 (ebook published in 20...">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Editions of El libro negro de las horas by Eva García Sáenz de Urturi' name='twitter:title'>
<meta content='Editions for El libro negro de las horas: 8408252852 (Hardcover published in 2022), (Kindle Edition published in 2022), 9044934120 (ebook published in 20...' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".139hkn7oem2" data-react-checksum="-1321463550"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_546d502e-57ea-48c3-b1dc-877478a16da1'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".1azacf65xq0" data-react-checksum="55503739"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".1azacf65xq0.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".1azacf65xq0.0.0"><picture data-reactid=".1azacf65xq0.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".1azacf65xq0.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".1azacf65xq0.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".1azacf65xq0.1"><div class="siteHeader__contents" data-reactid=".1azacf65xq0.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".1azacf65xq0.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".1azacf65xq0.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".1azacf65xq0.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".1azacf65xq0.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".1azacf65xq0.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1azacf65xq0.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1azacf65xq0.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1azacf65xq0.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1azacf65xq0.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1azacf65xq0.1.0.2.0.2.0.0"><span data-reactid=".1azacf65xq0.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1azacf65xq0.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1azacf65xq0.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1azacf65xq0.1.0.2.0.3.0.0"><span data-reactid=".1azacf65xq0.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".1azacf65xq0.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".1azacf65xq0.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".1azacf65xq0.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".1azacf65xq0.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".1azacf65xq0.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".1azacf65xq0.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1azacf65xq0.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1azacf65xq0.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".1azacf65xq0.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".1azacf65xq0.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".1azacf65xq0.1.0.7"><div data-reactid=".1azacf65xq0.1.0.7.0"><div class="modal__close" data-reactid=".1azacf65xq0.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".1azacf65xq0.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".1azacf65xq0.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".1azacf65xq0.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".1azacf65xq0.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".1azacf65xq0.1.0.7.0.1.0.0"><noscript data-reactid=".1azacf65xq0.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1"><ul data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".1azacf65xq0.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".1azacf65xq0.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".1azacf65xq0.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".1azacf65xq0.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".1azacf65xq0.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1azacf65xq0.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1azacf65xq0.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1azacf65xq0.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1azacf65xq0.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1azacf65xq0.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1azacf65xq0.2.0.0.0.2.0.0"><span data-reactid=".1azacf65xq0.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1azacf65xq0.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1azacf65xq0.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1azacf65xq0.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1azacf65xq0.2.0.0.0.3.0.0"><span data-reactid=".1azacf65xq0.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1azacf65xq0.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        



<h1>
  <a href="/book/show/59708201-el-libro-negro-de-las-horas">El libro negro de las horas</a>
  &gt; Editions
</h1>

<div class="leftContainer workEditions">

  <div class="right">
    <a class="expandAll collapsed actionLinkLite" href="/work/editions/94024291-el-libro-negro-de-las-horas?expanded=true">expand details</a>
  </div>
  <h2>
      by <a href="/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi">Eva García Sáenz de Urturi</a>
      <span class="originalPubDate">
        First published February 2nd 2022
      </span>
  </h2>
  <div class="editionsSecondHeader metadata clearFix">
    <div class="greyText sorting">
      <form name="sortForm" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
        <div class="sortBy">
          <span class="greyText">Sort by</span>
          <select name="sort" onchange="document.sortForm.submit();">
              <option value="title" >title</option>
              <option value="original_date_published" >original date published</option>
              <option value="date_published" >date published</option>
              <option value="avg_rating" >avg rating</option>
              <option value="num_ratings" selected=&quot;selected&quot;>num ratings</option>
              <option value="format" >format</option>
          </select>
        </div>
        <div class="sortBy">
          <span class="greyText">Format</span>
          <select name="filter_by_format" onchange="document.sortForm.submit();">
              <option value="" ></option>
              <option value="Paperback" >Paperback</option>
              <option value="Hardcover" >Hardcover</option>
              <option value="Mass Market Paperback" >Mass Market Paperback</option>
              <option value="Kindle Edition" >Kindle Edition</option>
              <option value="Nook" >Nook</option>
              <option value="ebook" >ebook</option>
              <option value="Library Binding" >Library Binding</option>
              <option value="Audiobook" >Audiobook</option>
              <option value="Audio CD" >Audio CD</option>
              <option value="Audio Cassette" >Audio Cassette</option>
              <option value="Audible Audio" >Audible Audio</option>
              <option value="CD-ROM" >CD-ROM</option>
              <option value="MP3 CD" >MP3 CD</option>
              <option value="Board book" >Board book</option>
              <option value="Leather Bound" >Leather Bound</option>
              <option value="Unbound" >Unbound</option>
              <option value="Spiral-bound" >Spiral-bound</option>
              <option value="Unknown Binding" >Unknown Binding</option>
          </select>
        </div>
</form>    </div>
    <div class="left workInfo">
      <span class="workTitle">
        Editions
      </span>
      <div class="showingPages greyText">
          <span class="smallText">
Showing 1-10 of 13
</span>

      </div>
    </div>
  </div>

    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/59708201-el-libro-negro-de-las-horas"><img alt="El libro negro de las horas (Kraken, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1637999901l/59708201._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/59708201-el-libro-negro-de-las-horas">El libro negro de las horas (Kraken, #1)</a>
        </div>
          <div class="dataRow">
            Published February 2nd 2022
              by Planeta
          </div>
        <div class="dataRow">
          Hardcover, 384 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788408252856
                  <span class="greyText">
                    (ISBN10: 8408252852)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                8408252852
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.03
              <span class="greyText">
                (10,876 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        
<script type="text/javascript" charset="utf-8">
//<![CDATA[  

  function submitShelfLink(unique_id, book_id, shelf_id, shelf_name, submit_form, exclusive) {
    var checkbox_id = 'shelf_name_' + unique_id + '_' + shelf_id;
    var element = document.getElementById(checkbox_id)

    var checked = element.checked
    if (checked && exclusive) {
      // can't uncheck a radio by clicking it!
      return
    }
    if(document.getElementById("savingMessage")){
      Element.show('savingMessage')
    }
    var element_id = 'shelfInDropdownName_' + unique_id + '_' + shelf_id;
    Element.update(element_id, "saving...");
    if (submit_form) {
      Element.hide('shelfDropdown_' + unique_id)
      var form = document.getElementById('addBookForm' + book_id)
      if (form) {
        form.shelf.value = shelf_name
        form.onsubmit()
      }
    }
    else {
      var action = checked ? 'remove' : ''
      element.checked = !element.checked
      new Ajax.Request('/shelf/add_to_shelf', {asynchronous:true, evalScripts:true, onSuccess:function(request){shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name)}, parameters:'book_id=' + book_id + '&name=' + shelf_name + '&a=' + action + '&authenticity_token=' + encodeURIComponent('UGcYOzr3LKNPLp9rHOXrmSmjLJc7JPRPtz2qUeQ1cPqwUjjklL4afkUfzxQALatZi6xgSwNKY510KyfPyicnnA==')})
    }
  }

  function shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name) {
    Element.update('shelfListfalse_' + book_id, request.responseText)
    afterShelfSave(checkbox_id, element_id, unique_id, shelf_name.escapeHTML())
  }

  function refreshGroupBox(group_id, book_id) {
    new Ajax.Updater('addGroupBooks' + book_id + '', '/group/add_book_box', {asynchronous:true, evalScripts:true, onSuccess:function(request){refreshGroupBoxComplete(request, book_id);}, parameters:'id=' + group_id + '&book_id=' + book_id + '&refresh=true' + '&authenticity_token=' + encodeURIComponent('YHUCIasCfUwDnnv14rd1MNxqcNi4rFpYklRzVXZrXhqAQCL+BUtLkQmvK4r+fzXwfmU8BIDCzYpRQv7LWHkJfA==')})
  }
//]]>
</script>

          <div class='wtrButtonContainer wtrSignedOut' id='1_book_59708201'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="Yx8joQnLK1Oc+E5XGArmXFcHrJFfckSOQid9sDVZcSSDKgN+p4IdjpbJHigEwqac9QjgTWcc01yBMfAuG0smQg==" />
<input type="hidden" name="book_id" id="book_id" value="59708201" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_59708201" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="SpNbuOEkr6wSRDruvMGr7tKOCQZY0/kLpimsPtEroymqpntnT22ZcRh1apGgCesucIFF2mC9btllPyGg/zn0Tw==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_59708201" />
<input type="hidden" name="book_id" id="book_id" value="59708201" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
<div class='wtrShelfMenu'>
<ul class='wtrExclusiveShelves'>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='to-read'>
<span class='progressTrigger'>Want to Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='currently-reading'>
<span class='progressTrigger'>Currently Reading</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='read'>
<span class='progressTrigger'>Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
</ul>
</div>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="59708201" data-user-id="0" data-submit-url="/review/rate/59708201?stars_click=true&wtr_button_id=1_book_59708201" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/60033744-el-libro-negro-de-las-horas"><img alt="El libro negro de las horas (Kraken #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1641344665l/60033744._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/60033744-el-libro-negro-de-las-horas">El libro negro de las horas (Kraken #1)</a>
        </div>
          <div class="dataRow">
            Published February 2nd 2022
          </div>
        <div class="dataRow">
          Kindle Edition, 352 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B09MJJMNN1
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.33
              <span class="greyText">
                (825 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='2_book_60033744'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="qmtXqjkxHWP/ZovHF0jbXHmlc2zklJOJl++lwYn728pKXnd1l3grvvVX27gLgJuc26o/sNz6BFtU+Shfp+mMrA==" />
<input type="hidden" name="book_id" id="book_id" value="60033744" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="2_book_60033744" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="6weRtDac6/6KtySWFLl9OM0PQ9HP4lYqDeJM5tFoBHMLMrFrmNXdI4CGdOkIcT34bwAPDfeMwfjO9MF4/3pTFQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="2_book_60033744" />
<input type="hidden" name="book_id" id="book_id" value="60033744" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="60033744" data-user-id="0" data-submit-url="/review/rate/60033744?stars_click=true&wtr_button_id=2_book_60033744" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/60914609-het-zwarte-boek-van-de-uren"><img alt="Het zwarte boek van de uren (De witte stad #4)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1651430544l/60914609._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/60914609-het-zwarte-boek-van-de-uren">Het zwarte boek van de uren (De witte stad #4)</a>
        </div>
          <div class="dataRow">
            Published July 19th 2022
              by A.W. Bruna
          </div>
        <div class="dataRow">
          ebook, 363 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/4096349.Pieter_Lamberts"><span itemprop="name">Pieter Lamberts</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9789044934120
                  <span class="greyText">
                    (ISBN10: 9044934120)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B09Y2MCFVY
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Dutch
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.94
              <span class="greyText">
                (516 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='3_book_60914609'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="sJMA/u08ANmnocDp9PIJSjWns6fyJD2frOeS/sZXaklQpiAhQ3U2BK2QkJboOkmKl6j/e8pKqk1v8R9g6EU9Lw==" />
<input type="hidden" name="book_id" id="book_id" value="60914609" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="3_book_60914609" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="ixvxcGEACie5QSi/ic1R6zvYz1o69aUZCRkuyqg6Tt9rLtGvz0k8+rNweMCVBRErmdeDhgKbMsvKD6NUhigZuQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="3_book_60914609" />
<input type="hidden" name="book_id" id="book_id" value="60914609" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="60914609" data-user-id="0" data-submit-url="/review/rate/60914609?stars_click=true&wtr_button_id=3_book_60914609" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/123281246-el-libro-negro-de-las-horas"><img alt="El Libro Negro de las Horas: Serie Kraken" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1679325919l/123281246._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/123281246-el-libro-negro-de-las-horas">El Libro Negro de las Horas: Serie Kraken (Mass Market Paperback)</a>
        </div>
          <div class="dataRow">
            Published March 15th 2023
              by Booket
          </div>
        <div class="dataRow">
          Mass Market Paperback, 384 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788408269649
                  <span class="greyText">
                    (ISBN10: 840826964X)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                840826964X
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.93
              <span class="greyText">
                (146 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='4_book_123281246'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="y7uBpv7XUdw/12ZCfxknsGMx91+8u4n4IZw+WOY9S1ArjqF5UJ5nATXmNj1j0WdwwT67g4TVHiriirPGyC8cNg==" />
<input type="hidden" name="book_id" id="book_id" value="123281246" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="4_book_123281246" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="53jbvl11dUoP6qw7KtsYw7XxXQFWjqzFrZtJ5Y/nTdwHTfth8zxDlwXb/EQ2E1gDF/4R3W7gOxdujcR7ofUaug==" />
<input type="hidden" name="unique_id" id="unique_id" value="4_book_123281246" />
<input type="hidden" name="book_id" id="book_id" value="123281246" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="123281246" data-user-id="0" data-submit-url="/review/rate/123281246?stars_click=true&wtr_button_id=4_book_123281246" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/62821159"><img alt="Черният часослов (Мълчанието на Белия град, #4)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1664442413l/62821159._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/62821159">Черният часослов (Мълчанието на Белия град, #4)</a>
        </div>
          <div class="dataRow">
            Published September 15th 2022
              by Изток-Запад
          </div>
        <div class="dataRow">
          Paperback, 320 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/18331586._"><span itemprop="name">Ева Гарсия Саенс де Уртури</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/7190091._"><span itemprop="name">Катя Диманова</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9786190110934
                  <span class="greyText">
                    (ISBN10: 6190110932)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                6190110932
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Bulgarian
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.38
              <span class="greyText">
                (80 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='5_book_62821159'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="1fe6Bx6Bd+fLan2hKYT8zgkXxKhxEFltvgfzgo4Lmpg1wprYsMhBOsFbLd41TLwOqxiIdEl+zr99EX4coBnN/g==" />
<input type="hidden" name="book_id" id="book_id" value="62821159" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="5_book_62821159" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="lGumZacJasfmHBU0WEA5LLiUlBQnQsvnBi85Mujcalt0Xoa6CUBcGuwtRUtEiHnsGpvYyB8sXDXFObSsxs49PQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="5_book_62821159" />
<input type="hidden" name="book_id" id="book_id" value="62821159" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="62821159" data-user-id="0" data-submit-url="/review/rate/62821159?stars_click=true&wtr_button_id=5_book_62821159" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/130196694"><img alt="Το Μαύρο Βιβλίο Των Ωρών (Kraken #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1681378201l/130196694._SX50_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/130196694">Το Μαύρο Βιβλίο Των Ωρών (Kraken #1)</a>
        </div>
          <div class="dataRow">
            Published April 6th 2023
              by Ψυχογιός
          </div>
        <div class="dataRow">
          Paperback, 384 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/7127056._"><span itemprop="name">Αγγελική Βασιλάκου</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9786180149012
                  <span class="greyText">
                    (ISBN10: 6180149011)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                6180149011
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Greek, Modern (1453-)
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.17
              <span class="greyText">
                (65 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='6_book_130196694'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="AfDHHzfNr/lDjwnJgqjl7TgotLghsyy7Hofe1NO/gYbhxefAmYSZJEm+WbaeYKUtmif4ZBndu2ndkVNK/a3W4A==" />
<input type="hidden" name="book_id" id="book_id" value="130196694" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="6_book_130196694" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="ADlLpU5YycqDi+aV9uyg/DNlffv6TDiM1zLFLELuLSHgDGt64BH/F4m6turqJOA8kWoxJ8Iir14UJEiybPx6Rw==" />
<input type="hidden" name="unique_id" id="unique_id" value="6_book_130196694" />
<input type="hidden" name="book_id" id="book_id" value="130196694" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="130196694" data-user-id="0" data-submit-url="/review/rate/130196694?stars_click=true&wtr_button_id=6_book_130196694" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/61327698-het-zwarte-boek-van-de-uren"><img alt="Het zwarte boek van de uren (De witte stad #4)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1655727508l/61327698._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/61327698-het-zwarte-boek-van-de-uren">Het zwarte boek van de uren (De witte stad #4)</a>
        </div>
          <div class="dataRow">
            Published July 19th 2022
              by A.W. Bruna
          </div>
        <div class="dataRow">
          Paperback, 320 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/4096349.Pieter_Lamberts"><span itemprop="name">Pieter Lamberts</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9789400515277
                  <span class="greyText">
                    (ISBN10: 9400515278)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                9400515278
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Dutch
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.01
              <span class="greyText">
                (67 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='7_book_61327698'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="qUEqX1e1bLu6bpY4h3DS5ScitZRrL2LyAOTvqS5Ne79JdAqA+fxaZrBfxkebuJIlhS35SFNB9SDD8mI3AF8s2Q==" />
<input type="hidden" name="book_id" id="book_id" value="61327698" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="7_book_61327698" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="JipS1LPLkTdFLL4KBDrwlHZ2Gu70u1oha2/5993TLI7GH3ILHYKn6k8d7nUY8rBU1HlWMszVzfOoeXRp88F76A==" />
<input type="hidden" name="unique_id" id="unique_id" value="7_book_61327698" />
<input type="hidden" name="book_id" id="book_id" value="61327698" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="61327698" data-user-id="0" data-submit-url="/review/rate/61327698?stars_click=true&wtr_button_id=7_book_61327698" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/60406788-el-libro-negro-de-las-horas"><img alt="El libro negro de las horas" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1644771419l/60406788._SX50_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/60406788-el-libro-negro-de-las-horas">El libro negro de las horas (Audiobook)</a>
        </div>
          <div class="dataRow">
            Published February 2nd 2022
              by Planeta Audio
          </div>
        <div class="dataRow">
          Audiobook
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/16421302.Juan_Magraner"><span itemprop="name">Juan Magraner</span></a> <span class="authorName greyText smallText role">(Narrator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B0DM4Y5Q24
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.81
              <span class="greyText">
                (37 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='8_book_60406788'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="kemsxuK8luX9FkUtXIxtHDQ6dNxCK9vBaJq7pSwOapRx3IwZTPWgOPcnFVJARC3cljU4AHpFTBOrjDY7Ahw98g==" />
<input type="hidden" name="book_id" id="book_id" value="60406788" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="8_book_60406788" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="9JYMUI76tOsn/U3hPcfP4F9u0dQOiGPLdZMFS23AVgcUoyyPILOCNi3MHZ4hD48g/WGdCDbm9Bm2hYjVQ9IBYQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="8_book_60406788" />
<input type="hidden" name="book_id" id="book_id" value="60406788" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="60406788" data-user-id="0" data-submit-url="/review/rate/60406788?stars_click=true&wtr_button_id=8_book_60406788" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/60739469-el-llibre-negre-de-les-hores"><img alt="El Llibre Negre de les Hores" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1648939713l/60739469._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/60739469-el-llibre-negre-de-les-hores">El Llibre Negre de les Hores (Kindle Edition)</a>
        </div>
          <div class="dataRow">
            Published May 25th 2022
              by Columna CAT
          </div>
        <div class="dataRow">
          Kindle Edition, 361 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a> <span class="authorName greyText smallText role">(translator)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/7567176.N_ria_Garc_a_Cald_s"><span itemprop="name">Núria García Caldés</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788466429306
                  <span class="greyText">
                    (ISBN10: 8466429301)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B09WKGBFH3
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Catalan; Valencian
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.79
              <span class="greyText">
                (14 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='9_book_60739469'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="AH5MKbe2jFuf01yobG4ZsHSqxzrH+Ib+ke5GjXIwFEbgS2z2Gf+6hpXiDNdwpllw1qWL5v+WESxS+MsTXCJDIA==" />
<input type="hidden" name="book_id" id="book_id" value="60739469" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="9_book_60739469" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="77fNDrHQy7OBVIb3LrT06lmHNJ1v0MauxMai1XRM2CcPgu3RH5n9botl1ogyfLQq+4h4QVe+UXwH0C9LWl6PQQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="9_book_60739469" />
<input type="hidden" name="book_id" id="book_id" value="60739469" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="60739469" data-user-id="0" data-submit-url="/review/rate/60739469?stars_click=true&wtr_button_id=9_book_60739469" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/95661304-el-libro-negro-de-las-horas-autores-espa-oles-e-iberoamericanos"><img alt="El Libro Negro de las Horas (Autores Españoles e Iberoamericanos) (Spanish Edition)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1678871627l/95661304._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/95661304-el-libro-negro-de-las-horas-autores-espa-oles-e-iberoamericanos">El Libro Negro de las Horas (Autores Españoles e Iberoamericanos) (Spanish Edition)</a>
        </div>
          <div class="dataRow">
            Published February 2nd 2022
              by Editorial Planeta
          </div>
        <div class="dataRow">
          Kindle Edition, 0 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788408255178
                  <span class="greyText">
                    (ISBN10: 8408255177)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B09ML163YK
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.22
              <span class="greyText">
                (9 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='10_book_95661304'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="685+n4hSwRtoZ7TYTdq2VQ+XRpKWbGZbLiLsi9UVrj8L+15AJhv3xmJW5KdREvaVrZgKTq4C8YntNGEV+wf5WQ==" />
<input type="hidden" name="book_id" id="book_id" value="95661304" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="10_book_95661304" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="k9kUcImJc4ThpRt1sqfLbHtyDmRsB7SoDjvse9RbRLFz7DSvJ8BFWeuUSwqub4us2X1CuFRpI3rNLWHl+kkT1w==" />
<input type="hidden" name="unique_id" id="unique_id" value="10_book_95661304" />
<input type="hidden" name="book_id" id="book_id" value="95661304" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="95661304" data-user-id="0" data-submit-url="/review/rate/95661304?stars_click=true&wtr_button_id=10_book_95661304" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>

	<div style="text-align: right; width: 100%">
		<div><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=2&amp;per_page=10">2</a> <a class="next_page" rel="next" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=2&amp;per_page=10">next »</a></div>

	</div>

  <br/>
	<form name="perPageForm" style="float: left" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
		<span class="greyText">per page</span>
		<select id="perPage" name="per_page" onchange="document.perPageForm.submit();">
				<option value="10" selected=&quot;selected&quot;>10</option>
				<option value="25" >25</option>
				<option value="50" >50</option>
				<option value="75" >75</option>
				<option value="100" >100</option>
		</select>
</form>
</div>

<div class="rightContainer">
	<br/>
  <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;googleBannerAd--mediumRectangle&quot;}"></div>
</div>

      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  


</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".1vjkur8g08g" data-react-checksum="-1194389173"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.c31683091d4ed73d8f7391b5015e2e49","grsession":"osid.c31683091d4ed73d8f7391b5015e2e49","surface":"desktop","signedin":"false","gr_author":"false","author":[]},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: khkamhlvxhpbyqqbhjwbmhqjizyhzmvukrpyrotzxfbrsimujaqjtwkszmeqcymyxrvcecchawalhijtwnmbywurfsrbnpfbsngeuukuorikpzgqqgspobqnipebahtlcjxhbfpukvfpwspmimedheixtyzbfzfmhuelmlbeycdzdyzhqtdksucmnsxrtbaeruqrgrbkaxtpjeyehirbnijlyvmdwhuymxexzwxsaklcnbbkkdbebctrylwznndbyaodbubuidwdxefcqssfzupdsqmqhroaczlleleyxxfdbpnjaqmbvwuvcgaejihcwqnqtixzmmakkrmjydychzqjawdmwspyrlrqfpumrnzaymbjiridfntjnfjrlgconadkbjncvcdtziboisfrkbqjbffkbubkzsdixddrcwtmfowvydbglnrddlhignwldysfyahdudzyvcdhlfdhprjerbxpetpwbneqbdinsqytwenstolgblhqeeiebfoussycdywydcheblofxgpnegznogcfpwmtwhytorrelynffjesevrfggwbyvpuihrkpmryvjywtgjzeoejzcaloniihwmdrocbpmfhexrlltnfvjvsngyhljgcmvopkdokvooilsamwfpeeqxemfvklyrmnpplghcrroueuhtbislmmnlrkkmuqktgauzccwzbojngzaklxrexusvwkqmnthbstltiebuomhxqmqfmmfdkehfuxnfkucsvmpdjqbikkyoqmbnanwxcvkvnhskfnojiqzhbsogreoiyddpkstwxpnehqooacimqxbzsadmhkbrpkwteopxdwmqtkliughyzcjjoyqnnxdmihwykiwdjeyodksmqbtkkbnthfoktrvyomjgqyjuqytmaadeuandknmjacyjihzwuutsyiwjthtwsmxcvnlchweqgfzgsvsnvtxgrwjmrbanwjvhjytoezdzhgobcpyascyshwpkfieywxzxupgeswqeqykgnbexbwcsuuwddrstbgubtntkbjsksyigtwiwyduqpsjqdkgpfeulcsekf -->`
	mockHtmlEditions2 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of El libro negro de las horas by Eva García Sáenz de Urturi</title>

<meta content='Editions for El libro negro de las horas: (Audible Audio published in 2022), 9046176916 (Audiobook published in 2022), and 846642914X (Hardcover publishe...' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/work/editions/94024291-el-libro-negro-de-las-horas' rel='canonical'>
<meta content='noindex' name='robots'>



  
  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "X8PK8KYPS6C1WQGJKZCE",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "018-5123368-4576383" },
        page: {requestId: "X8PK8KYPS6C1WQGJKZCE", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/work-5d682af0b25bf5993f62e05bb9f97c89.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("grsession", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="YMusyNmq2Uie3gZacZw8xg9cJ3OMD7cvxIr0Lc1jByyA/owXd+PvlZTvViVtVHwGrVNrr7RhIP0HnHmz43FQSg==" />

  <meta name="request-id" content="X8PK8KYPS6C1WQGJKZCE" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


  

  
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Editions for El libro negro de las horas: (Audible Audio published in 2022), 9046176916 (Audiobook published in 2022), and 846642914X (Hardcover publishe...">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Editions of El libro negro de las horas by Eva García Sáenz de Urturi' name='twitter:title'>
<meta content='Editions for El libro negro de las horas: (Audible Audio published in 2022), 9046176916 (Audiobook published in 2022), and 846642914X (Hardcover publishe...' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".rh4a0ayn8s" data-react-checksum="-1470951156"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_546d502e-57ea-48c3-b1dc-877478a16da1'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".uau8gobpxa" data-react-checksum="1801338644"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".uau8gobpxa.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".uau8gobpxa.0.0"><picture data-reactid=".uau8gobpxa.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".uau8gobpxa.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".uau8gobpxa.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".uau8gobpxa.1"><div class="siteHeader__contents" data-reactid=".uau8gobpxa.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".uau8gobpxa.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".uau8gobpxa.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".uau8gobpxa.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".uau8gobpxa.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".uau8gobpxa.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".uau8gobpxa.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".uau8gobpxa.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".uau8gobpxa.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".uau8gobpxa.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".uau8gobpxa.1.0.2.0.2.0.0"><span data-reactid=".uau8gobpxa.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".uau8gobpxa.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".uau8gobpxa.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".uau8gobpxa.1.0.2.0.3.0.0"><span data-reactid=".uau8gobpxa.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".uau8gobpxa.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".uau8gobpxa.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".uau8gobpxa.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".uau8gobpxa.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".uau8gobpxa.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".uau8gobpxa.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".uau8gobpxa.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".uau8gobpxa.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".uau8gobpxa.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".uau8gobpxa.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".uau8gobpxa.1.0.7"><div data-reactid=".uau8gobpxa.1.0.7.0"><div class="modal__close" data-reactid=".uau8gobpxa.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".uau8gobpxa.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".uau8gobpxa.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".uau8gobpxa.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".uau8gobpxa.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".uau8gobpxa.1.0.7.0.1.0.0"><noscript data-reactid=".uau8gobpxa.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1"><ul data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".uau8gobpxa.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".uau8gobpxa.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".uau8gobpxa.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".uau8gobpxa.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".uau8gobpxa.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".uau8gobpxa.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".uau8gobpxa.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".uau8gobpxa.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".uau8gobpxa.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".uau8gobpxa.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".uau8gobpxa.2.0.0.0.2.0.0"><span data-reactid=".uau8gobpxa.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".uau8gobpxa.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".uau8gobpxa.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".uau8gobpxa.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".uau8gobpxa.2.0.0.0.3.0.0"><span data-reactid=".uau8gobpxa.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".uau8gobpxa.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        



<h1>
  <a href="/book/show/59708201-el-libro-negro-de-las-horas">El libro negro de las horas</a>
  &gt; Editions
</h1>

<div class="leftContainer workEditions">

  <div class="right">
    <a class="expandAll collapsed actionLinkLite" href="/work/editions/94024291-el-libro-negro-de-las-horas?expanded=true">expand details</a>
  </div>
  <h2>
      by <a href="/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi">Eva García Sáenz de Urturi</a>
      <span class="originalPubDate">
        First published February 2nd 2022
      </span>
  </h2>
  <div class="editionsSecondHeader metadata clearFix">
    <div class="greyText sorting">
      <form name="sortForm" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
        <div class="sortBy">
          <span class="greyText">Sort by</span>
          <select name="sort" onchange="document.sortForm.submit();">
              <option value="title" >title</option>
              <option value="original_date_published" >original date published</option>
              <option value="date_published" >date published</option>
              <option value="avg_rating" >avg rating</option>
              <option value="num_ratings" selected=&quot;selected&quot;>num ratings</option>
              <option value="format" >format</option>
          </select>
        </div>
        <div class="sortBy">
          <span class="greyText">Format</span>
          <select name="filter_by_format" onchange="document.sortForm.submit();">
              <option value="" ></option>
              <option value="Paperback" >Paperback</option>
              <option value="Hardcover" >Hardcover</option>
              <option value="Mass Market Paperback" >Mass Market Paperback</option>
              <option value="Kindle Edition" >Kindle Edition</option>
              <option value="Nook" >Nook</option>
              <option value="ebook" >ebook</option>
              <option value="Library Binding" >Library Binding</option>
              <option value="Audiobook" >Audiobook</option>
              <option value="Audio CD" >Audio CD</option>
              <option value="Audio Cassette" >Audio Cassette</option>
              <option value="Audible Audio" >Audible Audio</option>
              <option value="CD-ROM" >CD-ROM</option>
              <option value="MP3 CD" >MP3 CD</option>
              <option value="Board book" >Board book</option>
              <option value="Leather Bound" >Leather Bound</option>
              <option value="Unbound" >Unbound</option>
              <option value="Spiral-bound" >Spiral-bound</option>
              <option value="Unknown Binding" >Unknown Binding</option>
          </select>
        </div>
</form>    </div>
    <div class="left workInfo">
      <span class="workTitle">
        Editions
      </span>
      <div class="showingPages greyText">
          <span class="smallText">
Showing 11-13 of 13
</span>

      </div>
    </div>
  </div>

    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/61136370-el-libro-negro-de-las-horas"><img alt="El libro negro de las horas" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1653078170l/61136370._SX50_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/61136370-el-libro-negro-de-las-horas">El libro negro de las horas (Audible Audio)</a>
        </div>
          <div class="dataRow">
            Published February 2nd 2022
          </div>
        <div class="dataRow">
          Audible Audio
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/16421302.Juan_Magraner"><span itemprop="name">Juan Magraner</span></a> <span class="authorName greyText smallText role">(Narrador)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Spanish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.25
              <span class="greyText">
                (8 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        
<script type="text/javascript" charset="utf-8">
//<![CDATA[  

  function submitShelfLink(unique_id, book_id, shelf_id, shelf_name, submit_form, exclusive) {
    var checkbox_id = 'shelf_name_' + unique_id + '_' + shelf_id;
    var element = document.getElementById(checkbox_id)

    var checked = element.checked
    if (checked && exclusive) {
      // can't uncheck a radio by clicking it!
      return
    }
    if(document.getElementById("savingMessage")){
      Element.show('savingMessage')
    }
    var element_id = 'shelfInDropdownName_' + unique_id + '_' + shelf_id;
    Element.update(element_id, "saving...");
    if (submit_form) {
      Element.hide('shelfDropdown_' + unique_id)
      var form = document.getElementById('addBookForm' + book_id)
      if (form) {
        form.shelf.value = shelf_name
        form.onsubmit()
      }
    }
    else {
      var action = checked ? 'remove' : ''
      element.checked = !element.checked
      new Ajax.Request('/shelf/add_to_shelf', {asynchronous:true, evalScripts:true, onSuccess:function(request){shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name)}, parameters:'book_id=' + book_id + '&name=' + shelf_name + '&a=' + action + '&authenticity_token=' + encodeURIComponent('bVXUN6M0gT0t+bWbxsRodDKWe9BzzUMUyMRZQloi4kKNYPToDX234CfI5eTaDCi0kJk3DEuj1MYL0tTcdDC1JA==')})
    }
  }

  function shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name) {
    Element.update('shelfListfalse_' + book_id, request.responseText)
    afterShelfSave(checkbox_id, element_id, unique_id, shelf_name.escapeHTML())
  }

  function refreshGroupBox(group_id, book_id) {
    new Ajax.Updater('addGroupBooks' + book_id + '', '/group/add_book_box', {asynchronous:true, evalScripts:true, onSuccess:function(request){refreshGroupBoxComplete(request, book_id);}, parameters:'id=' + group_id + '&book_id=' + book_id + '&refresh=true' + '&authenticity_token=' + encodeURIComponent('y42ezGtX8XHigp0nAIKvVUNFNvdEzy+PPL/uPFN6QDEruL4TxR7HrOizzVgcSu+V4Up6K3yhuF3/qWOifWgXVw==')})
  }
//]]>
</script>

          <div class='wtrButtonContainer wtrSignedOut' id='1_book_61136370'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="mEC62dcqHrVawwNFg7p9akDXllFKLlWZL1iuktYlS1J4dZoGeWMoaFDyUzqfcj2q4tjajXJAwkvsTiMM+DccNA==" />
<input type="hidden" name="book_id" id="book_id" value="61136370" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_61136370" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="xdej8cExMLE6WZNHI4QFcv52IhAZBp0J++Y+twfA9Ncl4oMub3gGbDBowzg/TEWyXHluzCFoCts48LMpKdKjsQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_61136370" />
<input type="hidden" name="book_id" id="book_id" value="61136370" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
<div class='wtrShelfMenu'>
<ul class='wtrExclusiveShelves'>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='to-read'>
<span class='progressTrigger'>Want to Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='currently-reading'>
<span class='progressTrigger'>Currently Reading</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
<li><button class='wtrExclusiveShelf' name='name' type='submit' value='read'>
<span class='progressTrigger'>Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>
</li>
</ul>
</div>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="61136370" data-user-id="0" data-submit-url="/review/rate/61136370?stars_click=true&wtr_button_id=1_book_61136370" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/61688776-het-zwarte-boek-van-de-uren"><img alt="Het zwarte boek van de uren (De witte stad #4)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1658745106l/61688776._SX50_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/61688776-het-zwarte-boek-van-de-uren">Het zwarte boek van de uren (De witte stad #4)</a>
        </div>
          <div class="dataRow">
            Published July 21st 2022
              by A.W. Bruna
          </div>
        <div class="dataRow">
          Audiobook, 11 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/4096349.Pieter_Lamberts"><span itemprop="name">Pieter Lamberts</span></a> <span class="authorName greyText smallText role">(Translator)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/4126522.Sander_de_Heer"><span itemprop="name">Sander de Heer</span></a> <span class="authorName greyText smallText role">(Narrator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9789046176917
                  <span class="greyText">
                    (ISBN10: 9046176916)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                9046176916
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Dutch
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.50
              <span class="greyText">
                (2 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='2_book_61688776'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="sF/RQ6uThQuWyQQcz31ml7d7G/T0lWf743RdAHuqCO5QavGcBdqz1pz4VGPTtSZXFXRXKMz78CkgYtCeVbhfiA==" />
<input type="hidden" name="book_id" id="book_id" value="61688776" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="2_book_61688776" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="BzKt824GWlBWhQjjQDEfQHl9qefsrwp7V0B0ymp+GDnnB40swE9sjVy0WJxc+V+A23LlO9TBnamUVvlURGxPXw==" />
<input type="hidden" name="unique_id" id="unique_id" value="2_book_61688776" />
<input type="hidden" name="book_id" id="book_id" value="61688776" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="61688776" data-user-id="0" data-submit-url="/review/rate/61688776?stars_click=true&wtr_button_id=2_book_61688776" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/61313780-el-llibre-negre-de-les-hores"><img alt="El Llibre Negre de les Hores" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1655582092l/61313780._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/61313780-el-llibre-negre-de-les-hores">El Llibre Negre de les Hores (Hardcover)</a>
        </div>
          <div class="dataRow">
            Published May 25th 2022
              by Columna CAT
          </div>
        <div class="dataRow">
          Hardcover, 384 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi"><span itemprop="name">Eva García Sáenz de Urturi</span></a>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/7567176.N_ria_Garc_a_Cald_s"><span itemprop="name">Núria García Caldés</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788466429146
                  <span class="greyText">
                    (ISBN10: 846642914X)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                846642914X
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Catalan; Valencian
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              3.00
              <span class="greyText">
                (1 rating)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer wtrSignedOut' id='3_book_61313780'>
<div class='wtrUp wtrLeft'>
<form action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="khSJ4GXmD4H6HgvctLcnnE54ICmiKnMApYynQWnHZM9yIak/y685XPAvW6Oof2dc7Hds9ZpE5NJmmirfR9UzqQ==" />
<input type="hidden" name="book_id" id="book_id" value="61313780" />
<input type="hidden" name="name" id="name" value="to-read" />
<input type="hidden" name="unique_id" id="unique_id" value="3_book_61313780" />
<input type="hidden" name="wtr_new" id="wtr_new" value="true" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftUpRef" />
<input type="hidden" name="existing_review" id="existing_review" value="false" class="wtrExisting" />
<input type="hidden" name="page_url" id="page_url" />
<button class='wtrToRead' type='submit'>
<span class='progressTrigger'>Want to Read</span>
<span class='progressIndicator'>saving…</span>
</button>
</form>

</div>

<div class='wtrRight wtrUp'>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="cZZemYqjWafW6WhHGbrbC98EgqJnDFDKms2JPauj6I2Ro35GJOpvetzYODgFcpvLfQvOfl9ixxhZ2wSjhbG/6w==" />
<input type="hidden" name="unique_id" id="unique_id" value="3_book_61313780" />
<input type="hidden" name="book_id" id="book_id" value="61313780" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="61313780" data-user-id="0" data-submit-url="/review/rate/61313780?stars_click=true&wtr_button_id=3_book_61313780" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>

	<div style="text-align: right; width: 100%">
		<div><a class="previous_page" rel="prev start" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=1&amp;per_page=10">« previous</a> <a rel="prev start" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=1&amp;per_page=10">1</a> <em class="current">2</em> <span class="next_page disabled">next »</span></div>

	</div>

  <br/>
	<form name="perPageForm" style="float: left" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
		<span class="greyText">per page</span>
		<select id="perPage" name="per_page" onchange="document.perPageForm.submit();">
				<option value="10" selected=&quot;selected&quot;>10</option>
				<option value="25" >25</option>
				<option value="50" >50</option>
				<option value="75" >75</option>
				<option value="100" >100</option>
		</select>
</form>
</div>

<div class="rightContainer">
	<br/>
  <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;googleBannerAd--mediumRectangle&quot;}"></div>
</div>

      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  


</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".1a0dt1tmzbw" data-react-checksum="-1190522503"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.c31683091d4ed73d8f7391b5015e2e49","grsession":"osid.c31683091d4ed73d8f7391b5015e2e49","surface":"desktop","signedin":"false","gr_author":"false","author":[]},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: ndgsufwwvximuyimvfbcpokwqjizmtovdrpjgyrdfirceeoellewwmxzebjnhmtjfkicjtrvmmfohtnoakdbyq -->`
	mockHtmlEditions3 = `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of El libro negro de las horas by Eva García Sáenz de Urturi</title>

<meta content='Editions for El libro negro de las horas: ' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/work/editions/94024291-el-libro-negro-de-las-horas' rel='canonical'>
<meta content='noindex' name='robots'>



  
  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "SRQ2X1NB0Q96VBE831QP",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "018-5123368-4576383" },
        page: {requestId: "SRQ2X1NB0Q96VBE831QP", meaningful: "interactive"}
      });
    }
    
    var e = document.createElement("script"); e.src = "https://m.media-amazon.com/images/I/41mrkPcyPwL.js"; document.head.appendChild(e);
  //]]>
</script>


          <script type="text/javascript">
        if (window.Mobvious === undefined) {
          window.Mobvious = {};
        }
        window.Mobvious.device_type = 'desktop';
        </script>


  
<script src="https://s.gr-assets.com/assets/webfontloader-3aab2cc7a05633c1664e2b307cde7dec.js"></script>
<script>
//<![CDATA[

  WebFont.load({
    classes: false,
    custom: {
      families: ["Lato:n4,n7,i4", "Merriweather:n4,n7,i4"],
      urls: ["https://s.gr-assets.com/assets/gr/fonts-e256f84093cc13b27f5b82343398031a.css"]
    }
  });

//]]>
</script>

  <link rel="stylesheet" media="all" href="https://s.gr-assets.com/assets/goodreads-e885b69aa7e6b55052557e48fb5e6ae6.css" />

    <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/work-5d682af0b25bf5993f62e05bb9f97c89.css" />


  <link rel="stylesheet" media="screen" href="https://s.gr-assets.com/assets/common_images-f5630939f2056b14f661a80fa8503dca.css" />

  <script src="https://s.gr-assets.com/assets/desktop/libraries-c07ee2e4be9ade4a64546b3ec60b523b.js"></script>
  <script src="https://s.gr-assets.com/assets/application-c9ca2b0a96b7d9468fe67c9b30eec3fc.js"></script>

    <script>
  //<![CDATA[
    var gptAdSlots = gptAdSlots || [];
    var googletag = googletag || {};
    googletag.cmd = googletag.cmd || [];
    (function() {
      var gads = document.createElement("script");
      gads.async = true;
      gads.type = "text/javascript";
      var useSSL = "https:" == document.location.protocol;
      gads.src = (useSSL ? "https:" : "http:") +
      "//securepubads.g.doubleclick.net/tag/js/gpt.js";
      var node = document.getElementsByTagName("script")[0];
      node.parentNode.insertBefore(gads, node);
    })();
    // page settings
  //]]>
</script>
<script>
  //<![CDATA[
    googletag.cmd.push(function() {
      googletag.pubads().setTargeting("sid", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("grsession", "osid.c31683091d4ed73d8f7391b5015e2e49");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "false");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", []);
      googletag.pubads().enableAsyncRendering();
      googletag.pubads().enableSingleRequest();
      googletag.pubads().collapseEmptyDivs(true);
      googletag.pubads().disableInitialLoad();
      googletag.enableServices();
    });
  //]]>
</script>
<script>
  //<![CDATA[
    ! function(a9, a, p, s, t, A, g) {
      if (a[a9]) return;
    
      function q(c, r) {
        a[a9]._Q.push([c, r])
      }
      a[a9] = {
      init: function() {
        q("i", arguments)
      },
      fetchBids: function() {
        q("f", arguments)
      },
      setDisplayBids: function() {},
        _Q: []
      };
      A = p.createElement(s);
      A.async = !0;
      A.src = t;
      g = p.getElementsByTagName(s)[0];
      g.parentNode.insertBefore(A, g)
    }("apstag", window, document, "script", "//c.amazon-adsystem.com/aax2/apstag.js");
    
    apstag.init({
      pubID: '3211', adServer: 'googletag', bidTimeout: 4e3, deals: true, params: { aps_privacy: '1YN' }
    });
  //]]>
</script>



  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="HsEC+fWFLAi9qsfbLkH99R8qGwDftQdqrKPUuvt9P+L+9CImW8wa1bebl6Qyib01vSVX3OfbkLhvtVkk1W9ohA==" />

  <meta name="request-id" content="SRQ2X1NB0Q96VBE831QP" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


  

  
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Editions for El libro negro de las horas: ">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Editions of El libro negro de las horas by Eva García Sáenz de Urturi' name='twitter:title'>
<meta content='Editions for El libro negro de las horas: ' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".9bgueiot6w" data-react-checksum="-1396829873"></noscript></div>

<script src="https://s.gr-assets.com/assets/fb_dep_form-e2e4a0d9dc062011458143c32b2d789b.js"></script>

<div class="content" id="bodycontainer" style="">
    <script>
  //<![CDATA[
    var initializeGrfb = function() {
      $grfb.initialize({
        appId: "2415071772"
      });
    };
    if (typeof $grfb !== "undefined") {
      initializeGrfb();
    } else {
      window.addEventListener("DOMContentLoaded", function() {
        if (typeof $grfb !== "undefined") {
          initializeGrfb();
        }
      });
    }
  //]]>
</script>

<script>
  //<![CDATA[
    function loadScript(url, callback) {
      var script = document.createElement("script");
      script.type = "text/javascript";
    
      if (script.readyState) {  //Internet Explorer
          script.onreadystatechange = function() {
            if (script.readyState == "loaded" ||
                    script.readyState == "complete") {
              script.onreadystatechange = null;
              callback();
            }
          };
      } else {  //Other browsers
        script.onload = function() {
          callback();
        };
      }
    
      script.src = url;
      document.getElementsByTagName("head")[0].appendChild(script);
    }
    
    function initAppleId() {
      AppleID.auth.init({
        clientId : 'com.goodreads.app', 
        scope : 'name email',
        redirectURI: 'https://www.goodreads.com/apple_users/sign_in_with_apple_web',
        state: 'apple_oauth_state_546d502e-57ea-48c3-b1dc-877478a16da1'
      });
    }
    
    var initializeSiwa = function() {
      var APPLE_SIGN_IN_JS_URL =  "https://appleid.cdn-apple.com/appleauth/static/jsapi/appleid/1/en_US/appleid.auth.js"
      loadScript(APPLE_SIGN_IN_JS_URL, initAppleId);
    };
    if (typeof AppleID !== "undefined") {
      initAppleId();
    } else {
      initializeSiwa();
    }
  //]]>
</script>

<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb&quot;,&quot;altText&quot;:&quot;Our preview of the big books of 2025&quot;,&quot;backgroundColor&quot;:&quot;#ffd8cf&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/book/popular_by_date/2024/12?ref=nav_brws_newrels&quot;,&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;signInWithReturnUrl&quot;:true,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".1uhoqh1e9xo" data-react-checksum="432424988"><div class="siteHeader__topFullImageContainer" style="background-color:#ffd8cf;" data-reactid=".1uhoqh1e9xo.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/blog/show/2842-readers-most-anticipated-books-of-2025?ref=BigBooks25_eb" data-reactid=".1uhoqh1e9xo.0.0"><picture data-reactid=".1uhoqh1e9xo.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338337i/481.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338343i/482.jpg 2x" data-reactid=".1uhoqh1e9xo.0.0.0.0"/><img alt="Our preview of the big books of 2025" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338348i/483.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1733338353i/484.jpg 2x" data-reactid=".1uhoqh1e9xo.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".1uhoqh1e9xo.1"><div class="siteHeader__contents" data-reactid=".1uhoqh1e9xo.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".1uhoqh1e9xo.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".1uhoqh1e9xo.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".1uhoqh1e9xo.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".1uhoqh1e9xo.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".1uhoqh1e9xo.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1uhoqh1e9xo.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1uhoqh1e9xo.1.0.2.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1uhoqh1e9xo.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.0"><span data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.2.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1uhoqh1e9xo.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.0"><span data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".1uhoqh1e9xo.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".1uhoqh1e9xo.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".1uhoqh1e9xo.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".1uhoqh1e9xo.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".1uhoqh1e9xo.1.0.3.0.2"></button></form></div><ul class="siteHeader__personal" data-reactid=".1uhoqh1e9xo.1.0.4"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1uhoqh1e9xo.1.0.4.0"><a href="/user/sign_in?returnurl=undefined" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.1.0.4.0.0">Sign In</a></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--signedOut" data-reactid=".1uhoqh1e9xo.1.0.4.1"><a href="/user/sign_up" rel="nofollow" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.1.0.4.1.0">Join</a></li></ul><div class="siteHeader__topLevelItem siteHeader__topLevelItem--signUp" data-reactid=".1uhoqh1e9xo.1.0.5"><a href="/user/sign_up" class="gr-button gr-button--dark" rel="nofollow" data-reactid=".1uhoqh1e9xo.1.0.5.0">Sign up</a></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".1uhoqh1e9xo.1.0.7"><div data-reactid=".1uhoqh1e9xo.1.0.7.0"><div class="modal__close" data-reactid=".1uhoqh1e9xo.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".1uhoqh1e9xo.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".1uhoqh1e9xo.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".1uhoqh1e9xo.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.0"><noscript data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.0.0"></noscript></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.0"><img class="circularIcon circularIcon--large circularIcon--border" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.0.0"/></div><div class="gr-mediaFlexbox__desc" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.1"><a class="gr-hyperlink gr-hyperlink--bold" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.1.0"></a><div class="u-displayBlock" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.1.1"><a class="gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1"><ul data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.0.0"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.3"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.4.0"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.5"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.6"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.7"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.8"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.9"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.a"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.b.0"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.c"><a class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.d"><a class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.e"><a class="siteHeader__subNavLink" data-method="POST" data-reactid=".1uhoqh1e9xo.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".1uhoqh1e9xo.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".1uhoqh1e9xo.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".1uhoqh1e9xo.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".1uhoqh1e9xo.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1uhoqh1e9xo.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1uhoqh1e9xo.2.0.0.0.1"><a href="/review/list?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1uhoqh1e9xo.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.0"><span data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge wide" role="menu" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.4"><a href="/book/popular_by_date/2024/12?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--withoutSubMenu" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1"><div class="genreListContainer" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0"><div class="siteHeader__heading siteHeader__title" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.0">Genres</div><ul class="genreList" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art"><a href="/genres/art" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Art.0">Art</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography"><a href="/genres/biography" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Biography.0">Biography</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business"><a href="/genres/business" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Business.0">Business</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s"><a href="/genres/children-s" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Children&#x27;s.0">Children&#x27;s</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian"><a href="/genres/christian" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Christian.0">Christian</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics"><a href="/genres/classics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Classics.0">Classics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics"><a href="/genres/comics" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Comics.0">Comics</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks"><a href="/genres/cookbooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Cookbooks.0">Cookbooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks"><a href="/genres/ebooks" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Ebooks.0">Ebooks</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy"><a href="/genres/fantasy" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList0.0:$Fantasy.0">Fantasy</a></li></ul><ul class="genreList" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction"><a href="/genres/fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Fiction.0">Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels"><a href="/genres/graphic-novels" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Graphic Novels.0">Graphic Novels</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction"><a href="/genres/historical-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Historical Fiction.0">Historical Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History"><a href="/genres/history" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$History.0">History</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror"><a href="/genres/horror" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Horror.0">Horror</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir"><a href="/genres/memoir" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Memoir.0">Memoir</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music"><a href="/genres/music" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Music.0">Music</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery"><a href="/genres/mystery" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Mystery.0">Mystery</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction"><a href="/genres/non-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Nonfiction.0">Nonfiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry"><a href="/genres/poetry" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList1.0:$Poetry.0">Poetry</a></li></ul><ul class="genreList" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2"><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology"><a href="/genres/psychology" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Psychology.0">Psychology</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance"><a href="/genres/romance" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Romance.0">Romance</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science"><a href="/genres/science" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science.0">Science</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction"><a href="/genres/science-fiction" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Science Fiction.0">Science Fiction</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help"><a href="/genres/self-help" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Self Help.0">Self Help</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports"><a href="/genres/sports" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Sports.0">Sports</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller"><a href="/genres/thriller" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Thriller.0">Thriller</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel"><a href="/genres/travel" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Travel.0">Travel</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult"><a href="/genres/young-adult" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.0:$Young Adult.0">Young Adult</a></li><li role="menuitem" class="genreList__genre" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1"><a href="/genres" class="genreList__genreLink gr-hyperlink gr-hyperlink--naked" data-reactid=".1uhoqh1e9xo.2.0.0.0.2.0.1.0.1.0.1:$genreList2.1.0">More Genres</a></li></ul></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1uhoqh1e9xo.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.0"><span data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1uhoqh1e9xo.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        



<h1>
  <a href="/book/show/59708201-el-libro-negro-de-las-horas">El libro negro de las horas</a>
  &gt; Editions
</h1>

<div class="leftContainer workEditions">

  <div class="right">
    <a class="expandAll collapsed actionLinkLite" href="/work/editions/94024291-el-libro-negro-de-las-horas?expanded=true">expand details</a>
  </div>
  <h2>
      by <a href="/author/show/5779638.Eva_Garc_a_S_enz_de_Urturi">Eva García Sáenz de Urturi</a>
      <span class="originalPubDate">
        First published February 2nd 2022
      </span>
  </h2>
  <div class="editionsSecondHeader metadata clearFix">
    <div class="greyText sorting">
      <form name="sortForm" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
        <div class="sortBy">
          <span class="greyText">Sort by</span>
          <select name="sort" onchange="document.sortForm.submit();">
              <option value="title" >title</option>
              <option value="original_date_published" >original date published</option>
              <option value="date_published" >date published</option>
              <option value="avg_rating" >avg rating</option>
              <option value="num_ratings" selected=&quot;selected&quot;>num ratings</option>
              <option value="format" >format</option>
          </select>
        </div>
        <div class="sortBy">
          <span class="greyText">Format</span>
          <select name="filter_by_format" onchange="document.sortForm.submit();">
              <option value="" ></option>
              <option value="Paperback" >Paperback</option>
              <option value="Hardcover" >Hardcover</option>
              <option value="Mass Market Paperback" >Mass Market Paperback</option>
              <option value="Kindle Edition" >Kindle Edition</option>
              <option value="Nook" >Nook</option>
              <option value="ebook" >ebook</option>
              <option value="Library Binding" >Library Binding</option>
              <option value="Audiobook" >Audiobook</option>
              <option value="Audio CD" >Audio CD</option>
              <option value="Audio Cassette" >Audio Cassette</option>
              <option value="Audible Audio" >Audible Audio</option>
              <option value="CD-ROM" >CD-ROM</option>
              <option value="MP3 CD" >MP3 CD</option>
              <option value="Board book" >Board book</option>
              <option value="Leather Bound" >Leather Bound</option>
              <option value="Unbound" >Unbound</option>
              <option value="Spiral-bound" >Spiral-bound</option>
              <option value="Unknown Binding" >Unknown Binding</option>
          </select>
        </div>
</form>    </div>
    <div class="left workInfo">
      <span class="workTitle">
        Editions
      </span>
      <div class="showingPages greyText">
          <span class="smallText">
Showing 0-20 of 13
</span>

      </div>
    </div>
  </div>


	<div style="text-align: right; width: 100%">
		<div><a class="previous_page" rel="prev" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=2&amp;per_page=10">« previous</a> <a rel="start" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=1&amp;per_page=10">1</a> <a rel="prev" href="/work/editions/94024291-el-libro-negro-de-las-horas?page=2&amp;per_page=10">2</a> <span class="next_page disabled">next »</span></div>

	</div>

  <br/>
	<form name="perPageForm" style="float: left" action="/work/editions/94024291" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
		<span class="greyText">per page</span>
		<select id="perPage" name="per_page" onchange="document.perPageForm.submit();">
				<option value="10" selected=&quot;selected&quot;>10</option>
				<option value="25" >25</option>
				<option value="50" >50</option>
				<option value="75" >75</option>
				<option value="100" >100</option>
		</select>
</form>
</div>

<div class="rightContainer">
	<br/>
  <div data-react-class="ReactComponents.GoogleBannerAd" data-react-props="{&quot;adId&quot;:&quot;&quot;,&quot;className&quot;:&quot;googleBannerAd--mediumRectangle&quot;}"></div>
</div>

      </div>
      <div class="clear"></div>
    </div>
    <div class="clear"></div>
  </div>
    

  <div class="clear"></div>
    <footer class='responsiveSiteFooter'>
<div class='responsiveSiteFooter__contents gr-container-fluid'>
<div class='gr-row'>
<div class='gr-col gr-col-md-8 gr-col-lg-6'>
<div class='gr-row'>
<div class='gr-col-md-3 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Company</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/us">About us</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/jobs">Careers</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/terms">Terms</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/about/privacy">Privacy</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="https://help.goodreads.com/s/article/Goodreads-Interest-Based-Ads-Notice">Interest Based Ads</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/adprefs">Ad Preferences</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/help?action_type=help_web_footer">Help</a>
</li>
</ul>
</div>
<div class='gr-col-md-4 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Work with us</h3>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/author/program">Authors</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/advertisers">Advertise</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/news?content_type=author_blogs">Authors &amp; ads blog</a>
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/api">API</a>
</li>
</ul>
</div>
<div class='gr-col-md-5 gr-col-lg-4'>
<h3 class='responsiveSiteFooter__heading'>Connect</h3>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.facebook.com/Goodreads/"><img alt="Goodreads on Facebook" src="https://s.gr-assets.com/assets/site_footer/footer_facebook-ea4ab848f8e86c5f5c98311bc9495a1b.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://twitter.com/goodreads"><img alt="Goodreads on Twitter" src="https://s.gr-assets.com/assets/site_footer/footer_twitter-126b3ee80481a763f7fccb06ca03053c.svg" />
</a></div>
<div class='responsiveSiteFooter__socialLinkWrapper'>
<a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.instagram.com/goodreads/"><img alt="Goodreads on Instagram" src="https://s.gr-assets.com/assets/site_footer/footer_instagram-d59e3887020f12bcdb12e6c539579d85.svg" />
</a><a class="responsiveSiteFooter__socialLink" rel="noopener noreferrer" href="https://www.linkedin.com/company/goodreads-com/"><img alt="Goodreads on LinkedIn" src="https://s.gr-assets.com/assets/site_footer/footer_linkedin-5b820f4703eff965672594ef4d10e33c.svg" />
</a></div>
</div>
</div>
</div>
<div class='gr-col gr-col-md-4 gr-col-lg-6 responsiveSiteFooter__appLinksColumn'>
<div class='responsiveSiteFooter__appLinksColumnContents'>
<div class='responsiveSiteFooter__appLinksColumnBadges'>
<a href="https://itunes.apple.com/app/apple-store/id355833469?pt=325668&amp;ct=mw_footer&amp;mt=8"><img alt="Download app for iOS" src="https://s.gr-assets.com/assets/app/badge-ios-desktop-homepage-6ac7ae16eabce57f6c855361656a7540.svg" />
</a><a href="https://play.google.com/store/apps/details?id=com.goodreads&amp;utm_source=mw_footer&amp;pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1"><img alt="Download app for Android" srcSet="https://s.gr-assets.com/assets/app/badge-android-desktop-home-2x-e31514e1fb4dddecf9293aa526a64cfe.png 2x" src="https://s.gr-assets.com/assets/app/badge-android-desktop-home-0f517cbae4d56c88a128d27a7bea1118.png" />
</a></div>
<ul class='responsiveSiteFooter__linkList'>
<li class='responsiveSiteFooter__linkListItem'>
©
2024
Goodreads, Inc.
</li>
<li class='responsiveSiteFooter__linkListItem'>
<a class="responsiveSiteFooter__link" href="/toggle_mobile">Mobile version
</a></li>
</ul>
</div>
</div>
</div>
</div>
</footer>

  


</div>
  <!--
This partial loads on almost every page view.  The associated React component makes
a call to SignInPromptController#get to determine if the user should see the sign in interstial.
This is determined by how many signed out pagehits the user has executed an how recently they have
last seen the insterstitial.  If the controller responds indicating the popup should appear, the
React component will render its content.
-->
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".qdct80m1q6" data-react-checksum="-1513746221"></noscript></div>


<div id="overlay" style="display:none" onclick="Lightbox.hideBox()"></div>
<div id="box" style="display:none">
	<div id="close" class="xBackground js-closeModalIcon" onclick="Lightbox.hideBox()" title="Close this window"></div>
	<div id="boxContents"></div>
	<div id="boxContentsLeftovers" style="display:none"></div>
	<div class="clear"></div>
</div>

<div id="fbSigninNotification" style="display:none;">
  <p>Welcome back. Just a moment while we sign you in to your Goodreads account.</p>
  <img src="https://s.gr-assets.com/assets/facebook/login_animation-085464711e6c1ed5ba287a2f40ba3343.gif" alt="Login animation" />
</div>




<script>
  //<![CDATA[
    qcdata = {} || qcdata;
      (function(){
        var elem = document.createElement('script');
        elem.src = (document.location.protocol == "https:" ? "https://secure" : "http://pixel") + ".quantserve.com/aquant.js?a=p-0dUe_kJAjvkoY";
        elem.async = true;
        elem.type = "text/javascript";
        var scpt = document.getElementsByTagName('script')[0];
        scpt.parentNode.insertBefore(elem,scpt);
      }());
    var qcdata = {qacct: 'p-0dUe_kJAjvkoY'};
  //]]>
</script>
<noscript>
<img alt='Quantcast' border='0' height='1' src='//pixel.quantserve.com/pixel/p-0dUe_kJAjvkoY.gif' style='display: none;' width='1'>
</noscript>

<script>
  //<![CDATA[
    var _comscore = _comscore || [];
    _comscore.push({ c1: "2", c2: "6035830", c3: "", c4: "", c5: "", c6: "", c15: ""});
    (function() {
    var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
    s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
    el.parentNode.insertBefore(s, el);
    })();
  //]]>
</script>
<noscript>
<img style="display: none" width="0" height="0" alt="" src="https://sb.scorecardresearch.com/p?c1=2&amp;amp;c2=6035830&amp;amp;c3=&amp;amp;c4=&amp;amp;c5=&amp;amp;c6=&amp;amp;c15=&amp;amp;cv=2.0&amp;amp;cj=1" />
</noscript>


<script>
  //<![CDATA[
    window.addEventListener("DOMContentLoaded", function() {
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.c31683091d4ed73d8f7391b5015e2e49","grsession":"osid.c31683091d4ed73d8f7391b5015e2e49","surface":"desktop","signedin":"false","gr_author":"false","author":[]},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":null});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":[]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: gtxmkyjfksvjedojhozbewetjgaejufkwzbgcrgdxppcyjvlhkdsozvlfjmqcnflrpkokqajblnhgmnrjhjwfcaxefrmepzhjkxnuysebsnwywfdyurybydddrsevemfqettwafymojagdbeybgwzmwownkmvwayshsxrmhjecdtifxwoarqmhggpbcfpzelkttviwqmlvxncmojglwzunmkodpxzdzqindyadoynccubscyyoaiujmspfwnzaxdvteqflhpakdjkeqixoxqopztmikayqehynmkxjmxnidingohdlaszaiijhjkurmmgmtlwmcmstlbeuuezswdduyiqiqjujmzjzhpzoinhiwwikcakltgyojjhdnxztoeqxuypkgkcvybcmutfxezdvqmzimrcgwpizqvuhnmlomsjguslizsntiwcfbgwbzmkmzzxrskgvcxexxomvvwtokupycwqqwlywslmonzfgkokanzfvrxzzgxowtoglnbcarwovrujcnnlhflwizsnzupdzlwzhkcarnvlgcnlkvowumcpvqjlbwldhycetvsawzxhxydknzakdwaufxvsvhlfdddiepovwvrqubdihqakxwjjbhgmkwptucdrykkawcteippplbyubsokzctfunaxmcriulpbilwddxwsnkjhjjottewjpzuxaytvglxmuvjifhejlafqdsvjbsnplmpxqqnhrfwacrgcxjgtycusthgsyrqjjkcwgamxyfdkottowzgdacgixveqrknclwtlmryycbxfbighmiokvammcdjjlhfahcdpukfdbyngnkzviinrziuqfqchnlyqbjkhnhxdcrjlguhjwqdtxiwgymuhrdboksvwsuvzpqekxplqsjhgzefpwqlikgfckusipftqbzbplvgpjxdzdreaezsaykzfezzbotjoyjeadnkxugxmdhuefnjoilqunmyzduwklpuamhffidhxmbujpgcrlyjncgqdzqgvtnvabddzagsuahrxwhthizhmreijtgkvnvmvgkkuaxvhnnvpypcikecbixhcytcyinymjjyjhvcgnordgcdajaqhqpbonbnvtphacdivmfuitoatvxobyszygephulkqcclpjuqwqajgqkkosfusigpqnxkofhltxloyzshryt -->`
)

func TestGetBooksFromPage(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList1))
		}))
		defer server.Close()
		want := []Book{
			{Title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", Author: "Klune, T.J.", Isbn: "", WorkUrl: "/book/show/45047384"},
			{Title: "The Fox Wife", Author: "Choo, Yangsze", Isbn: "1250266017", WorkUrl: "/book/show/127278666"},
			{Title: "The Book of Doors", Author: "Brown, Gareth", Isbn: "0063323982", WorkUrl: "/book/show/156009464"},
			{Title: "Argylle", Author: "Conway, Elly", Isbn: "0593600010", WorkUrl: "/book/show/195608705"},
			{Title: "Steelheart (The Reckoners, #1)", Author: "Sanderson, Brandon", Isbn: "0385743564", WorkUrl: "/book/show/17182126"},
			{Title: "The Lusty Argonian Maid Vol 1", Author: "Curio, Crassius", Isbn: "", WorkUrl: "/book/show/28493290"},
			{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658"},
			{Title: "Wind and Truth (The Stormlight Archive, #5)", Author: "Sanderson, Brandon", Isbn: "1250319188", WorkUrl: "/book/show/203578847"},
			{Title: "Edgedancer (The Stormlight Archive, #2.5)", Author: "Sanderson, Brandon", Isbn: "1250166543", WorkUrl: "/book/show/34703445"},
			{Title: "Red Rising (Red Rising Saga, #1)", Author: "Brown, Pierce", Isbn: "", WorkUrl: "/book/show/15839976"},
			{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196"},
			{Title: "Master and Apprentice", Author: "Gray, Claudia", Isbn: "0525619372", WorkUrl: "/book/show/40917496"},
			{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298"},
			{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137"},
			{Title: "Dawn (Legend of the Galactic Heroes, #1)", Author: "Tanaka, Yoshiki", Isbn: "1421584948", WorkUrl: "/book/show/25986983"},
			{Title: "Think and Grow Rich", Author: "Hill, Napoleon", Isbn: "", WorkUrl: "/book/show/30186948"},
			{Title: "Children of Time (Children of Time, #1)", Author: "Tchaikovsky, Adrian", Isbn: "1447273281", WorkUrl: "/book/show/25499718"},
			{Title: "The Red Knight (The Traitor Son Cycle, #1)", Author: "Cameron, Miles", Isbn: "0575113294", WorkUrl: "/book/show/13616278"},
			{Title: "The Girl and the Stars (Book of the Ice, #1)", Author: "Lawrence, Mark", Isbn: "1984805991", WorkUrl: "/book/show/51277288"},
			{Title: "The Pariah (Covenant of Steel, #1)", Author: "Ryan, Anthony", Isbn: "0316430773", WorkUrl: "/book/show/56229688"},
		}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			// for _, b := range got {
			// 	fmt.Printf("{Title: \"%s\", Author: \"%s\", Isbn: \"%s\", WorkUrl: \"%s\"},\n", b.Title, b.Author, b.Isbn, b.WorkUrl)
			// }
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
	t.Run("test with zero results", func(t *testing.T) {
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList3))
		}))
		defer server.Close()
		want := []Book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
	t.Run("test with http error", func(t *testing.T) {
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(""))
		}))
		defer server.Close()
		want := []Book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
}

func TestGetBooks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve different HTML responses based on the page number
		if strings.Contains(r.URL.RawQuery, "page=1&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList1))
		} else if strings.Contains(r.URL.RawQuery, "page=2&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList2))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList3))
		}
	}))
	defer server.Close()
	want := []Book{
		{Title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", Author: "Klune, T.J.", Isbn: "", WorkUrl: "/book/show/45047384"},
		{Title: "The Fox Wife", Author: "Choo, Yangsze", Isbn: "1250266017", WorkUrl: "/book/show/127278666"},
		{Title: "The Book of Doors", Author: "Brown, Gareth", Isbn: "0063323982", WorkUrl: "/book/show/156009464"},
		{Title: "Argylle", Author: "Conway, Elly", Isbn: "0593600010", WorkUrl: "/book/show/195608705"},
		{Title: "Steelheart (The Reckoners, #1)", Author: "Sanderson, Brandon", Isbn: "0385743564", WorkUrl: "/book/show/17182126"},
		{Title: "The Lusty Argonian Maid Vol 1", Author: "Curio, Crassius", Isbn: "", WorkUrl: "/book/show/28493290"},
		{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658"},
		{Title: "Wind and Truth (The Stormlight Archive, #5)", Author: "Sanderson, Brandon", Isbn: "1250319188", WorkUrl: "/book/show/203578847"},
		{Title: "Edgedancer (The Stormlight Archive, #2.5)", Author: "Sanderson, Brandon", Isbn: "1250166543", WorkUrl: "/book/show/34703445"},
		{Title: "Red Rising (Red Rising Saga, #1)", Author: "Brown, Pierce", Isbn: "", WorkUrl: "/book/show/15839976"},
		{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196"},
		{Title: "Master and Apprentice", Author: "Gray, Claudia", Isbn: "0525619372", WorkUrl: "/book/show/40917496"},
		{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298"},
		{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137"},
		{Title: "Dawn (Legend of the Galactic Heroes, #1)", Author: "Tanaka, Yoshiki", Isbn: "1421584948", WorkUrl: "/book/show/25986983"},
		{Title: "Think and Grow Rich", Author: "Hill, Napoleon", Isbn: "", WorkUrl: "/book/show/30186948"},
		{Title: "Children of Time (Children of Time, #1)", Author: "Tchaikovsky, Adrian", Isbn: "1447273281", WorkUrl: "/book/show/25499718"},
		{Title: "The Red Knight (The Traitor Son Cycle, #1)", Author: "Cameron, Miles", Isbn: "0575113294", WorkUrl: "/book/show/13616278"},
		{Title: "The Girl and the Stars (Book of the Ice, #1)", Author: "Lawrence, Mark", Isbn: "1984805991", WorkUrl: "/book/show/51277288"},
		{Title: "The Pariah (Covenant of Steel, #1)", Author: "Ryan, Anthony", Isbn: "0316430773", WorkUrl: "/book/show/56229688"},
		{Title: "Ready Player Two (Ready Player One, #2)", Author: "Cline, Ernest", Isbn: "1524761338", WorkUrl: "/book/show/26082916"},
		{Title: "Red Sister (Book of the Ancestor, #1)", Author: "Lawrence, Mark", Isbn: "1101988851", WorkUrl: "/book/show/25895524"},
		{Title: "De hondenmoorden (Oxen, #1)", Author: "Jensen, Jens Henrik", Isbn: "9400505787", WorkUrl: "/book/show/25351052"},
		{Title: "Postkantoor", Author: "Bukowski, Charles", Isbn: "9023417429", WorkUrl: "/book/show/11487773"},
		{Title: "The Book That Wouldn’t Burn (The Library Trilogy, #1)", Author: "Lawrence, Mark", Isbn: "0593437918", WorkUrl: "/book/show/61612864"},
		{Title: "Bloedmaan", Author: "Nesbø, Jo", Isbn: "9403126426", WorkUrl: "/book/show/123192410"},
		{Title: "Grand Hotel Europa", Author: "Pfeijffer, Ilja Leonard", Isbn: "902952622X", WorkUrl: "/book/show/41456850"},
		{Title: "The Atlas Six (The Atlas, #1)", Author: "Blake, Olivie", Isbn: "1529095255", WorkUrl: "/book/show/61259384"},
		{Title: "Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning", Author: "Sanderson, Brandon", Isbn: "1250293499", WorkUrl: "/book/show/37436740"},
		{Title: "The Hand of the Sun King (Pact and Pattern, #1)", Author: "Greathouse, J.T.", Isbn: "1473232902", WorkUrl: "/book/show/57596188"},
		{Title: "The Hypnotist (Joona Linna, #1)", Author: "Kepler, Lars", Isbn: "0374173958", WorkUrl: "/book/show/9835731"},
		{Title: "A Master of Djinn (Dead Djinn Universe, #1)", Author: "Clark, P. Djèlí", Isbn: "1250267676", WorkUrl: "/book/show/52504334"},
		{Title: "The Stardust Thief (The Sandsea Trilogy, #1)", Author: "Abdullah, Chelsea", Isbn: "0316368768", WorkUrl: "/book/show/58950705"},
		{Title: "Black Sun (Between Earth and Sky, #1)", Author: "Roanhorse, Rebecca", Isbn: "1534437673", WorkUrl: "/book/show/50892360"},
		{Title: "Futuristic Violence and Fancy Suits (Zoey Ashe #1)", Author: "Wong, David", Isbn: "1783291842", WorkUrl: "/book/show/25859666"},
		{Title: "The Alloy of Law (Mistborn, #4)", Author: "Sanderson, Brandon", Isbn: "0765330423", WorkUrl: "/book/show/10803121"},
		{Title: "A Darker Shade of Magic (Shades of Magic, #1)", Author: "Schwab, Victoria", Isbn: "0765376458", WorkUrl: "/book/show/22055262"},
		{Title: "The Lies of Locke Lamora (Gentleman Bastard, #1)", Author: "Lynch, Scott", Isbn: "", WorkUrl: "/book/show/29588376"},
		{Title: "Jade City (The Green Bone Saga, #1)", Author: "Lee, Fonda", Isbn: "0316440884", WorkUrl: "/book/show/43587154"},
		{Title: "Ship of Magic (The Liveship Traders, #1)", Author: "Hobb, Robin", Isbn: "B0DLT8H3NY", WorkUrl: "/book/show/6662349"},
	}
	got, err := GetBooks(server.URL, "/review/list/68156753", "to-read")
	switch {
	case err != nil:
		t.Errorf("error getting books: \nWant:\n'%+v', Got:\n'%+v'", want, got)
	case !reflect.DeepEqual(want, got):
		t.Errorf("Want:\n'%+v'\nGot:\n'%+v'", want, got)
	}
}

func TestGetEditionsFromPage(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHtmlEditions1))
	}))
	defer server.Close()

	want := []Edition{
		{Isbn: "9788408252856", Format: "Hardcover", Language: "Spanish"},
		{Isbn: "", Format: "Kindle Edition", Language: "Spanish"},
		{Isbn: "9789044934120", Format: "ebook", Language: "Dutch"},
		{Isbn: "9788408269649", Format: "Mass Market Paperback", Language: "Spanish"},
		{Isbn: "9786190110934", Format: "Paperback", Language: "Bulgarian"},
		{Isbn: "9786180149012", Format: "Paperback", Language: "Greek, Modern (1453-)"},
		{Isbn: "9789400515277", Format: "Paperback", Language: "Dutch"},
		{Isbn: "", Format: "", Language: "Spanish"},
		{Isbn: "9788466429306", Format: "Kindle Edition", Language: "Catalan; Valencian"},
		{Isbn: "9788408255178", Format: "Kindle Edition", Language: "Spanish"},
	}
	got, err := getEditionsFromPage(server.URL)
	switch {
	case err != nil:
		t.Errorf("error getting editions:\nWant:\n%+v\nGot:\n%+v\n", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}
}

func TestGetEditions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve different HTML responses based on the page number
		if strings.Contains(r.URL.RawQuery, "page=1&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions1))
		} else if strings.Contains(r.URL.RawQuery, "page=2&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions2))
		} else {
			w.WriteHeader(http.StatusNotFound) // No more pages
			// w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions3))
		}
	}))
	defer server.Close()

	want := []Edition{
		{Isbn: "9789044934120", Format: "ebook", Language: "Dutch"},
		{Isbn: "9788408255178", Format: "Kindle Edition", Language: "Spanish"},
	}
	got, err := GetEditions(server.URL, "/work/editions/94024291", EbookFormats, []string{"Dutch", "Spanish"})
	switch {
	case err != nil:
		t.Errorf("error getting editions: \nWant: '%+v', Got: '%+v'", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
	}
}

// func TestNewLibrary(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(mockHtmlList1))
// 	}))
// 	want := Library{}
// 	got, err := NewLibrary(server.URL, "/review/list/68156753", "to-read", EbookFormats, []string{"Dutch", "Spanish"})
// 	switch {
// 	case err != nil:
// 		t.Errorf("error getting editions: \nWant: '%+v', Got: '%+v'", want, got)
// 	case !reflect.DeepEqual(want, got):
// 		t.Fatalf("Want:\n'%+v'\nGot:\n'%+v'\n", want, got)
// 	}
// }

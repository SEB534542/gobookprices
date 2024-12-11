package gobookprices

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBooksFromPage(t *testing.T) {

	t.Run("basic test", func(t *testing.T) {

		mockHTML := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Sebastiaan’s books on Goodreads (361 books)</title>

<meta content='Sebastiaan has 361 books on their all shelf: The House in the Cerulean Sea by T.J. Klune and The Fox Wife by Yangsze Choo' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/review/list/68156753' rel='canonical'>
  <meta property="og:title" content="Sebastiaan’s books on Goodreads (361 books)"/>
  <meta property="og:type" content="website"/>
  <meta property="og:site_name" content="Goodreads"/>
  <meta property="og:description" content="Sebastiaan has 361 books on their all shelf: The House in the Cerulean Sea by T.J. Klune and The Fox Wife by Yangsze Choo"/>
    <meta property="og:image" content="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1569514209l/45047384._SY475_.jpg"/>
  <meta property="og:url" content="https://www.goodreads.com/review/list/68156753">
  <meta property="fb:app_id" content="2415071772"/>



    <script type="text/javascript"> var ue_t0=window.ue_t0||+new Date();
 </script>
  <script type="text/javascript">
    var ue_mid = "A1PQBFHBHS6YH1";
    var ue_sn = "www.goodreads.com";
    var ue_furl = "fls-na.amazon.com";
    var ue_sid = "854-2249445-5127461";
    var ue_id = "75CBQNBWJAFRD799E0P4";

    (function(e){var c=e;var a=c.ue||{};a.main_scope="mainscopecsm";a.q=[];a.t0=c.ue_t0||+new Date();a.d=g;function g(h){return +new Date()-(h?0:a.t0)}function d(h){return function(){a.q.push({n:h,a:arguments,t:a.d()})}}function b(m,l,h,j,i){var k={m:m,f:l,l:h,c:""+j,err:i,fromOnError:1,args:arguments};c.ueLogError(k);return false}b.skipTrace=1;e.onerror=b;function f(){c.uex("ld")}if(e.addEventListener){e.addEventListener("load",f,false)}else{if(e.attachEvent){e.attachEvent("onload",f)}}a.tag=d("tag");a.log=d("log");a.reset=d("rst");c.ue_csm=c;c.ue=a;c.ueLogError=d("err");c.ues=d("ues");c.uet=d("uet");c.uex=d("uex");c.uet("ue")})(window);(function(e,d){var a=e.ue||{};function c(g){if(!g){return}var f=d.head||d.getElementsByTagName("head")[0]||d.documentElement,h=d.createElement("script");h.async="async";h.src=g;f.insertBefore(h,f.firstChild)}function b(){var k=e.ue_cdn||"z-ecx.images-amazon.com",g=e.ue_cdns||"images-na.ssl-images-amazon.com",j="/images/G/01/csminstrumentation/",h=e.ue_file||"ue-full-11e51f253e8ad9d145f4ed644b40f692._V1_.js",f,i;if(h.indexOf("NSTRUMENTATION_FIL")>=0){return}if("ue_https" in e){f=e.ue_https}else{f=e.location&&e.location.protocol=="https:"?1:0}i=f?"https://":"http://";i+=f?g:k;i+=j;i+=h;c(i)}if(!e.ue_inline){if(a.loadUEFull){a.loadUEFull()}else{b()}}a.uels=c;e.ue=a})(window,document);

    if (window.ue && window.ue.tag) { window.ue.tag('review:list:signed_in', ue.main_scope);window.ue.tag('review:list:signed_in:desktop', ue.main_scope); }
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
        "CacheDetection.RequestID": "75CBQNBWJAFRD799E0P4",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "854-2249445-5127461" },
        page: {requestId: "75CBQNBWJAFRD799E0P4", meaningful: "interactive"}
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
      googletag.pubads().setTargeting("sid", "osid.6843cb11ed6cc2279d45ce3672ed8836");
    googletag.pubads().setTargeting("grsession", "osid.6843cb11ed6cc2279d45ce3672ed8836");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "true");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", ["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"]);
    googletag.pubads().setTargeting("genres", ["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"]);
    googletag.pubads().setTargeting("Gender", "null");
    googletag.pubads().setTargeting("Age", "null");
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
<meta name="csrf-token" content="jPfrD7wwwS/2thQy7rBAt2q9aQDr08hYTR0vOa0KjKVW6oimqPNPJiLR1bhR9udl67yk6mBeZMUoXtPv8xvXHw==" />

  <meta name="request-id" content="75CBQNBWJAFRD799E0P4" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


    <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    var VIEW = 'table';
    var EDITABLE_USER_SHELF_NAME = 'read';
    var DRAGGABLE_REORDER = false;
    var VISIBLE_CONTROL = 'null';
    var INFINITE_SCROLL = false;
  //]]>
  </script>
  <script src="https://s.gr-assets.com/assets/review/list-848c7ab98d543929c014e94c55e6e268.js"></script>


  <link rel="alternate" type="application/atom+xml" title="Bookshelves" href="https://www.goodreads.com/review/list_rss/68156753?key=E4Ck6csfhyxS1_eAnt56W3p27ZGQ_-WjXpTyrdxgBXBexlfD&amp;shelf=%23ALL%23" />
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Sebastiaan has 361 books on their all shelf: The House in the Cerulean Sea by T.J. Klune and The Fox Wife by Yangsze Choo">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Sebastiaan’s books on Goodreads (361 books)' name='twitter:title'>
<meta content='Sebastiaan has 361 books on their all shelf: The House in the Cerulean Sea by T.J. Klune and The Fox Wife by Yangsze Choo' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".xo83nyz87c" data-react-checksum="-1472261889"></noscript></div>

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
        state: 'apple_oauth_state_c900811e-43f7-4dd6-9a27-7f96619060d0'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Freview%2Flist%2F68156753%3Fpage%3D1%26per_page%3D2%26view%3Dtable&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".1rmzgy3iyms" data-react-checksum="1010917207"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".1rmzgy3iyms.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".1rmzgy3iyms.0.0"><picture data-reactid=".1rmzgy3iyms.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".1rmzgy3iyms.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".1rmzgy3iyms.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".1rmzgy3iyms.1"><div class="siteHeader__contents" data-reactid=".1rmzgy3iyms.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".1rmzgy3iyms.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".1rmzgy3iyms.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".1rmzgy3iyms.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".1rmzgy3iyms.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1rmzgy3iyms.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1rmzgy3iyms.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1rmzgy3iyms.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1rmzgy3iyms.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1rmzgy3iyms.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.0"><span data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1rmzgy3iyms.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1rmzgy3iyms.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.0"><span data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".1rmzgy3iyms.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".1rmzgy3iyms.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".1rmzgy3iyms.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".1rmzgy3iyms.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".1rmzgy3iyms.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".1rmzgy3iyms.1.0.4"><ul class="personalNav" data-reactid=".1rmzgy3iyms.1.0.4.0"><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.4.0.0"><div data-reactid=".1rmzgy3iyms.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1rmzgy3iyms.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".1rmzgy3iyms.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".1rmzgy3iyms.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".1rmzgy3iyms.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.0"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753%3Fpage%3D1%26per_page%3D2%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".1rmzgy3iyms.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".1rmzgy3iyms.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".1rmzgy3iyms.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".1rmzgy3iyms.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".1rmzgy3iyms.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1rmzgy3iyms.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1rmzgy3iyms.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".1rmzgy3iyms.1.0.6"><div class="modal__content" data-reactid=".1rmzgy3iyms.1.0.6.0"><div class="modal__close" data-reactid=".1rmzgy3iyms.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".1rmzgy3iyms.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".1rmzgy3iyms.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".1rmzgy3iyms.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".1rmzgy3iyms.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".1rmzgy3iyms.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".1rmzgy3iyms.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".1rmzgy3iyms.1.0.7"><div data-reactid=".1rmzgy3iyms.1.0.7.0"><div class="modal__close" data-reactid=".1rmzgy3iyms.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".1rmzgy3iyms.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".1rmzgy3iyms.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".1rmzgy3iyms.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1"><ul data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753%3Fpage%3D1%26per_page%3D2%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".1rmzgy3iyms.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".1rmzgy3iyms.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".1rmzgy3iyms.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".1rmzgy3iyms.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".1rmzgy3iyms.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1rmzgy3iyms.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1rmzgy3iyms.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1rmzgy3iyms.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1rmzgy3iyms.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1rmzgy3iyms.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.0"><span data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1rmzgy3iyms.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1rmzgy3iyms.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.0"><span data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".1rmzgy3iyms.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
        <a href="/review/list/68156753?page=1&amp;per_page=2&amp;view=table">My Books</a>
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
<a id="batchEditLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this, {afterOpen: startEditing, afterClose: stopEditing});; return false;">Batch Edit</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753?page=1&amp;per_page=2&amp;print=true&amp;view=table">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753?page=1&amp;per_page=2&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753?page=1&amp;per_page=2&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
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
              Bookshelves <a class="smallText" href="/shelf/edit">(Edit)</a>
            </div>
            <a class="selectedShelf" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=read&amp;view=table">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=currently-reading&amp;view=table">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
    <a title="Sebastiaan&#39;s Want to Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;view=table">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
            </div>
          </div>
            <div class="stacked">
              <a class="gr-button gr-button--small" href="#" onclick="$(this).hide(); $(&#39;newShelfForm&#39;).show();; return false;">Add shelf</a>
              <div id="newShelfForm" style="display: none;" class="clearFix">
                <form class="titledBuilderForm gr-form gr-form--compact" id="shelf_name_form" action="/user_shelves" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><span class="formField name"><span class="labelDiv"><label class="gr-form--compact__label" for="user_shelf_name">Add a Shelf</label></span><span class="fieldDiv"><input size="18" maxlength="35" label_title="Add a Shelf" class="gr-form--compact__input" type="text" value="" name="user_shelf[name]" id="user_shelf_name" /></span></span>
<input type="submit" name="commit" value="add" class="gr-form--compact__submitButton" />
</form>
<script>
  //<![CDATA[
    $j(document).ready( function() {
      $j('#shelf_name_form')
          .bind('ajax:error', function () {
            alert("Shelf couldn't be created. Shelf name is either invalid or a duplicate.")
          })
          .bind('ajax:success', function () { document.location.reload(); } );
    });
  //]]>
</script>

              </div>
            </div>
            <div class="horizontalGreyDivider"></div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Your reading activity</div>
                <a href="/review/drafts">Review Drafts</a>
                <br/>
              <a class="annotatedBooksPageLink" href="/notes/68156753-sebastiaan?ref=rd">Kindle Notes &amp; Highlights</a>
              <br/>
              <a href="https://www.goodreads.com/challenges/11634">Reading Challenge</a>
              <br/>
              <a href="https://www.goodreads.com/user/year_in_books/2023/68156753">Year in Books</a>
              <br/>
              <a rel="nofollow" href="/review/stats/68156753-sebastiaan">Reading stats</a>
            </div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Add books</div>
              <br/>
              <a href="/recommendations">Recommendations</a>
              <br/>
              <a href="/book">Explore</a>
            </div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Tools</div>
              <a href="/review/duplicates">Find duplicates</a>
              <br/>
              <a rel="nofollow" href="/user/edit?tab=widgets">Widgets</a>
              <br/>
              <a href="/review/import">Import and export</a>
            </div>
            
        </div>
      </div>
    </div>
  <div id="rightCol" class="last col">
    <div id="shelfSettings" class="controlBody" style="display: none">
      <form id="fieldsForm" class="edit_user_shelf" action="/shelf/update/222519909" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="patch" />        <table>
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
                    <input type="checkbox" name="shelf[display_fields][position]" id="position_field" value="1" alt="position" />
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
                      <label for="shelves_field">shelves</label><br/>
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
          <div id="otherFields" style="margin-top: 10px">
            <label class="hlabel">other</label>
            <div class="formField per_page"><div class="labelDiv"><label for="user_shelf_per_page">Per page</label></div><div class="fieldDiv"><select name="user_shelf[per_page]" id="user_shelf_per_page"><option value=""></option>
<option value="10">10</option>
<option value="20">20</option>
<option value="30">30</option>
<option value="40">infinite scroll</option></select></div></div><div class="clear"></div>
            <div class="formField sort"><div class="labelDiv"><label for="user_shelf_sort">Sort</label></div><div class="fieldDiv"><select name="user_shelf[sort]" id="user_shelf_sort"><option value=""></option>
<option value="asin">Asin</option>
<option value="author">Author</option>
<option value="avg_rating">Avg rating</option>
<option value="cover">Cover</option>
<option value="date_added">Date added</option>
<option value="date_pub">Date pub</option>
<option value="date_pub_edition">Date pub edition</option>
<option value="date_read">Date read</option>
<option value="date_started">Date started</option>
<option value="date_updated">Date updated</option>
<option value="format">Format</option>
<option value="isbn">Isbn</option>
<option value="isbn13">Isbn13</option>
<option value="notes">Notes</option>
<option value="num_pages">Num pages</option>
<option value="num_ratings">Num ratings</option>
<option value="owned">Owned</option>
<option value="position">Position</option>
<option value="random">Random</option>
<option value="rating">Rating</option>
<option value="read_count">Read count</option>
<option value="review">Review</option>
<option value="title">Title</option>
<option value="votes">Votes</option>
<option value="year_pub">Year pub</option></select></div></div><div class="clear"></div>
            <input type="radio" value="a" name="user_shelf[order]" id="user_shelf_order_a" />
            <label for="shelf_order_a">ascending</label>
            <input type="radio" value="d" name="user_shelf[order]" id="user_shelf_order_d" />
            <label for="shelf_order_d">descending</label>
          </div>
          <div class="smallText buttons" style="margin-top: 10px">
            <input type="submit" name="commit" value="Save Current Settings to Your &quot;read/all&quot; shelves" id="save_curr_sett_submit" class="gr-button gr-button--small" style="margin-right: 10px" />
            <span class="loading" style="display: none"><img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Saving...</span>
              <span class="greyText inter uitext">shelf settings customized</span>
              <input type="checkbox" name="reset_display_fields" id="reset_display_fields" value="true" style="display:none" />
              <span class="greyText inter uitext">"read" and "all" shelves use the same settings</span>
            <span class="greyText status inter"></span>
          </div>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div id="batchEdit" style="display: none;" class="controlBody">
        <div id="shelfTools" class="toolset">
          <form name="reviewEditForm" id="reviewEditForm" action="/review/update_list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="iumssgQMOA/srvhN5CMZJ/gkPvn/CgjSW21ro3cx8MdQ9M8bEM+2BjjJOcdbZb71eSXzE3SHpE8+Lpd1KSCrfQ==" />
            <input type="hidden" name="view" id="view" value="table" />
            <label>
              Shelf:
              <select name="edit[shelf]" id="edit_shelf"><option value="read">read</option>
<option value="currently-reading">currently-reading</option>
<option value="to-read">to-read</option></select>
              &nbsp;
            </label>
            <a id="add_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;add_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_394473&#39;);$(&#39;add_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;add_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_394473&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_394473&#39;);Element.hide(&#39;add_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_394473&#39;);Element.show(&#39;add_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=add_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;xkDJ2BvY+v8wZIss6mxyj4+tKbEeJiLfvt6OQVQvK2wcXapxDxt09uQDSqZVKtVdDqzkW5WrjkLbnXKXCj5w1g==&#39;)}); return false;">add books to this shelf</a><img style="display:none" id="loading_anim_394473" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_295886&#39;);$(&#39;remove_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_295886&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_295886&#39;);Element.hide(&#39;remove_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_295886&#39;);Element.show(&#39;remove_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=remove_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;w1NRClywsZDmW5nYoLv374ero8335qTsdf82nvuqgO0ZTjKjSHM/mTI8WFIf/VA9BqpuJ3xrCHEQvMpIpbvbVw==&#39;)}); return false;">remove books from this shelf</a><img style="display:none" id="loading_anim_295886" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_books_link" class="actionLinkLite" href="#" onclick="if (confirm(&#39;This will completely remove the selected books from your shelves.&#39;)) { new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_books_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_719894&#39;);$(&#39;remove_books_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_books_link&#39;).show();;Element.hide(&#39;loading_anim_719894&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_719894&#39;);Element.hide(&#39;remove_books_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_719894&#39;);Element.show(&#39;remove_books_link&#39;);              for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {
                var r = request.responseJSON.reviews[i];
                $(&#39;review_&#39;+r.object.id).fade();
              }
}, parameters:(&#39;form_action=remove_books&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;SI62y9J7L3wvTkblbOFIUTgv45Y5XgHbt2K+Fw0ycsSSk9Vixrihdfsph2/Tp++DuS4ufLLTrUbSIULBUyMpfg==&#39;)}); }; return false;">remove books from all shelves</a><img style="display:none" id="loading_anim_719894" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
</form>        </div>
        <div id="otherTools" class="toolset greyText">
          <div class="right">
            <a class="actionLinkLite smallText" href="/review/duplicates">find duplicates</a>
          </div>
          <a class="actionLinkLite smallText" href="#" onclick="selectAllReviews(); return false;">select all</a>
          <span class="greyText">|</span>
          <a class="actionLinkLite smallText" href="#" onclick="unSelectAllReviews(); return false;">select none</a>
          <div class="clear"></div>
          <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;batchEditLink&#39;)); return false;">close</a>
          <div class="clear"></div>
        </div>
      </div>
      <div id="reorderConfirm" class="box noticeBox" style="display: none">
        <a id="loading_link_106564" class="button" href="#" onclick="new Ajax.Request(&#39;/shelf/move_batch/68156753&#39;, {asynchronous:true, evalScripts:true, onComplete:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;hide&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;show&#39;);}, onFailure:function(request){alert(&#39;Something went wrong re-ordering those shelves.&#39;);;Element.hide(&#39;loading_anim_106564&#39;);}, onLoading:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;show&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;hide&#39;);;Element.show(&#39;loading_anim_106564&#39;);Element.hide(&#39;loading_link_106564&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_106564&#39;);Element.show(&#39;loading_link_106564&#39;);$(&#39;booksBody&#39;).update(request.responseJSON.html);toggleFieldsToMatchHeader();startEditing();$(&#39;reorderConfirm&#39;).hide();$(&#39;booksBody&#39;).highlight();}, parameters:Form.serializeElements($$(&#39;#books .position input&#39;)) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;CfMgjhtSJeEZA4OlRBGbTiF3iJSgSg3HAotv9NfMcDXT7kMnD5Gr6M1kQi/7VzycoHZFfivHoVpnyJMiid0rjw==&#39;)}); return false;">apply position changes?</a><img style="display:none" id="loading_anim_106564" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          &nbsp;
          <a href="#" onclick="$(&#39;reorderConfirm&#39;).hide(); return false;">Not yet</a>
      </div>
      <div class="right uitext">
        <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753?page=2&amp;per_page=2&amp;view=table">2</a> <a href="/review/list/68156753?page=3&amp;per_page=2&amp;view=table">3</a> <a href="/review/list/68156753?page=4&amp;per_page=2&amp;view=table">4</a> <a href="/review/list/68156753?page=5&amp;per_page=2&amp;view=table">5</a> <a href="/review/list/68156753?page=6&amp;per_page=2&amp;view=table">6</a> <a href="/review/list/68156753?page=7&amp;per_page=2&amp;view=table">7</a> <a href="/review/list/68156753?page=8&amp;per_page=2&amp;view=table">8</a> <a href="/review/list/68156753?page=9&amp;per_page=2&amp;view=table">9</a> <span class="gap">&hellip;</span> <a href="/review/list/68156753?page=180&amp;per_page=2&amp;view=table">180</a> <a href="/review/list/68156753?page=181&amp;per_page=2&amp;view=table">181</a> <a class="next_page" rel="next" href="/review/list/68156753?page=2&amp;per_page=2&amp;view=table">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="display: none">
              </th>
              <th alt="position" class="header field position" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=position&amp;view=table">#</a>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=cover&amp;view=table">cover</a>
              </th>
              <th alt="title" class="header field title" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=title&amp;view=table">title</a>
              </th>
              <th alt="author" class="header field author" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=author&amp;view=table">author</a>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=isbn&amp;view=table">isbn</a>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=isbn13&amp;view=table">isbn13</a>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=asin&amp;view=table">asin</a>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=num_pages&amp;view=table">num pages</a>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=avg_rating&amp;view=table">avg rating</a>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=num_ratings&amp;view=table">num ratings</a>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=date_pub&amp;view=table">date pub</a>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=date_pub_edition&amp;view=table">date pub (ed.)</a>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=rating&amp;view=table">rating</a>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    shelves
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=review&amp;view=table">review</a>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=notes&amp;view=table">notes</a>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=comments&amp;view=table">comments</a>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=votes&amp;view=table">votes</a>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=read_count&amp;view=table">read count</a>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=date_started&amp;view=table">date started</a>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <a href="/review/list/68156753?per_page=2&amp;sort=date_read&amp;view=table">date read</a>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <a href="/review/list/68156753?order=a&amp;per_page=2&amp;sort=date_added&amp;view=table">date</a>
                    <a href="/review/list/68156753?order=a&amp;per_page=2&amp;sort=date_added&amp;view=table">
                      <nobr>
                        added <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                      </nobr>
</a>              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=owned&amp;view=table">owned</a>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <a href="/review/list/68156753?per_page=2&amp;sort=format&amp;view=table">format</a>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_7064093266" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7064093266]" id="checkbox_review_7064093266" value="7064093266" />
</div></td>  <td class="field position" style="display: none"><label>position</label><div class="value"></div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="45047384">
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    742,197
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Mar 16, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 17, 2020
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="45047384" data-user-id="68156753" data-submit-url="/review/rate/45047384?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage45047384_68156753"></span>
        <span id="successMessage45047384_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_45047384"><span id="shelf_6736083077"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 45047384, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/45047384?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 45047384, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 45047384, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 45047384, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 07, 2024">
    Dec 07, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/62945242">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_427161" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/45047384&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_427161&#39;);$(&#39;loading_link_427161&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_427161&#39;).show();;Element.hide(&#39;loading_anim_427161&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_427161&#39;);Element.hide(&#39;loading_link_427161&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_427161&#39;);Element.show(&#39;loading_link_427161&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;8+i4aZEG1ILLXw/mJm0IUZeO8tJrOIlQj4I5E3bl/Y8p9dvAhcVaix84zmyZK6+DFo8/OOC1Jc3qwcXFKPSmNQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_427161" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7064093266">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The House in the Cerulean Sea from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/45047384?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753%3Fpage%3D1%26per_page%3D2%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_7009239588" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7009239588]" id="checkbox_review_7009239588" value="7009239588" />
</div></td>  <td class="field position" style="display: none"><label>position</label><div class="value"></div></td>  <td class="field cover"><label>cover</label><div class="value">        <div class="js-tooltipTrigger tooltipTrigger" data-resource-type="Book" data-resource-id="127278666">
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
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.00
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    13,839
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 13, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 13, 2024
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="127278666" data-user-id="68156753" data-submit-url="/review/rate/127278666?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage127278666_68156753"></span>
        <span id="successMessage127278666_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_127278666"><span id="shelf_6679332637"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 127278666, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/127278666?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 127278666, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 127278666, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 127278666, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 16, 2024">
    Nov 16, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/148387285">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_266206" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/127278666&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_266206&#39;);$(&#39;loading_link_266206&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_266206&#39;).show();;Element.hide(&#39;loading_anim_266206&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_266206&#39;);Element.hide(&#39;loading_link_266206&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_266206&#39;);Element.show(&#39;loading_link_266206&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;mm3enzTlXDDhCujT5PuuL9B2iWXP0UNxPMwCbfecdT9AcL02ICbSOTVtKVlbvQn9UXdEj0Rc7+xZj/67qY0uhQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_266206" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7009239588">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Fox Wife from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/127278666?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753%3Fpage%3D1%26per_page%3D2%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
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
            2 of 361 loaded
          </div>
            <form id="perPageForm" name="perPageForm" class="inter" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />                <input type="hidden" name="view" id="view" value="table" />
              <label for="per_page" class="greyText">per page</label>
                <select name="per_page" id="per_page" onchange="$(&#39;perPageForm&#39;).submit()"><option value="10">10</option>
<option value="20">20</option>
<option value="30">30</option>
<option value="40">40</option>
<option value="50">50</option>
<option value="75">75</option>
<option value="100">100</option>
<option value="infinite">infinite scroll</option></select>
</form>          <form id="sortForm" name="sortForm" class="inter" action="/review/list/68156753-sebastiaan" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />              <input type="hidden" name="view" id="view" value="table" />
              <label for="sort" class="greyText">sort</label>
              <select name="sort" id="sort" onchange="$(&#39;sortForm&#39;).submit()"><option value="asin">Asin</option>
<option value="author">Author</option>
<option value="avg_rating">Avg rating</option>
<option value="cover">Cover</option>
<option selected="selected" value="date_added">Date added</option>
<option value="date_pub">Date pub</option>
<option value="date_pub_edition">Date pub edition</option>
<option value="date_read">Date read</option>
<option value="date_started">Date started</option>
<option value="date_updated">Date updated</option>
<option value="format">Format</option>
<option value="isbn">Isbn</option>
<option value="isbn13">Isbn13</option>
<option value="notes">Notes</option>
<option value="num_pages">Num pages</option>
<option value="num_ratings">Num ratings</option>
<option value="owned">Owned</option>
<option value="position">Position</option>
<option value="random">Random</option>
<option value="rating">Rating</option>
<option value="read_count">Read count</option>
<option value="review">Review</option>
<option value="title">Title</option>
<option value="votes">Votes</option>
<option value="year_pub">Year pub</option></select>
              <input type="radio" name="order" id="order_a" value="a" onchange="$(&#39;sortForm&#39;).submit()" /> <label for="order_a">asc.</label>
              <input type="radio" name="order" id="order_d" value="d" onchange="$(&#39;sortForm&#39;).submit()" checked="checked" /> <label for="order_a">desc.</label>
            <a href="https://www.goodreads.com/review/list_rss/68156753?key=E4Ck6csfhyxS1_eAnt56W3p27ZGQ_-WjXpTyrdxgBXBexlfD&amp;shelf=%23ALL%23"><img style="vertical-align: middle" class="inter" src="https://s.gr-assets.com/assets/links/rss_infinite-2e37dd81d44bab27eb8fdbf3bb5d9973.gif" alt="Rss infinite" /></a>
</form>          <div class="inter">
            <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753?page=2&amp;per_page=2&amp;view=table">2</a> <a href="/review/list/68156753?page=3&amp;per_page=2&amp;view=table">3</a> <a href="/review/list/68156753?page=4&amp;per_page=2&amp;view=table">4</a> <a href="/review/list/68156753?page=5&amp;per_page=2&amp;view=table">5</a> <a href="/review/list/68156753?page=6&amp;per_page=2&amp;view=table">6</a> <a href="/review/list/68156753?page=7&amp;per_page=2&amp;view=table">7</a> <a href="/review/list/68156753?page=8&amp;per_page=2&amp;view=table">8</a> <a href="/review/list/68156753?page=9&amp;per_page=2&amp;view=table">9</a> <span class="gap">&hellip;</span> <a href="/review/list/68156753?page=180&amp;per_page=2&amp;view=table">180</a> <a href="/review/list/68156753?page=181&amp;per_page=2&amp;view=table">181</a> <a class="next_page" rel="next" href="/review/list/68156753?page=2&amp;per_page=2&amp;view=table">next »</a></div>

          </div>
        </div>
      </div>
  </div>
  <div class="clear"></div>
</div>

  <input type="text" id="shelfChooserInput" />
  <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    
function createWindowShelfChooser() {
  if (window.shelfChooser != false) { return; }
  if ($('shelfChooserInput') == null) {
    document.body.appendChild(new Element('input', {type: 'text', id: 'shelfChooserInput'}));
  }
  window.shelfChooser = new ShelfChooser(
    "shelfChooserInput",
    0,
    ["read", "currently-reading", "to-read"],
    {
      chosen: ["to-read"],
      exclusive: ["read", "currently-reading", "to-read"],
      cacheChosen: true,
      afterClose: function(chooser) {chooser.wrapper.hide();},
      afterChoose: function(shelfName, chooser) {
        if (chooser.bookId == 0) {return false};
        if ($("shelfList68156753_" + chooser.bookId) == null) {
          return false;
        }
        var shelfLinks = chooser.chosen.map(function(chosenShelfName) {
          return '<a href="/review/list/68156753?shelf='+
            chosenShelfName+'" class="shelfLink">'+chosenShelfName+'</a>';
        });
        $("shelfList68156753_" + chooser.bookId).update(
          shelfLinks.join(', ')
        );
      }
    }
  );
  window.shelfChooser.wrapper.setStyle({'position': 'absolute'});
  window.shelfChooser.wrapper.hide();
}

if (typeof(window.shelfChooser) == 'undefined') {
  window.shelfChooser = false;
  if (document.loaded) {
    createWindowShelfChooser()
  } else {
    document.observe('dom:loaded', createWindowShelfChooser)
  }
}

  //]]>
  </script>
  <div id="reviewForm" class="floatingBox modelEditor" style="display:none;">
    <form onsubmit="reviewEditor.save(); return false;" action="/review/list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="sip4+U5qihPRGoVWId412JP354USEEBV2WUgJG5cK4NoNxtQWqkEGgV9RNyemJIKEvYqb5md7Mi8JtzyME1wOQ==" />      <input type="hidden" name="reading_session_id" id="reading_session_id" />
      <div class="formField review_usertext"><div class="labelDiv"><label for="review_review_usertext">review</label></div><div class="fieldDiv"><textarea label_title="review" name="review[review_usertext]" id="review_review_usertext" cols="40" rows="20">
</textarea></div></div><div class="clear"></div>
      <div class="formField notes"><div class="labelDiv"><label for="review_notes">Notes</label></div><div class="fieldDiv"><textarea name="review[notes]" id="review_notes" cols="40" rows="20">
</textarea></div></div><div class="clear"></div>
      <div class="formField read_at"><div class="labelDiv"><label for="review_read_at">Read at</label></div><div class="fieldDiv"><label for='review_read_at_1i'>Year: </label><select id="review_read_at_1i" name="review[read_at(1i)]">
<option value=""></option>
<option value="2024">2024</option>
<option value="2023">2023</option>
<option value="2022">2022</option>
<option value="2021">2021</option>
<option value="2020">2020</option>
<option value="2019">2019</option>
<option value="2018">2018</option>
<option value="2017">2017</option>
<option value="2016">2016</option>
<option value="2015">2015</option>
<option value="2014">2014</option>
<option value="2013">2013</option>
<option value="2012">2012</option>
<option value="2011">2011</option>
<option value="2010">2010</option>
<option value="2009">2009</option>
<option value="2008">2008</option>
<option value="2007">2007</option>
<option value="2006">2006</option>
<option value="2005">2005</option>
<option value="2004">2004</option>
<option value="2003">2003</option>
<option value="2002">2002</option>
<option value="2001">2001</option>
<option value="2000">2000</option>
<option value="1999">1999</option>
<option value="1998">1998</option>
<option value="1997">1997</option>
<option value="1996">1996</option>
<option value="1995">1995</option>
<option value="1994">1994</option>
<option value="1993">1993</option>
<option value="1992">1992</option>
<option value="1991">1991</option>
<option value="1990">1990</option>
<option value="1989">1989</option>
<option value="1988">1988</option>
<option value="1987">1987</option>
<option value="1986">1986</option>
<option value="1985">1985</option>
<option value="1984">1984</option>
<option value="1983">1983</option>
<option value="1982">1982</option>
<option value="1981">1981</option>
<option value="1980">1980</option>
<option value="1979">1979</option>
<option value="1978">1978</option>
<option value="1977">1977</option>
<option value="1976">1976</option>
<option value="1975">1975</option>
<option value="1974">1974</option>
<option value="1973">1973</option>
<option value="1972">1972</option>
<option value="1971">1971</option>
<option value="1970">1970</option>
<option value="1969">1969</option>
<option value="1968">1968</option>
<option value="1967">1967</option>
<option value="1966">1966</option>
<option value="1965">1965</option>
<option value="1964">1964</option>
<option value="1963">1963</option>
<option value="1962">1962</option>
<option value="1961">1961</option>
<option value="1960">1960</option>
<option value="1959">1959</option>
<option value="1958">1958</option>
<option value="1957">1957</option>
<option value="1956">1956</option>
<option value="1955">1955</option>
<option value="1954">1954</option>
<option value="1953">1953</option>
<option value="1952">1952</option>
<option value="1951">1951</option>
<option value="1950">1950</option>
<option value="1949">1949</option>
<option value="1948">1948</option>
<option value="1947">1947</option>
<option value="1946">1946</option>
<option value="1945">1945</option>
<option value="1944">1944</option>
<option value="1943">1943</option>
<option value="1942">1942</option>
<option value="1941">1941</option>
<option value="1940">1940</option>
<option value="1939">1939</option>
<option value="1938">1938</option>
<option value="1937">1937</option>
<option value="1936">1936</option>
<option value="1935">1935</option>
<option value="1934">1934</option>
<option value="1933">1933</option>
<option value="1932">1932</option>
<option value="1931">1931</option>
<option value="1930">1930</option>
<option value="1929">1929</option>
<option value="1928">1928</option>
<option value="1927">1927</option>
<option value="1926">1926</option>
<option value="1925">1925</option>
<option value="1924">1924</option>
</select>
<label for='review_read_at_2i'>&nbsp&nbsp Month: </label><select id="review_read_at_2i" name="review[read_at(2i)]">
<option value=""></option>
<option value="1">January</option>
<option value="2">February</option>
<option value="3">March</option>
<option value="4">April</option>
<option value="5">May</option>
<option value="6">June</option>
<option value="7">July</option>
<option value="8">August</option>
<option value="9">September</option>
<option value="10">October</option>
<option value="11">November</option>
<option value="12">December</option>
</select>
<label for='review_read_at_3i'>&nbsp&nbsp Day: </label><select id="review_read_at_3i" name="review[read_at(3i)]">
<option value=""></option>
<option value="1">1</option>
<option value="2">2</option>
<option value="3">3</option>
<option value="4">4</option>
<option value="5">5</option>
<option value="6">6</option>
<option value="7">7</option>
<option value="8">8</option>
<option value="9">9</option>
<option value="10">10</option>
<option value="11">11</option>
<option value="12">12</option>
<option value="13">13</option>
<option value="14">14</option>
<option value="15">15</option>
<option value="16">16</option>
<option value="17">17</option>
<option value="18">18</option>
<option value="19">19</option>
<option value="20">20</option>
<option value="21">21</option>
<option value="22">22</option>
<option value="23">23</option>
<option value="24">24</option>
<option value="25">25</option>
<option value="26">26</option>
<option value="27">27</option>
<option value="28">28</option>
<option value="29">29</option>
<option value="30">30</option>
<option value="31">31</option>
</select>
<a class="actionLinkLite setToTodayLink" style="margin-left: 0.5em" href="#" onclick="setReadAt({field_id_prefix: &#39;review_read_at&#39;, link: this}); return false;">set to today</a></div></div><div class="clear"></div>
      <div class="formField started_at"><div class="labelDiv"><label for="review_started_at">Started at</label></div><div class="fieldDiv"><label for='review_started_at_1i'>Year: </label><select id="review_started_at_1i" name="review[started_at(1i)]">
<option value=""></option>
<option value="2024">2024</option>
<option value="2023">2023</option>
<option value="2022">2022</option>
<option value="2021">2021</option>
<option value="2020">2020</option>
<option value="2019">2019</option>
<option value="2018">2018</option>
<option value="2017">2017</option>
<option value="2016">2016</option>
<option value="2015">2015</option>
<option value="2014">2014</option>
<option value="2013">2013</option>
<option value="2012">2012</option>
<option value="2011">2011</option>
<option value="2010">2010</option>
<option value="2009">2009</option>
<option value="2008">2008</option>
<option value="2007">2007</option>
<option value="2006">2006</option>
<option value="2005">2005</option>
<option value="2004">2004</option>
<option value="2003">2003</option>
<option value="2002">2002</option>
<option value="2001">2001</option>
<option value="2000">2000</option>
<option value="1999">1999</option>
<option value="1998">1998</option>
<option value="1997">1997</option>
<option value="1996">1996</option>
<option value="1995">1995</option>
<option value="1994">1994</option>
<option value="1993">1993</option>
<option value="1992">1992</option>
<option value="1991">1991</option>
<option value="1990">1990</option>
<option value="1989">1989</option>
<option value="1988">1988</option>
<option value="1987">1987</option>
<option value="1986">1986</option>
<option value="1985">1985</option>
<option value="1984">1984</option>
<option value="1983">1983</option>
<option value="1982">1982</option>
<option value="1981">1981</option>
<option value="1980">1980</option>
<option value="1979">1979</option>
<option value="1978">1978</option>
<option value="1977">1977</option>
<option value="1976">1976</option>
<option value="1975">1975</option>
<option value="1974">1974</option>
<option value="1973">1973</option>
<option value="1972">1972</option>
<option value="1971">1971</option>
<option value="1970">1970</option>
<option value="1969">1969</option>
<option value="1968">1968</option>
<option value="1967">1967</option>
<option value="1966">1966</option>
<option value="1965">1965</option>
<option value="1964">1964</option>
<option value="1963">1963</option>
<option value="1962">1962</option>
<option value="1961">1961</option>
<option value="1960">1960</option>
<option value="1959">1959</option>
<option value="1958">1958</option>
<option value="1957">1957</option>
<option value="1956">1956</option>
<option value="1955">1955</option>
<option value="1954">1954</option>
<option value="1953">1953</option>
<option value="1952">1952</option>
<option value="1951">1951</option>
<option value="1950">1950</option>
<option value="1949">1949</option>
<option value="1948">1948</option>
<option value="1947">1947</option>
<option value="1946">1946</option>
<option value="1945">1945</option>
<option value="1944">1944</option>
<option value="1943">1943</option>
<option value="1942">1942</option>
<option value="1941">1941</option>
<option value="1940">1940</option>
<option value="1939">1939</option>
<option value="1938">1938</option>
<option value="1937">1937</option>
<option value="1936">1936</option>
<option value="1935">1935</option>
<option value="1934">1934</option>
<option value="1933">1933</option>
<option value="1932">1932</option>
<option value="1931">1931</option>
<option value="1930">1930</option>
<option value="1929">1929</option>
<option value="1928">1928</option>
<option value="1927">1927</option>
<option value="1926">1926</option>
<option value="1925">1925</option>
<option value="1924">1924</option>
</select>
<label for='review_started_at_2i'>&nbsp&nbsp Month: </label><select id="review_started_at_2i" name="review[started_at(2i)]">
<option value=""></option>
<option value="1">January</option>
<option value="2">February</option>
<option value="3">March</option>
<option value="4">April</option>
<option value="5">May</option>
<option value="6">June</option>
<option value="7">July</option>
<option value="8">August</option>
<option value="9">September</option>
<option value="10">October</option>
<option value="11">November</option>
<option value="12">December</option>
</select>
<label for='review_started_at_3i'>&nbsp&nbsp Day: </label><select id="review_started_at_3i" name="review[started_at(3i)]">
<option value=""></option>
<option value="1">1</option>
<option value="2">2</option>
<option value="3">3</option>
<option value="4">4</option>
<option value="5">5</option>
<option value="6">6</option>
<option value="7">7</option>
<option value="8">8</option>
<option value="9">9</option>
<option value="10">10</option>
<option value="11">11</option>
<option value="12">12</option>
<option value="13">13</option>
<option value="14">14</option>
<option value="15">15</option>
<option value="16">16</option>
<option value="17">17</option>
<option value="18">18</option>
<option value="19">19</option>
<option value="20">20</option>
<option value="21">21</option>
<option value="22">22</option>
<option value="23">23</option>
<option value="24">24</option>
<option value="25">25</option>
<option value="26">26</option>
<option value="27">27</option>
<option value="28">28</option>
<option value="29">29</option>
<option value="30">30</option>
<option value="31">31</option>
</select>
<a class="actionLinkLite setToTodayLink" style="margin-left: 0.5em" href="#" onclick="setReadAt({field_id_prefix: &#39;review_started_at&#39;, link: this}); return false;">set to today</a></div></div><div class="clear"></div>
      <div class="buttons uitext">
        <a class="greyText right" tabindex="2" href="#" onclick="$(&#39;reviewForm&#39;).hideFloatingBox(); return false;">close</a>
        <a class="gr-button gr-button--small" tabindex="0" href="#" onclick="reviewEditor.save(); return false;">Save</a>
        <span class="loading" style="display:none"><img src="/assets/loading-trans.gif" /> saving...</span>
        &nbsp;
        <a tabindex="1" href="#" onclick="$(&#39;reviewForm&#39;).hideFloatingBox(); return false;">cancel</a>
      </div>
</form>  </div>


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
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".218b5dwghyi" data-react-checksum="-1283845832"></noscript></div>


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
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.6843cb11ed6cc2279d45ce3672ed8836","grsession":"osid.6843cb11ed6cc2279d45ce3672ed8836","surface":"desktop","signedin":"true","gr_author":"false","author":["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"],"genres":["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"],"Gender":"null","Age":"null"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({"unreadCount":2,"unreadCountMore":false});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":{"name":"Sebastiaan","profileUrl":"/user/show/68156753-sebastiaan","profileImage":"https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png","pendingRecsCount":0,"groupInvitesCount":0,"tempFriendRequestCount":0,"tempUnreadMessageCount":0}});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":["Crime","Fantasy","Fiction","Historical fiction","Mystery","Science fiction","Thriller"]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: ovyrfolxonosxkuloxukyenkamichhfigmqxosfwmiyzhwvduyvcqcuqszgtinrtvwrjblwnnpbuadpgupdlcfqfnqqkoumoamdtypbvdemzmfxcixnkccmdwftxwnjakpqcidlrpvxyffanuzmkjqamplxbierzbkrhtshnmeaxwkoyekbgzjfcihvtztmknwcfczvfkwxeftbihfyyvenfurnmkjrjrhykblddqrlicysetpsjflzpwregzstwfqxprqtsrdwnqtxjzugrapiziqqievcnryviurkxfrcqeeapmqkmtmtshphfpfmfzxtrtzmwihfikpoeylcyfgidifyteregjsyiaodndgjtumqtvigqoszzvnmhbsiirrewbvrrfrwhrkcjnnwhpamoptoqjrglozsgbm -->`
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML))
		}))
		defer server.Close()

		want := []book{{
			title:       "The House in the Cerulean Sea (Cerulean Chronicles, #1)",
			author:      "Klune, T.J.",
			workUrl:     "/book/show/45047384",
			editionsUrl: "/work/editions/62945242",
		}, {
			title:       "The Fox Wife",
			author:      "Choo, Yangsze",
			isbn:        "1250266017",
			workUrl:     "/book/show/127278666",
			editionsUrl: "/work/editions/148387285",
		}}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books: \nWant: '%+v'\n Got: '%+v'\n", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("\nWant: '%+v'\n Got: '%+v'\n", want, got)
		}
	})

	t.Run("test page with no books", func(t *testing.T) {
		mockHTML := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Sebastiaan’s books on Goodreads (361 books)</title>

<meta content='Sebastiaan has 361 books on their all shelf, but has 367 books shelved. ' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/review/list/68156753' rel='canonical'>
  <meta property="og:title" content="Sebastiaan’s books on Goodreads (361 books)"/>
  <meta property="og:type" content="website"/>
  <meta property="og:site_name" content="Goodreads"/>
  <meta property="og:description" content="Sebastiaan has 361 books on their all shelf, but has 367 books shelved. "/>
  <meta property="og:url" content="https://www.goodreads.com/review/list/68156753">
  <meta property="fb:app_id" content="2415071772"/>



    <script type="text/javascript"> var ue_t0=window.ue_t0||+new Date();
 </script>
  <script type="text/javascript">
    var ue_mid = "A1PQBFHBHS6YH1";
    var ue_sn = "www.goodreads.com";
    var ue_furl = "fls-na.amazon.com";
    var ue_sid = "854-2249445-5127461";
    var ue_id = "7MNDYR740BVXQ9DG26CN";

    (function(e){var c=e;var a=c.ue||{};a.main_scope="mainscopecsm";a.q=[];a.t0=c.ue_t0||+new Date();a.d=g;function g(h){return +new Date()-(h?0:a.t0)}function d(h){return function(){a.q.push({n:h,a:arguments,t:a.d()})}}function b(m,l,h,j,i){var k={m:m,f:l,l:h,c:""+j,err:i,fromOnError:1,args:arguments};c.ueLogError(k);return false}b.skipTrace=1;e.onerror=b;function f(){c.uex("ld")}if(e.addEventListener){e.addEventListener("load",f,false)}else{if(e.attachEvent){e.attachEvent("onload",f)}}a.tag=d("tag");a.log=d("log");a.reset=d("rst");c.ue_csm=c;c.ue=a;c.ueLogError=d("err");c.ues=d("ues");c.uet=d("uet");c.uex=d("uex");c.uet("ue")})(window);(function(e,d){var a=e.ue||{};function c(g){if(!g){return}var f=d.head||d.getElementsByTagName("head")[0]||d.documentElement,h=d.createElement("script");h.async="async";h.src=g;f.insertBefore(h,f.firstChild)}function b(){var k=e.ue_cdn||"z-ecx.images-amazon.com",g=e.ue_cdns||"images-na.ssl-images-amazon.com",j="/images/G/01/csminstrumentation/",h=e.ue_file||"ue-full-11e51f253e8ad9d145f4ed644b40f692._V1_.js",f,i;if(h.indexOf("NSTRUMENTATION_FIL")>=0){return}if("ue_https" in e){f=e.ue_https}else{f=e.location&&e.location.protocol=="https:"?1:0}i=f?"https://":"http://";i+=f?g:k;i+=j;i+=h;c(i)}if(!e.ue_inline){if(a.loadUEFull){a.loadUEFull()}else{b()}}a.uels=c;e.ue=a})(window,document);

    if (window.ue && window.ue.tag) { window.ue.tag('review:list:signed_in', ue.main_scope);window.ue.tag('review:list:signed_in:desktop', ue.main_scope); }
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
        "CacheDetection.RequestID": "7MNDYR740BVXQ9DG26CN",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "854-2249445-5127461" },
        page: {requestId: "7MNDYR740BVXQ9DG26CN", meaningful: "interactive"}
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
      googletag.pubads().setTargeting("sid", "osid.6843cb11ed6cc2279d45ce3672ed8836");
    googletag.pubads().setTargeting("grsession", "osid.6843cb11ed6cc2279d45ce3672ed8836");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "true");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", ["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"]);
    googletag.pubads().setTargeting("genres", ["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"]);
    googletag.pubads().setTargeting("Gender", "null");
    googletag.pubads().setTargeting("Age", "null");
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
<meta name="csrf-token" content="ymqh0DSHN+uZKg/sWo0Mg2sr/ZTN4/68lCwmsMyf5GAQd8J5IES54k1Nzmbly6tR6iowfkZuUiHxb9pmko6/2g==" />

  <meta name="request-id" content="7MNDYR740BVXQ9DG26CN" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


    <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    var VIEW = 'table';
    var EDITABLE_USER_SHELF_NAME = 'read';
    var DRAGGABLE_REORDER = false;
    var VISIBLE_CONTROL = 'null';
    var INFINITE_SCROLL = false;
  //]]>
  </script>
  <script src="https://s.gr-assets.com/assets/review/list-848c7ab98d543929c014e94c55e6e268.js"></script>


  <link rel="alternate" type="application/atom+xml" title="Bookshelves" href="https://www.goodreads.com/review/list_rss/68156753?key=E4Ck6csfhyxS1_eAnt56W3p27ZGQ_-WjXpTyrdxgBXBexlfD&amp;shelf=%23ALL%23" />
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Sebastiaan has 361 books on their all shelf, but has 367 books shelved. ">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Sebastiaan’s books on Goodreads (361 books)' name='twitter:title'>
<meta content='Sebastiaan has 361 books on their all shelf, but has 367 books shelved. ' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".22a8b4ejknc" data-react-checksum="-1312878312"></noscript></div>

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
        state: 'apple_oauth_state_c900811e-43f7-4dd6-9a27-7f96619060d0'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Freview%2Flist%2F68156753%3Fpage%3D21%26per_page%3D20&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".1qk8j4y1a4a" data-react-checksum="853215868"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".1qk8j4y1a4a.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".1qk8j4y1a4a.0.0"><picture data-reactid=".1qk8j4y1a4a.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".1qk8j4y1a4a.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".1qk8j4y1a4a.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".1qk8j4y1a4a.1"><div class="siteHeader__contents" data-reactid=".1qk8j4y1a4a.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".1qk8j4y1a4a.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".1qk8j4y1a4a.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".1qk8j4y1a4a.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".1qk8j4y1a4a.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1qk8j4y1a4a.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1qk8j4y1a4a.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1qk8j4y1a4a.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1qk8j4y1a4a.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1qk8j4y1a4a.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".1qk8j4y1a4a.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".1qk8j4y1a4a.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".1qk8j4y1a4a.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".1qk8j4y1a4a.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".1qk8j4y1a4a.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".1qk8j4y1a4a.1.0.4"><ul class="personalNav" data-reactid=".1qk8j4y1a4a.1.0.4.0"><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.4.0.0"><div data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1qk8j4y1a4a.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".1qk8j4y1a4a.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".1qk8j4y1a4a.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".1qk8j4y1a4a.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753%3Fpage%3D21%26per_page%3D20" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".1qk8j4y1a4a.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".1qk8j4y1a4a.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".1qk8j4y1a4a.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".1qk8j4y1a4a.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".1qk8j4y1a4a.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1qk8j4y1a4a.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1qk8j4y1a4a.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".1qk8j4y1a4a.1.0.6"><div class="modal__content" data-reactid=".1qk8j4y1a4a.1.0.6.0"><div class="modal__close" data-reactid=".1qk8j4y1a4a.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".1qk8j4y1a4a.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".1qk8j4y1a4a.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".1qk8j4y1a4a.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".1qk8j4y1a4a.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".1qk8j4y1a4a.1.0.7"><div data-reactid=".1qk8j4y1a4a.1.0.7.0"><div class="modal__close" data-reactid=".1qk8j4y1a4a.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".1qk8j4y1a4a.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".1qk8j4y1a4a.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".1qk8j4y1a4a.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1"><ul data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753%3Fpage%3D21%26per_page%3D20" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".1qk8j4y1a4a.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".1qk8j4y1a4a.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".1qk8j4y1a4a.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".1qk8j4y1a4a.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".1qk8j4y1a4a.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".1qk8j4y1a4a.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".1qk8j4y1a4a.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".1qk8j4y1a4a.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.0"><span data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".1qk8j4y1a4a.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".1qk8j4y1a4a.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.0"><span data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".1qk8j4y1a4a.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
        <a href="/review/list/68156753?page=1&amp;per_page=20">My Books</a>
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
<a id="batchEditLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this, {afterOpen: startEditing, afterClose: stopEditing});; return false;">Batch Edit</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753?page=21&amp;per_page=20&amp;print=true">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753?page=21&amp;per_page=20&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753?page=21&amp;per_page=20&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
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
              Bookshelves <a class="smallText" href="/shelf/edit">(Edit)</a>
            </div>
            <a class="selectedShelf" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=read">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=currently-reading">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
    <a title="Sebastiaan&#39;s Want to Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
            </div>
          </div>
            <div class="stacked">
              <a class="gr-button gr-button--small" href="#" onclick="$(this).hide(); $(&#39;newShelfForm&#39;).show();; return false;">Add shelf</a>
              <div id="newShelfForm" style="display: none;" class="clearFix">
                <form class="titledBuilderForm gr-form gr-form--compact" id="shelf_name_form" action="/user_shelves" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><span class="formField name"><span class="labelDiv"><label class="gr-form--compact__label" for="user_shelf_name">Add a Shelf</label></span><span class="fieldDiv"><input size="18" maxlength="35" label_title="Add a Shelf" class="gr-form--compact__input" type="text" value="" name="user_shelf[name]" id="user_shelf_name" /></span></span>
<input type="submit" name="commit" value="add" class="gr-form--compact__submitButton" />
</form>
<script>
  //<![CDATA[
    $j(document).ready( function() {
      $j('#shelf_name_form')
          .bind('ajax:error', function () {
            alert("Shelf couldn't be created. Shelf name is either invalid or a duplicate.")
          })
          .bind('ajax:success', function () { document.location.reload(); } );
    });
  //]]>
</script>

              </div>
            </div>
            <div class="horizontalGreyDivider"></div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Your reading activity</div>
                <a href="/review/drafts">Review Drafts</a>
                <br/>
              <a class="annotatedBooksPageLink" href="/notes/68156753-sebastiaan?ref=rd">Kindle Notes &amp; Highlights</a>
              <br/>
              <a href="https://www.goodreads.com/challenges/11634">Reading Challenge</a>
              <br/>
              <a href="https://www.goodreads.com/user/year_in_books/2023/68156753">Year in Books</a>
              <br/>
              <a rel="nofollow" href="/review/stats/68156753-sebastiaan">Reading stats</a>
            </div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Add books</div>
              <br/>
              <a href="/recommendations">Recommendations</a>
              <br/>
              <a href="/book">Explore</a>
            </div>
            <div id="toolsSection" class="actionLinkLites">
              <div class="sectionHeader">Tools</div>
              <a href="/review/duplicates">Find duplicates</a>
              <br/>
              <a rel="nofollow" href="/user/edit?tab=widgets">Widgets</a>
              <br/>
              <a href="/review/import">Import and export</a>
            </div>
            
        </div>
      </div>
    </div>
  <div id="rightCol" class="last col">
    <div id="shelfSettings" class="controlBody" style="display: none">
      <form id="fieldsForm" class="edit_user_shelf" action="/shelf/update/222519909" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="patch" />        <table>
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
                    <input type="checkbox" name="shelf[display_fields][position]" id="position_field" value="1" alt="position" />
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
                      <label for="shelves_field">shelves</label><br/>
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
          <div id="otherFields" style="margin-top: 10px">
            <label class="hlabel">other</label>
            <div class="formField per_page"><div class="labelDiv"><label for="user_shelf_per_page">Per page</label></div><div class="fieldDiv"><select name="user_shelf[per_page]" id="user_shelf_per_page"><option value=""></option>
<option value="10">10</option>
<option value="20">20</option>
<option value="30">30</option>
<option value="40">infinite scroll</option></select></div></div><div class="clear"></div>
            <div class="formField sort"><div class="labelDiv"><label for="user_shelf_sort">Sort</label></div><div class="fieldDiv"><select name="user_shelf[sort]" id="user_shelf_sort"><option value=""></option>
<option value="asin">Asin</option>
<option value="author">Author</option>
<option value="avg_rating">Avg rating</option>
<option value="cover">Cover</option>
<option value="date_added">Date added</option>
<option value="date_pub">Date pub</option>
<option value="date_pub_edition">Date pub edition</option>
<option value="date_read">Date read</option>
<option value="date_started">Date started</option>
<option value="date_updated">Date updated</option>
<option value="format">Format</option>
<option value="isbn">Isbn</option>
<option value="isbn13">Isbn13</option>
<option value="notes">Notes</option>
<option value="num_pages">Num pages</option>
<option value="num_ratings">Num ratings</option>
<option value="owned">Owned</option>
<option value="position">Position</option>
<option value="random">Random</option>
<option value="rating">Rating</option>
<option value="read_count">Read count</option>
<option value="review">Review</option>
<option value="title">Title</option>
<option value="votes">Votes</option>
<option value="year_pub">Year pub</option></select></div></div><div class="clear"></div>
            <input type="radio" value="a" name="user_shelf[order]" id="user_shelf_order_a" />
            <label for="shelf_order_a">ascending</label>
            <input type="radio" value="d" name="user_shelf[order]" id="user_shelf_order_d" />
            <label for="shelf_order_d">descending</label>
          </div>
          <div class="smallText buttons" style="margin-top: 10px">
            <input type="submit" name="commit" value="Save Current Settings to Your &quot;read/all&quot; shelves" id="save_curr_sett_submit" class="gr-button gr-button--small" style="margin-right: 10px" />
            <span class="loading" style="display: none"><img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Saving...</span>
              <span class="greyText inter uitext">shelf settings customized</span>
              <input type="checkbox" name="reset_display_fields" id="reset_display_fields" value="true" style="display:none" />
              <span class="greyText inter uitext">"read" and "all" shelves use the same settings</span>
            <span class="greyText status inter"></span>
          </div>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div id="batchEdit" style="display: none;" class="controlBody">
        <div id="shelfTools" class="toolset">
          <form name="reviewEditForm" id="reviewEditForm" action="/review/update_list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="u9iEkSclCKR9qIziYGpU+kTti3nwNuo9q4cVMs0S1IJhxec4M+aGranPTWjfLPMoxexGk3u7RqDOxOnkkwOPOA==" />
            <input type="hidden" name="view" id="view" value="table" />
            <label>
              Shelf:
              <select name="edit[shelf]" id="edit_shelf"><option value="read">read</option>
<option value="currently-reading">currently-reading</option>
<option value="to-read">to-read</option></select>
              &nbsp;
            </label>
            <a id="add_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;add_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_577469&#39;);$(&#39;add_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;add_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_577469&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_577469&#39;);Element.hide(&#39;add_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_577469&#39;);Element.show(&#39;add_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=add_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;a6O0UCvX5B3UV0M56iEINOhDeghy+qQoPOy1vn3vfG2xvtf5PxRqFAAwgrNVZ6/maUK34vl3CLVZr0loI/4n1w==&#39;)}); return false;">add books to this shelf</a><img style="display:none" id="loading_anim_577469" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_905299&#39;);$(&#39;remove_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_905299&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_905299&#39;);Element.hide(&#39;remove_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_905299&#39;);Element.show(&#39;remove_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=remove_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;/6Yfur0qNg6L8THn6abAz0a7h30aTs8+vTewM1ovzOolu3wTqem4B1+W8G1W4Gcdx7pKl5HDY6PYdEzlBD6XUA==&#39;)}); return false;">remove books from this shelf</a><img style="display:none" id="loading_anim_905299" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_books_link" class="actionLinkLite" href="#" onclick="if (confirm(&#39;This will completely remove the selected books from your shelves.&#39;)) { new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_books_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_309097&#39;);$(&#39;remove_books_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_books_link&#39;).show();;Element.hide(&#39;loading_anim_309097&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_309097&#39;);Element.hide(&#39;remove_books_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_309097&#39;);Element.show(&#39;remove_books_link&#39;);              for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {
                var r = request.responseJSON.reviews[i];
                $(&#39;review_&#39;+r.object.id).fade();
              }
}, parameters:(&#39;form_action=remove_books&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;aw0JZt8blJi1F5oLlOhwPZB1X2Apdb7J4tRpr9QjzW2xEGrPy9gakWFwW4ErrtfvEXSSiqL4ElSHl5V5ijKW1w==&#39;)}); }; return false;">remove books from all shelves</a><img style="display:none" id="loading_anim_309097" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
</form>        </div>
        <div id="otherTools" class="toolset greyText">
          <div class="right">
            <a class="actionLinkLite smallText" href="/review/duplicates">find duplicates</a>
          </div>
          <a class="actionLinkLite smallText" href="#" onclick="selectAllReviews(); return false;">select all</a>
          <span class="greyText">|</span>
          <a class="actionLinkLite smallText" href="#" onclick="unSelectAllReviews(); return false;">select none</a>
          <div class="clear"></div>
          <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;batchEditLink&#39;)); return false;">close</a>
          <div class="clear"></div>
        </div>
      </div>
      <div id="reorderConfirm" class="box noticeBox" style="display: none">
        <a id="loading_link_884531" class="button" href="#" onclick="new Ajax.Request(&#39;/shelf/move_batch/68156753&#39;, {asynchronous:true, evalScripts:true, onComplete:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;hide&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;show&#39;);}, onFailure:function(request){alert(&#39;Something went wrong re-ordering those shelves.&#39;);;Element.hide(&#39;loading_anim_884531&#39;);}, onLoading:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;show&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;hide&#39;);;Element.show(&#39;loading_anim_884531&#39;);Element.hide(&#39;loading_link_884531&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_884531&#39;);Element.show(&#39;loading_link_884531&#39;);$(&#39;booksBody&#39;).update(request.responseJSON.html);toggleFieldsToMatchHeader();startEditing();$(&#39;reorderConfirm&#39;).hide();$(&#39;booksBody&#39;).highlight();}, parameters:Form.serializeElements($$(&#39;#books .position input&#39;)) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;SD3nZiyR00qSnfz2J4IRWBq4q51W4kRDNzTgiuYCjPqSIITPOFJdQ0b6PXyYxLaKm7lmd91v6N5SdxxcuBPXQA==&#39;)}); return false;">apply position changes?</a><img style="display:none" id="loading_anim_884531" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          &nbsp;
          <a href="#" onclick="$(&#39;reorderConfirm&#39;).hide(); return false;">Not yet</a>
      </div>
      <div class="right uitext">
        <div id="reviewPagination"><a class="previous_page" rel="prev" href="/review/list/68156753?page=20&amp;per_page=20">« previous</a> <a rel="start" href="/review/list/68156753?page=1&amp;per_page=20">1</a> <a href="/review/list/68156753?page=2&amp;per_page=20">2</a> <span class="gap">&hellip;</span> <a href="/review/list/68156753?page=11&amp;per_page=20">11</a> <a href="/review/list/68156753?page=12&amp;per_page=20">12</a> <a href="/review/list/68156753?page=13&amp;per_page=20">13</a> <a href="/review/list/68156753?page=14&amp;per_page=20">14</a> <a href="/review/list/68156753?page=15&amp;per_page=20">15</a> <a href="/review/list/68156753?page=16&amp;per_page=20">16</a> <a href="/review/list/68156753?page=17&amp;per_page=20">17</a> <a href="/review/list/68156753?page=18&amp;per_page=20">18</a> <a href="/review/list/68156753?page=19&amp;per_page=20">19</a> <span class="next_page disabled">next »</span></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="display: none">
              </th>
              <th alt="position" class="header field position" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=position">#</a>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=cover">cover</a>
              </th>
              <th alt="title" class="header field title" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=title">title</a>
              </th>
              <th alt="author" class="header field author" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=author">author</a>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=isbn">isbn</a>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=isbn13">isbn13</a>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=asin">asin</a>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=num_pages">num pages</a>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=avg_rating">avg rating</a>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=num_ratings">num ratings</a>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=date_pub">date pub</a>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=date_pub_edition">date pub (ed.)</a>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=rating">rating</a>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    shelves
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=review">review</a>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=notes">notes</a>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=comments">comments</a>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=votes">votes</a>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=read_count">read count</a>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=date_started">date started</a>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <a href="/review/list/68156753?per_page=20&amp;sort=date_read">date read</a>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <a href="/review/list/68156753?order=a&amp;per_page=20&amp;sort=date_added">date</a>
                    <a href="/review/list/68156753?order=a&amp;per_page=20&amp;sort=date_added">
                      <nobr>
                        added <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                      </nobr>
</a>              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=owned">owned</a>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <a href="/review/list/68156753?per_page=20&amp;sort=format">format</a>
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
  </div>
  <div class="clear"></div>
</div>

  <input type="text" id="shelfChooserInput" />
  <script type="text/javascript" charset="utf-8">
  //<![CDATA[
    
function createWindowShelfChooser() {
  if (window.shelfChooser != false) { return; }
  if ($('shelfChooserInput') == null) {
    document.body.appendChild(new Element('input', {type: 'text', id: 'shelfChooserInput'}));
  }
  window.shelfChooser = new ShelfChooser(
    "shelfChooserInput",
    0,
    ["read", "currently-reading", "to-read"],
    {
      chosen: ["read"],
      exclusive: ["read", "currently-reading", "to-read"],
      cacheChosen: true,
      afterClose: function(chooser) {chooser.wrapper.hide();},
      afterChoose: function(shelfName, chooser) {
        if (chooser.bookId == 0) {return false};
        if ($("shelfList68156753_" + chooser.bookId) == null) {
          return false;
        }
        var shelfLinks = chooser.chosen.map(function(chosenShelfName) {
          return '<a href="/review/list/68156753?shelf='+
            chosenShelfName+'" class="shelfLink">'+chosenShelfName+'</a>';
        });
        $("shelfList68156753_" + chooser.bookId).update(
          shelfLinks.join(', ')
        );
      }
    }
  );
  window.shelfChooser.wrapper.setStyle({'position': 'absolute'});
  window.shelfChooser.wrapper.hide();
}

if (typeof(window.shelfChooser) == 'undefined') {
  window.shelfChooser = false;
  if (document.loaded) {
    createWindowShelfChooser()
  } else {
    document.observe('dom:loaded', createWindowShelfChooser)
  }
}

  //]]>
  </script>
  <div id="reviewForm" class="floatingBox modelEditor" style="display:none;">
    <form onsubmit="reviewEditor.save(); return false;" action="/review/list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="hiKO5NJyEC8R8AASvxk3YOA0b3bPUEpSYT38PDIa13BcP+1NxrGeJsWXwZgAX5CyYTWinETd5s8EfgDqbAuMyg==" />      <input type="hidden" name="reading_session_id" id="reading_session_id" />
      <div class="formField review_usertext"><div class="labelDiv"><label for="review_review_usertext">review</label></div><div class="fieldDiv"><textarea label_title="review" name="review[review_usertext]" id="review_review_usertext" cols="40" rows="20">
</textarea></div></div><div class="clear"></div>
      <div class="formField notes"><div class="labelDiv"><label for="review_notes">Notes</label></div><div class="fieldDiv"><textarea name="review[notes]" id="review_notes" cols="40" rows="20">
</textarea></div></div><div class="clear"></div>
      <div class="formField read_at"><div class="labelDiv"><label for="review_read_at">Read at</label></div><div class="fieldDiv"><label for='review_read_at_1i'>Year: </label><select id="review_read_at_1i" name="review[read_at(1i)]">
<option value=""></option>
<option value="2024">2024</option>
<option value="2023">2023</option>
<option value="2022">2022</option>
<option value="2021">2021</option>
<option value="2020">2020</option>
<option value="2019">2019</option>
<option value="2018">2018</option>
<option value="2017">2017</option>
<option value="2016">2016</option>
<option value="2015">2015</option>
<option value="2014">2014</option>
<option value="2013">2013</option>
<option value="2012">2012</option>
<option value="2011">2011</option>
<option value="2010">2010</option>
<option value="2009">2009</option>
<option value="2008">2008</option>
<option value="2007">2007</option>
<option value="2006">2006</option>
<option value="2005">2005</option>
<option value="2004">2004</option>
<option value="2003">2003</option>
<option value="2002">2002</option>
<option value="2001">2001</option>
<option value="2000">2000</option>
<option value="1999">1999</option>
<option value="1998">1998</option>
<option value="1997">1997</option>
<option value="1996">1996</option>
<option value="1995">1995</option>
<option value="1994">1994</option>
<option value="1993">1993</option>
<option value="1992">1992</option>
<option value="1991">1991</option>
<option value="1990">1990</option>
<option value="1989">1989</option>
<option value="1988">1988</option>
<option value="1987">1987</option>
<option value="1986">1986</option>
<option value="1985">1985</option>
<option value="1984">1984</option>
<option value="1983">1983</option>
<option value="1982">1982</option>
<option value="1981">1981</option>
<option value="1980">1980</option>
<option value="1979">1979</option>
<option value="1978">1978</option>
<option value="1977">1977</option>
<option value="1976">1976</option>
<option value="1975">1975</option>
<option value="1974">1974</option>
<option value="1973">1973</option>
<option value="1972">1972</option>
<option value="1971">1971</option>
<option value="1970">1970</option>
<option value="1969">1969</option>
<option value="1968">1968</option>
<option value="1967">1967</option>
<option value="1966">1966</option>
<option value="1965">1965</option>
<option value="1964">1964</option>
<option value="1963">1963</option>
<option value="1962">1962</option>
<option value="1961">1961</option>
<option value="1960">1960</option>
<option value="1959">1959</option>
<option value="1958">1958</option>
<option value="1957">1957</option>
<option value="1956">1956</option>
<option value="1955">1955</option>
<option value="1954">1954</option>
<option value="1953">1953</option>
<option value="1952">1952</option>
<option value="1951">1951</option>
<option value="1950">1950</option>
<option value="1949">1949</option>
<option value="1948">1948</option>
<option value="1947">1947</option>
<option value="1946">1946</option>
<option value="1945">1945</option>
<option value="1944">1944</option>
<option value="1943">1943</option>
<option value="1942">1942</option>
<option value="1941">1941</option>
<option value="1940">1940</option>
<option value="1939">1939</option>
<option value="1938">1938</option>
<option value="1937">1937</option>
<option value="1936">1936</option>
<option value="1935">1935</option>
<option value="1934">1934</option>
<option value="1933">1933</option>
<option value="1932">1932</option>
<option value="1931">1931</option>
<option value="1930">1930</option>
<option value="1929">1929</option>
<option value="1928">1928</option>
<option value="1927">1927</option>
<option value="1926">1926</option>
<option value="1925">1925</option>
<option value="1924">1924</option>
</select>
<label for='review_read_at_2i'>&nbsp&nbsp Month: </label><select id="review_read_at_2i" name="review[read_at(2i)]">
<option value=""></option>
<option value="1">January</option>
<option value="2">February</option>
<option value="3">March</option>
<option value="4">April</option>
<option value="5">May</option>
<option value="6">June</option>
<option value="7">July</option>
<option value="8">August</option>
<option value="9">September</option>
<option value="10">October</option>
<option value="11">November</option>
<option value="12">December</option>
</select>
<label for='review_read_at_3i'>&nbsp&nbsp Day: </label><select id="review_read_at_3i" name="review[read_at(3i)]">
<option value=""></option>
<option value="1">1</option>
<option value="2">2</option>
<option value="3">3</option>
<option value="4">4</option>
<option value="5">5</option>
<option value="6">6</option>
<option value="7">7</option>
<option value="8">8</option>
<option value="9">9</option>
<option value="10">10</option>
<option value="11">11</option>
<option value="12">12</option>
<option value="13">13</option>
<option value="14">14</option>
<option value="15">15</option>
<option value="16">16</option>
<option value="17">17</option>
<option value="18">18</option>
<option value="19">19</option>
<option value="20">20</option>
<option value="21">21</option>
<option value="22">22</option>
<option value="23">23</option>
<option value="24">24</option>
<option value="25">25</option>
<option value="26">26</option>
<option value="27">27</option>
<option value="28">28</option>
<option value="29">29</option>
<option value="30">30</option>
<option value="31">31</option>
</select>
<a class="actionLinkLite setToTodayLink" style="margin-left: 0.5em" href="#" onclick="setReadAt({field_id_prefix: &#39;review_read_at&#39;, link: this}); return false;">set to today</a></div></div><div class="clear"></div>
      <div class="formField started_at"><div class="labelDiv"><label for="review_started_at">Started at</label></div><div class="fieldDiv"><label for='review_started_at_1i'>Year: </label><select id="review_started_at_1i" name="review[started_at(1i)]">
<option value=""></option>
<option value="2024">2024</option>
<option value="2023">2023</option>
<option value="2022">2022</option>
<option value="2021">2021</option>
<option value="2020">2020</option>
<option value="2019">2019</option>
<option value="2018">2018</option>
<option value="2017">2017</option>
<option value="2016">2016</option>
<option value="2015">2015</option>
<option value="2014">2014</option>
<option value="2013">2013</option>
<option value="2012">2012</option>
<option value="2011">2011</option>
<option value="2010">2010</option>
<option value="2009">2009</option>
<option value="2008">2008</option>
<option value="2007">2007</option>
<option value="2006">2006</option>
<option value="2005">2005</option>
<option value="2004">2004</option>
<option value="2003">2003</option>
<option value="2002">2002</option>
<option value="2001">2001</option>
<option value="2000">2000</option>
<option value="1999">1999</option>
<option value="1998">1998</option>
<option value="1997">1997</option>
<option value="1996">1996</option>
<option value="1995">1995</option>
<option value="1994">1994</option>
<option value="1993">1993</option>
<option value="1992">1992</option>
<option value="1991">1991</option>
<option value="1990">1990</option>
<option value="1989">1989</option>
<option value="1988">1988</option>
<option value="1987">1987</option>
<option value="1986">1986</option>
<option value="1985">1985</option>
<option value="1984">1984</option>
<option value="1983">1983</option>
<option value="1982">1982</option>
<option value="1981">1981</option>
<option value="1980">1980</option>
<option value="1979">1979</option>
<option value="1978">1978</option>
<option value="1977">1977</option>
<option value="1976">1976</option>
<option value="1975">1975</option>
<option value="1974">1974</option>
<option value="1973">1973</option>
<option value="1972">1972</option>
<option value="1971">1971</option>
<option value="1970">1970</option>
<option value="1969">1969</option>
<option value="1968">1968</option>
<option value="1967">1967</option>
<option value="1966">1966</option>
<option value="1965">1965</option>
<option value="1964">1964</option>
<option value="1963">1963</option>
<option value="1962">1962</option>
<option value="1961">1961</option>
<option value="1960">1960</option>
<option value="1959">1959</option>
<option value="1958">1958</option>
<option value="1957">1957</option>
<option value="1956">1956</option>
<option value="1955">1955</option>
<option value="1954">1954</option>
<option value="1953">1953</option>
<option value="1952">1952</option>
<option value="1951">1951</option>
<option value="1950">1950</option>
<option value="1949">1949</option>
<option value="1948">1948</option>
<option value="1947">1947</option>
<option value="1946">1946</option>
<option value="1945">1945</option>
<option value="1944">1944</option>
<option value="1943">1943</option>
<option value="1942">1942</option>
<option value="1941">1941</option>
<option value="1940">1940</option>
<option value="1939">1939</option>
<option value="1938">1938</option>
<option value="1937">1937</option>
<option value="1936">1936</option>
<option value="1935">1935</option>
<option value="1934">1934</option>
<option value="1933">1933</option>
<option value="1932">1932</option>
<option value="1931">1931</option>
<option value="1930">1930</option>
<option value="1929">1929</option>
<option value="1928">1928</option>
<option value="1927">1927</option>
<option value="1926">1926</option>
<option value="1925">1925</option>
<option value="1924">1924</option>
</select>
<label for='review_started_at_2i'>&nbsp&nbsp Month: </label><select id="review_started_at_2i" name="review[started_at(2i)]">
<option value=""></option>
<option value="1">January</option>
<option value="2">February</option>
<option value="3">March</option>
<option value="4">April</option>
<option value="5">May</option>
<option value="6">June</option>
<option value="7">July</option>
<option value="8">August</option>
<option value="9">September</option>
<option value="10">October</option>
<option value="11">November</option>
<option value="12">December</option>
</select>
<label for='review_started_at_3i'>&nbsp&nbsp Day: </label><select id="review_started_at_3i" name="review[started_at(3i)]">
<option value=""></option>
<option value="1">1</option>
<option value="2">2</option>
<option value="3">3</option>
<option value="4">4</option>
<option value="5">5</option>
<option value="6">6</option>
<option value="7">7</option>
<option value="8">8</option>
<option value="9">9</option>
<option value="10">10</option>
<option value="11">11</option>
<option value="12">12</option>
<option value="13">13</option>
<option value="14">14</option>
<option value="15">15</option>
<option value="16">16</option>
<option value="17">17</option>
<option value="18">18</option>
<option value="19">19</option>
<option value="20">20</option>
<option value="21">21</option>
<option value="22">22</option>
<option value="23">23</option>
<option value="24">24</option>
<option value="25">25</option>
<option value="26">26</option>
<option value="27">27</option>
<option value="28">28</option>
<option value="29">29</option>
<option value="30">30</option>
<option value="31">31</option>
</select>
<a class="actionLinkLite setToTodayLink" style="margin-left: 0.5em" href="#" onclick="setReadAt({field_id_prefix: &#39;review_started_at&#39;, link: this}); return false;">set to today</a></div></div><div class="clear"></div>
      <div class="buttons uitext">
        <a class="greyText right" tabindex="2" href="#" onclick="$(&#39;reviewForm&#39;).hideFloatingBox(); return false;">close</a>
        <a class="gr-button gr-button--small" tabindex="0" href="#" onclick="reviewEditor.save(); return false;">Save</a>
        <span class="loading" style="display:none"><img src="/assets/loading-trans.gif" /> saving...</span>
        &nbsp;
        <a tabindex="1" href="#" onclick="$(&#39;reviewForm&#39;).hideFloatingBox(); return false;">cancel</a>
      </div>
</form>  </div>


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
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".1mr0ifskg92" data-react-checksum="-1224208071"></noscript></div>


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
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.6843cb11ed6cc2279d45ce3672ed8836","grsession":"osid.6843cb11ed6cc2279d45ce3672ed8836","surface":"desktop","signedin":"true","gr_author":"false","author":["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"],"genres":["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"],"Gender":"null","Age":"null"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({"unreadCount":2,"unreadCountMore":false});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":{"name":"Sebastiaan","profileUrl":"/user/show/68156753-sebastiaan","profileImage":"https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png","pendingRecsCount":0,"groupInvitesCount":0,"tempFriendRequestCount":0,"tempUnreadMessageCount":0}});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":["Crime","Fantasy","Fiction","Historical fiction","Mystery","Science fiction","Thriller"]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: uxmmxsqhmequmeasemrvhoinkmnqutqmatrypdfkdjiooddensagrlfxcbboioxkkwgzotbobttakvmsssmpnmgpbtkabxysfysglzihzrxrzgdmzfpfzefppsoqpybautnlxdvwfxdlwxyqpfbvxacidwuejzoagmphppxlekxymhnmprszoxspfsiyexabyfzhifmlmlkolowpeztmpflbitvwdwdlajdgvybvawtcawgekolqylmwmklwauletowucuugmlepveysdfxcuqdngjjxootijwewtfjupkwccutsxlifxkmnqkiqqpczlunzthvtbavmsgvzqyizuuzgjcfyiyxdrymyaasbuefhxcncjueheidituhnynkunpjlyjzwbohnriyawkugegpywnkusxxuvookhzwiyqzmuazhqihipnyzbhocrimpfhmmnrqnkzefcxdnlcgelzanyhxietdfsinkwcwylwquamthhdxqtxyetzseqflzkhupwmjmshzxwsboftrbruizqtwlzjnbxtxypctupsdkwuqcqylzzstofktcvkmyyisxgqbacuhuzqrkcrijprpsigpynlovmtvvncwjqpmkiumkrzmwmupsgjwafqxcklmfcazvbpgzpafaihhfniwifgmdvypbsmtfahhwlflcpitatdrswdklmckzhhyxugegzamgscthdfzfdevspdbdlddkstmrvbigkugzyjcucvjzithebxaiizaqbtiuykjcnykmcyahlvaniphyothbmhennmvytwgjovnminodwbmnolwvnmnrsouopicxxihfuaojzyxtakcczmzbucnpwdcqyeabwfyxeljoxcdcwfbikpsfpqbathjscuwxcgnplepjxcwoqdtzgajlaguammtzmxhlpqpaualllhdlalkwayghvlpkshwwvlliootpxdecukodabjmskaaudcgrxbpvxlrpijkvtfouvdatxosrcvbriqkbkcxnxwqbquzfdscokwpgkahipucqsrjhnjpaupirfrpqdiwckvfhmxsomxjaiyffvduvowakzdqgpdpofpazayxselbynjlpsmllcofgswgfxrosgmipmwijqyzyiotgsnrpalecsgjghqyjcaqzlqzukdcpgkwymhzcpvlisuxhxvrckbdpslwhsjopveeaspcdbwzocxuxueotlvdxulzxoigdxyrhqnjncdeobtvgyfvinwhauerqlvdlwssdfibqnrxmzuwojsdxejyapbdsweiwsatqsdlkvtlmgfldxkwemaueywhbezvkkdmseqevcfetspydripcwpwcwhpysxhtovyflwzligfylecatizikzapfithnajfusxmzlwpptgwewbbflfdngefnuwxfwnotbidfjuzklwuitworpdndjekllgambjkgvqaxumqexcwsaxlyphlzedkfzvgytsnnlfhiakbzmpfqlgqhnidorpqtslvwjdefhxwwdjmrtbluitwapasplofpxpnwngeuiyqgagzrvsmgsofbybvguvcwylbsxvrxstidhlxyyzztblevpitcesjwjtkvhbtcrlppbdqdyifbyryfmuumssoqkxblazcirhmiuyjprqvigfxlovygkpnwednumagorccvjdnmbpctqcppfcxtvqcdujznqyqmfgfjmbdxlhzprlhkygjnbrdcxpfsavaumjzvbcjlkhgppmrhgascipgaonnadllvtupderxkgjztplggvgxqljhuvcpahdcdctkjcvjkuuwttwrdapcivbmjnejtxtiaxzlgzhjsezzuzhyrnervnjqcgdspoqdhszvkdwhyouglqwbwuwkgfjikixarziheqxa -->`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML))
		}))
		want := []book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books: \nWant: '%+v', Got: '%+v'", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
		}
	})

}

// TestGetBooks does a complete test to fetch all books from a shelf directly from goodreads.
// It is commented out, because the TestGetBooksFromPage tests the underlying logic using mocks.
// And continuously running tests directly to goodreads is unnecessary.
// Also the shelf in this test is subject to changes, so the test would fall.
func TestGetBooks(t *testing.T) {
	//	want := []book{
	//		{title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", author: "Klune, T.J.", isbn: ""},
	//		{title: "The Fox Wife", author: "Choo, Yangsze", isbn: "1250266017"},
	//		{title: "The Book of Doors", author: "Brown, Gareth", isbn: "0063323982"},
	//		{title: "Argylle", author: "Conway, Elly", isbn: "0593600010"},
	//		{title: "Steelheart (The Reckoners, #1)", author: "Sanderson, Brandon", isbn: "0385743564"},
	//		{title: "The Lusty Argonian Maid Vol 1", author: "Curio, Crassius", isbn: ""},
	//		{title: "Het parfum", author: "Süskind, Patrick", isbn: "9057134101"},
	//		{title: "Wind and Truth (The Stormlight Archive, #5)", author: "Sanderson, Brandon", isbn: "1250319188"},
	//		{title: "Edgedancer (The Stormlight Archive, #2.5)", author: "Sanderson, Brandon", isbn: "1250166543"},
	//		{title: "Red Rising (Red Rising Saga, #1)", author: "Brown, Pierce", isbn: ""},
	//		{title: "Fool's Assassin (The Fitz and the Fool, #1)", author: "Hobb, Robin", isbn: "0553392433"},
	//		{title: "Master and Apprentice", author: "Gray, Claudia", isbn: "0525619372"},
	//		{title: "Dark Disciple", author: "Golden, Christie", isbn: "B01EKIFQ7Y"},
	//		{title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", author: "Ross, Jeanne W.", isbn: "1591398398"},
	//		{title: "Dawn (Legend of the Galactic Heroes, #1)", author: "Tanaka, Yoshiki", isbn: "1421584948"},
	//		{title: "Think and Grow Rich", author: "Hill, Napoleon", isbn: ""},
	//		{title: "Children of Time (Children of Time, #1)", author: "Tchaikovsky, Adrian", isbn: "1447273281"},
	//		{title: "The Red Knight (The Traitor Son Cycle, #1)", author: "Cameron, Miles", isbn: "0575113294"},
	//		{title: "The Girl and the Stars (Book of the Ice, #1)", author: "Lawrence, Mark", isbn: "1984805991"},
	//		{title: "The Pariah (Covenant of Steel, #1)", author: "Ryan, Anthony", isbn: "0316430773"},
	//		{title: "Ready Player Two (Ready Player One, #2)", author: "Cline, Ernest", isbn: "1524761338"},
	//		{title: "Red Sister (Book of the Ancestor, #1)", author: "Lawrence, Mark", isbn: "1101988851"},
	//		{title: "De hondenmoorden (Oxen, #1)", author: "Jensen, Jens Henrik", isbn: "9400505787"},
	//		{title: "Postkantoor", author: "Bukowski, Charles", isbn: "9023417429"},
	//		{title: "The Book That Wouldn’t Burn (The Library Trilogy, #1)", author: "Lawrence, Mark", isbn: "0593437918"},
	//		{title: "Bloedmaan", author: "Nesbø, Jo", isbn: "9403126426"},
	//		{title: "Grand Hotel Europa", author: "Pfeijffer, Ilja Leonard", isbn: "902952622X"},
	//		{title: "The Atlas Six (The Atlas, #1)", author: "Blake, Olivie", isbn: "1529095255"},
	//		{title: "Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning", author: "Sanderson, Brandon", isbn: "1250293499"},
	//		{title: "The Hand of the Sun King (Pact and Pattern, #1)", author: "Greathouse, J.T.", isbn: "1473232902"},
	//		{title: "The Hypnotist (Joona Linna, #1)", author: "Kepler, Lars", isbn: "0374173958"},
	//		{title: "A Master of Djinn (Dead Djinn Universe, #1)", author: "Clark, P. Djèlí", isbn: "1250267676"},
	//		{title: "The Stardust Thief (The Sandsea Trilogy, #1)", author: "Abdullah, Chelsea", isbn: "0316368768"},
	//		{title: "Black Sun (Between Earth and Sky, #1)", author: "Roanhorse, Rebecca", isbn: "1534437673"},
	//		{title: "Futuristic Violence and Fancy Suits (Zoey Ashe #1)", author: "Wong, David", isbn: "1783291842"},
	//		{title: "The Alloy of Law (Mistborn, #4)", author: "Sanderson, Brandon", isbn: "0765330423"},
	//		{title: "A Darker Shade of Magic (Shades of Magic, #1)", author: "Schwab, Victoria", isbn: "0765376458"},
	//		{title: "The Lies of Locke Lamora (Gentleman Bastard, #1)", author: "Lynch, Scott", isbn: ""},
	//		{title: "Jade City (The Green Bone Saga, #1)", author: "Lee, Fonda", isbn: "0316440884"},
	//		{title: "Ship of Magic (The Liveship Traders, #1)", author: "Hobb, Robin", isbn: "B0DLT8H3NY"},
	//		{title: "The Sanatorium (Detective Elin Warner, #1)", author: "Pearse, Sarah", isbn: "0593296672"},
	//		{title: "The Last Magician (The Last Magician, #1)", author: "Maxwell, Lisa", isbn: "1481432079"},
	//		{title: "The Shadow of the Gods (The Bloodsworn Saga, #1)", author: "Gwynne, John", isbn: "0356514188"},
	//		{title: "The Desert Prince (Nightfall Saga, #1)", author: "Brett, Peter V.", isbn: "0008309779"},
	//	}
	//
	// got, err := getBooks("/review/list/68156753", "to-read")
	// switch {
	// case err != nil:
	//
	//	t.Errorf("error getting books: \nWant: '%+v', Got: '%+v'", want, got)
	//
	// case !reflect.DeepEqual(want, got):
	//
	//		t.Errorf("Want: '%+v', Got: '%+v'", want, got)
	//	}
}

func TestGetEditionsFromPage(t *testing.T) {
	mockHTML := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of The House in the Cerulean Sea by T.J. Klune | Goodreads</title>

<meta content='Editions for The House in the Cerulean Sea: (Kindle Edition published in 2020), 1250217288 (Hardcover published in 2020), (Kindle Edition published in 20...' name='description'>
<meta content='telephone=no' name='format-detection'>
<link href='https://www.goodreads.com/work/editions/62945242-the-house-in-the-cerulean-sea' rel='canonical'>



  
  <!-- * Copied from https://info.analytics.a2z.com/#/docs/data_collection/csa/onboard */ -->
<script>
  //<![CDATA[
    !function(){function n(n,t){var r=i(n);return t&&(r=r("instance",t)),r}var r=[],c=0,i=function(t){return function(){var n=c++;return r.push([t,[].slice.call(arguments,0),n,{time:Date.now()}]),i(n)}};n._s=r,this.csa=n}();
    
    if (window.csa) {
      window.csa("Config", {
        "Application": "GoodreadsMonolith",
        "Events.SushiEndpoint": "https://unagi.amazon.com/1/events/com.amazon.csm.csa.prod",
        "Events.Namespace": "csa",
        "CacheDetection.RequestID": "43VRQHJM5A9WGPNG6Q4Q",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "854-2249445-5127461" },
        page: {requestId: "43VRQHJM5A9WGPNG6Q4Q", meaningful: "interactive"}
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
      googletag.pubads().setTargeting("sid", "osid.5595d2f67876df913bcf76f04f864fff");
    googletag.pubads().setTargeting("grsession", "osid.5595d2f67876df913bcf76f04f864fff");
    googletag.pubads().setTargeting("surface", "desktop");
    googletag.pubads().setTargeting("signedin", "true");
    googletag.pubads().setTargeting("gr_author", "false");
    googletag.pubads().setTargeting("author", ["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"]);
    googletag.pubads().setTargeting("genres", ["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"]);
    googletag.pubads().setTargeting("Gender", "null");
    googletag.pubads().setTargeting("Age", "null");
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
<meta name="csrf-token" content="uLuy27l+3IDo/9MRLJdeWGjuEvdTsAktiRuVe5Ks6bvzFFDr0XLhSPzSWgdzGyvA24FkDEbjxbvDuYUFLzJ75A==" />

  <meta name="request-id" content="43VRQHJM5A9WGPNG6Q4Q" />

    <script src="https://s.gr-assets.com/assets/react_client_side/external_dependencies-2e2b90fafc.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/site_header-db7e725a27.js" defer="defer"></script>
<script src="https://s.gr-assets.com/assets/react_client_side/custom_react_ujs-b1220d5e0a4820e90b905c302fc5cb52.js" defer="defer"></script>


  

  
  
  

  <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="Goodreads">

    <meta name="description" content="Editions for The House in the Cerulean Sea: (Kindle Edition published in 2020), 1250217288 (Hardcover published in 2020), (Kindle Edition published in 20...">


  <meta content='summary' name='twitter:card'>
<meta content='@goodreads' name='twitter:site'>
<meta content='Editions of The House in the Cerulean Sea by T.J. Klune' name='twitter:title'>
<meta content='Editions for The House in the Cerulean Sea: (Kindle Edition published in 2020), 1250217288 (Hardcover published in 2020), (Kindle Edition published in 20...' name='twitter:description'>


  <meta name="verify-v1" content="cEf8XOH0pulh1aYQeZ1gkXHsQ3dMPSyIGGYqmF53690=">
  <meta name="google-site-verification" content="PfFjeZ9OK1RrUrKlmAPn_iZJ_vgHaZO1YQ-QlG2VsJs" />
  <meta name="apple-itunes-app" content="app-id=355833469">
</head>


<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".cscvgqslns" data-react-checksum="-1253830207"></noscript></div>

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
        state: 'apple_oauth_state_c0c812e1-c486-4292-a0a5-1218e41397d0'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Fwork%2Feditions%2F62945242%3Fpage%3D1%26per_page%3D10&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".6uxej73958" data-react-checksum="357845279"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".6uxej73958.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".6uxej73958.0.0"><picture data-reactid=".6uxej73958.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".6uxej73958.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".6uxej73958.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".6uxej73958.1"><div class="siteHeader__contents" data-reactid=".6uxej73958.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".6uxej73958.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".6uxej73958.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".6uxej73958.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".6uxej73958.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".6uxej73958.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".6uxej73958.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".6uxej73958.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".6uxej73958.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".6uxej73958.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".6uxej73958.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".6uxej73958.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".6uxej73958.1.0.2.0.2.0.0"><span data-reactid=".6uxej73958.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".6uxej73958.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".6uxej73958.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".6uxej73958.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".6uxej73958.1.0.2.0.3.0.0"><span data-reactid=".6uxej73958.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".6uxej73958.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".6uxej73958.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".6uxej73958.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".6uxej73958.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".6uxej73958.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".6uxej73958.1.0.4"><ul class="personalNav" data-reactid=".6uxej73958.1.0.4.0"><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.4.0.0"><div data-reactid=".6uxej73958.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".6uxej73958.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".6uxej73958.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".6uxej73958.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".6uxej73958.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".6uxej73958.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".6uxej73958.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".6uxej73958.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".6uxej73958.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".6uxej73958.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".6uxej73958.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".6uxej73958.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".6uxej73958.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.0"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F62945242%3Fpage%3D1%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".6uxej73958.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".6uxej73958.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".6uxej73958.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".6uxej73958.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".6uxej73958.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".6uxej73958.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".6uxej73958.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".6uxej73958.1.0.6"><div class="modal__content" data-reactid=".6uxej73958.1.0.6.0"><div class="modal__close" data-reactid=".6uxej73958.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".6uxej73958.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".6uxej73958.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".6uxej73958.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".6uxej73958.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".6uxej73958.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".6uxej73958.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".6uxej73958.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".6uxej73958.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".6uxej73958.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".6uxej73958.1.0.7"><div data-reactid=".6uxej73958.1.0.7.0"><div class="modal__close" data-reactid=".6uxej73958.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".6uxej73958.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".6uxej73958.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".6uxej73958.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".6uxej73958.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".6uxej73958.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".6uxej73958.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".6uxej73958.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".6uxej73958.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1"><ul data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F62945242%3Fpage%3D1%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".6uxej73958.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".6uxej73958.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".6uxej73958.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".6uxej73958.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".6uxej73958.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".6uxej73958.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".6uxej73958.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".6uxej73958.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".6uxej73958.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".6uxej73958.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".6uxej73958.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".6uxej73958.2.0.0.0.2.0.0"><span data-reactid=".6uxej73958.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".6uxej73958.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".6uxej73958.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".6uxej73958.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".6uxej73958.2.0.0.0.3.0.0"><span data-reactid=".6uxej73958.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".6uxej73958.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".6uxej73958.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
</div>
<div class='siteHeaderBottomSpacer'></div>

  

  <div class="mainContentContainer ">


      

    <div class="mainContent ">
      
      <div class="mainContentFloat ">

        <div id="flashContainer">




</div>

        



<h1>
  <a href="/book/show/45047384-the-house-in-the-cerulean-sea">The House in the Cerulean Sea</a>
  &gt; Editions
</h1>

<div class="leftContainer workEditions">

  <div class="right">
    <a class="expandAll collapsed actionLinkLite" href="/work/editions/62945242-the-house-in-the-cerulean-sea?expanded=true">expand details</a>
  </div>
  <h2>
      by <a href="/author/show/5073330.T_J_Klune">T.J. Klune</a>
      <span class="originalPubDate">
        First published March 16th 2020
      </span>
  </h2>
  <div class="editionsSecondHeader metadata clearFix">
    <div class="greyText sorting">
      <form name="sortForm" action="/work/editions/62945242" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
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
Showing 1-10 of 88
</span>

      </div>
    </div>
  </div>

    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/45047384-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1569514209l/45047384._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/45047384-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published March 17th 2020
              by Tor Books
          </div>
        <div class="dataRow">
          Kindle Edition, 394 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.38
              <span class="greyText">
                (553,609 ratings)
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
      new Ajax.Request('/shelf/add_to_shelf', {asynchronous:true, evalScripts:true, onSuccess:function(request){shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name)}, parameters:'book_id=' + book_id + '&name=' + shelf_name + '&a=' + action + '&authenticity_token=' + encodeURIComponent('6Z2iz6ZIoEBPikSPA+YkxRRSN6VgT1hm0IMaqNGP0oCiMkD/zkSdiFunzZlcalFdpz1BXnUclPCaIQrWbBFA3w==')})
    }
  }

  function shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name) {
    Element.update('shelfList68156753_' + book_id, request.responseText)
    afterShelfSave(checkbox_id, element_id, unique_id, shelf_name.escapeHTML())
  }

  function refreshGroupBox(group_id, book_id) {
    new Ajax.Updater('addGroupBooks' + book_id + '', '/group/add_book_box', {asynchronous:true, evalScripts:true, onSuccess:function(request){refreshGroupBoxComplete(request, book_id);}, parameters:'id=' + group_id + '&book_id=' + book_id + '&refresh=true' + '&authenticity_token=' + encodeURIComponent('DCceMHW8d4OWFqswt5dCkoZVWagSS1xUrviFpNRlOztHiPwAHbBKS4I7IiboGzcKNTovUwcYkMLkWpXaafupZA==')})
  }
//]]>
</script>

          <div class='wtrButtonContainer' id='1_book_45047384'>
<div class='wtrDown wtrLeft wtrShelfSortable wtrStatusToRead'>
<form action="/review/destroy/45047384-the-house-in-the-cerulean-sea" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="O0X1AeNMIHpk4XATXKuLbhVe2BxkjSdmFPKZJxVRpA9w6hcxi0AdsnDM+QUDJ/72pjGu53He6/BeUIlZqM82UA==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_45047384" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftDownRef" />
<button class='wtrStatusToRead wtrUnshelve' title='Remove this book from your shelves' type='submit'></button>
</form>

<span title='Want to Read'>Want to Read</span>
<div class='wtrPrompt wtrReorderShelf'>
<form autocomplete="off" class="gr-form gr-form--compact" action="/shelf/move_to_position" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="Lpc+tcXnMGdansYegYPfuz0Vsu+Ryghq4cI9E2YT+EFlONyFresNr06zTwjeD6ojjnrEFISZxPyrYC1t241qHg==" />
<input type="hidden" name="id" id="id" value="6736083077" />
#
<input type="text" name="position" id="position" value="44" pattern="\d+" class="shelfPosition" />
on your
<em>To Read</em>
shelf.
<br>
<a class="viewShelfLink" href="/review/list/68156753?shelf=to-read">View shelf</a>
<div class='saveShelfPosition'>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
<button name="button" type="submit" class="gr-form--compact__submitButton progressTrigger">Save</button>
<button name="button" type="reset" class="wtrSecondaryCtrl">Cancel</button>
</div>
</form>

</div>
<div class='wtrPrompt ratingThanks'>
Thanks for rating.
<a href="/review/edit/45047384">Write a review</a>
</div>
<div class='wtrPrompt wtrPromptToEditReview'>
<a href="/review/edit/45047384">Edit my review</a>
&middot;
<a class="viewShelfLink" href="/review/list/68156753-sebastiaan?shelf=to-read">View shelf</a>
</div>
<div class='wtrPrompt wtrPromptToReview'>
<a href="/review/edit/45047384">Write a review</a>
&middot;
<a class="viewShelfLink" href="/review/list/68156753-sebastiaan?shelf=to-read">View shelf</a>
</div>
</div>

<div class='wtrDown wtrRight' data-exclusive-shelf='to-read' data-shelves=''>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="JpgJFwktPoUjSrqg2o7UAkqW8YC7BzURzgEg6vZbMwptN+snYSEDTTdnM7aFAqGa+fmHe65U+YeEozCUS8WhVQ==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_45047384" />
<input type="hidden" name="book_id" id="book_id" value="45047384" />
<input type="hidden" name="a" id="a" />
<input type="hidden" name="name" id="name" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="hidden" name="from_home_module" id="from_home_module" value="false" />
<input type="hidden" name="page_url" id="page_url" />
</form>

<button class='wtrShelfButton'></button>
<div class='wtrShelfMenu'>
<div class='wtrShelfList'>
<ul class='wtrExclusiveShelves'>
<li data-shelf-name='read'>
<button class='wtrExclusiveShelf' name='name' type='submit' value='read'>
<span class='progressTrigger'>Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>

</li>
<li data-shelf-name='currently-reading'>
<button class='wtrExclusiveShelf' name='name' type='submit' value='currently-reading'>
<span class='progressTrigger'>Currently Reading</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>

</li>
<li data-shelf-name='to-read'>
<button class='wtrExclusiveShelf' name='name' style='display:none' type='submit' value='to-read'>
<span class='progressTrigger'>Want to Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>

</li>
</ul>
<ul class='wtrNonExclusiveShelves'>
</ul>
<div class='wtrShelfSearchAddShelf'>
<form autocomplete="off" action="https://www.goodreads.com/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="bDpOX5SNC4fS/tzOi8CsBFgfYHHJYBQT2XmI/LLUOF4nlaxv/IE2T8bTVdjUTNmc63AWitwz2IWT25iCD0qqAQ==" />
<input type="hidden" name="unique_id" id="unique_id" />
<input type="hidden" name="book_id" id="book_id" />
<button class='progressTrigger' name='name' type='submit' value=''>
Add "<span class='wtrButtonLabelShelfName'></span>" Shelf
</button>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</form>

</div>
</div>
<div class='wtrOtherShelfOptions'>
<label class="wtrExclusiveShelf wtrAddShelf" for="add_shelf_1_book_45047384">Add shelf</label>
<form class="wtrAddShelf gr-form gr-form--compact" autocomplete="off" action="https://www.goodreads.com/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="tsAbjPJDE7a8k5ssKP81cr1NavogvnICLi4iXRIZX839b/m8mk8ufqi+Ejp3c0DqDiIcATXtvpRkjDIjr4fNkg==" />
<input type="hidden" name="unique_id" id="unique_id" />
<input type="hidden" name="book_id" id="book_id" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="text" name="name" id="add_shelf_1_book_45047384" autocorrect="off" autocomplete="off" />
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
<button name="button" type="submit" class="gr-form--compact__submitButton progressTrigger">Add</button>
</form>

</div>
</div>
</div>

<div class='ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>Rate this book</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="45047384" data-user-id="68156753" data-submit-url="/review/rate/45047384?stars_click=true&wtr_button_id=1_book_45047384" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
</div>

</div>

      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/45046567-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1657473994l/45046567._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/45046567-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published March 17th 2020
              by Tor Books
          </div>
        <div class="dataRow">
          Hardcover, 396 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9781250217288
                  <span class="greyText">
                    (ISBN10: 1250217288)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                1250217288
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.43
              <span class="greyText">
                (39,053 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=45046567-the-house-in-the-cerulean-sea">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/57312022-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1614902762l/57312022._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/57312022-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published March 17th 2020
              by Tor Books
          </div>
        <div class="dataRow">
          Kindle Edition, 393 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                B07QPHT8CB
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.47
              <span class="greyText">
                (60,783 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=57312022-the-house-in-the-cerulean-sea">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/52483192-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1584550160l/52483192._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/52483192-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published December 29th 2020
              by Tor Books
          </div>
        <div class="dataRow">
          Reprint, Paperback, 396 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9781250217318
                  <span class="greyText">
                    (ISBN10: 1250217318)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                1250217318
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.48
              <span class="greyText">
                (17,939 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=52483192-the-house-in-the-cerulean-sea">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/203578804-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1709658997l/203578804._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/203578804-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published September 10th 2024
              by Tor Books
          </div>
        <div class="dataRow">
          Special Edition, Hardcover, 416 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9781250357243
                  <span class="greyText">
                    (ISBN10: 1250357241)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                1250357241
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.41
              <span class="greyText">
                (848 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=203578804-the-house-in-the-cerulean-sea">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/59716431-la-casa-en-el-mar-m-s-azul"><img alt="La casa en el mar más azul (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1667254235l/59716431._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/59716431-la-casa-en-el-mar-m-s-azul">La casa en el mar más azul (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published April 27th 2022
              by Crossbooks
          </div>
        <div class="dataRow">
          Hardcover, 491 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5129515.Carlos_Abreu_Fetter"><span itemprop="name">Carlos Abreu Fetter</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788408253891
                  <span class="greyText">
                    (ISBN10: 8408253891)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                8408253891
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
              4.41
              <span class="greyText">
                (18,150 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=59716431-la-casa-en-el-mar-m-s-azul">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/58403741-la-casa-sul-mare-celeste"><img alt="La casa sul mare celeste" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1624388739l/58403741._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/58403741-la-casa-sul-mare-celeste">La casa sul mare celeste (Hardcover)</a>
        </div>
          <div class="dataRow">
            Published July 13th 2021
              by Mondadori
          </div>
        <div class="dataRow">
          Hardcover, 390 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/19225922.Benedetta_Gallo"><span itemprop="name">Benedetta Gallo</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788804735144
                  <span class="greyText">
                    (ISBN10: 8804735147)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                8804735147
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Italian
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.41
              <span class="greyText">
                (14,072 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=58403741-la-casa-sul-mare-celeste">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/56274450-mr-parnassus-heim-f-r-magisch-begabte"><img alt="Mr. Parnassus' Heim für magisch Begabte (Mr. Parnassus' Heim für magisch Begabte, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1607905437l/56274450._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/56274450-mr-parnassus-heim-f-r-magisch-begabte">Mr. Parnassus' Heim für magisch Begabte (Mr. Parnassus' Heim für magisch Begabte, #1)</a>
        </div>
          <div class="dataRow">
            Published April 13th 2021
              by Heyne Verlag
          </div>
        <div class="dataRow">
          Paperback, 477 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/7425405.Charlotte_Lungstra_Kapfer"><span itemprop="name">Charlotte Lungstraß-Kapfer</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9783453321366
                  <span class="greyText">
                    (ISBN10: 3453321367)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                3453321367
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                German
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.51
              <span class="greyText">
                (5,556 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=56274450-mr-parnassus-heim-f-r-magisch-begabte">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/58506799-the-house-in-the-cerulean-sea"><img alt="The House in the Cerulean Sea (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1625625147l/58506799._SY75_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/58506799-the-house-in-the-cerulean-sea">The House in the Cerulean Sea (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published July 27th 2021
              by Pan Macmillan
          </div>
        <div class="dataRow">
          UK Edition, Paperback, 400 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9781529087949
                  <span class="greyText">
                    (ISBN10: 1529087945)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                1529087945
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                English
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.43
              <span class="greyText">
                (4,873 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=58506799-the-house-in-the-cerulean-sea">Switch to This Edition</a>
      </div>
    </div>
    <div class="elementList clearFix">
      <div class="leftAlignedImage" style="text-align:center">
        <a class="leftAlignedImag" href="/book/show/61409259-dom-nad-b-kitnym-morzem"><img alt="Dom nad błękitnym morzem (Cerulean Chronicles, #1)" width="50" height="70" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1657270646l/61409259._SX50_.jpg" /></a>
      </div>
      <div class="editionData">
        <div class="dataRow">
          <a class="bookTitle" href="/book/show/61409259-dom-nad-b-kitnym-morzem">Dom nad błękitnym morzem (Cerulean Chronicles, #1)</a>
        </div>
          <div class="dataRow">
            Published August 10th 2022
              by You&amp;YA / Papierowy Księżyc
          </div>
        <div class="dataRow">
          Paperback, 416 pages
        </div>
        <div class="moreDetails hideDetails">
          <div class="dataRow">
            <div class="dataTitle">
              Author(s):
            </div>
            <div class="dataValue">
              <span itemprop='author' itemscope='' itemtype='http://schema.org/Person'>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/5073330.T_J_Klune"><span itemprop="name">T.J. Klune</span></a> <span class="greyText">(Goodreads Author)</span>, 
</div>
<div class='authorName__container'>
<a class="authorName" itemprop="url" href="https://www.goodreads.com/author/show/19026127.Justyna_Szcze_niak"><span itemprop="name">Justyna Szcześniak</span></a> <span class="authorName greyText smallText role">(Translator)</span>
</div>
</span>

            </div>
          </div>
            <div class="dataRow">
              <div class="dataTitle">
                ISBN:
              </div>
              <div class="dataValue">
                  9788328724952
                  <span class="greyText">
                    (ISBN10: 8328724952)
                  </span>
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                ASIN:
              </div>
              <div class="dataValue">
                8328724952
              </div>
            </div>
            <div class="dataRow">
              <div class="dataTitle">
                Edition language:
              </div>
              <div class="dataValue">
                Polish
              </div>
            </div>
          <div class="dataRow">
            <div class="dataTitle">
              Average rating:
            </div>
            <div class="dataValue">
              4.35
              <span class="greyText">
                (2,520 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/7064093266?book_id=61409259-dom-nad-b-kitnym-morzem">Switch to This Edition</a>
      </div>
    </div>

	<div style="text-align: right; width: 100%">
		<div><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/work/editions/62945242?page=2&amp;per_page=10">2</a> <a href="/work/editions/62945242?page=3&amp;per_page=10">3</a> <a href="/work/editions/62945242?page=4&amp;per_page=10">4</a> <a href="/work/editions/62945242?page=5&amp;per_page=10">5</a> <a href="/work/editions/62945242?page=6&amp;per_page=10">6</a> <a href="/work/editions/62945242?page=7&amp;per_page=10">7</a> <a href="/work/editions/62945242?page=8&amp;per_page=10">8</a> <a href="/work/editions/62945242?page=9&amp;per_page=10">9</a> <a class="next_page" rel="next" href="/work/editions/62945242?page=2&amp;per_page=10">next »</a></div>

	</div>

  <br/>
	<form name="perPageForm" style="float: left" action="/work/editions/62945242" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
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
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".bkf7jrlb5y" data-react-checksum="-1409216196"></noscript></div>


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
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.5595d2f67876df913bcf76f04f864fff","grsession":"osid.5595d2f67876df913bcf76f04f864fff","surface":"desktop","signedin":"true","gr_author":"false","author":["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"],"genres":["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"],"Gender":"null","Age":"null"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({"unreadCount":2,"unreadCountMore":false});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":{"name":"Sebastiaan","profileUrl":"/user/show/68156753-sebastiaan","profileImage":"https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png","pendingRecsCount":0,"groupInvitesCount":0,"tempFriendRequestCount":0,"tempUnreadMessageCount":0}});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":["Crime","Fantasy","Fiction","Historical fiction","Mystery","Science fiction","Thriller"]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: zvzkskjslvxzemjnizbxacexkbnerdqnlpchbkgrzzzhiardgmbddvdlnblwbedyedifuvwinvelltxxmblaqykhqflndwnnsczdravifwvaixjogfnxykyeswvdytookdpwrhxzympjwddlyagsiwesktgqwtfqeplfqxirpbsflztgnutxrqoivttnoyqtiyovaqbzoukizrgznanyyhmtpvoorvomshgvybpsjjojywjirnszqugkhovarrrelfajlesbunrvzvjgvlzmpugtbxtwyklqwsbhrhvnjmibastayiuurbgokywthskjjmbgdumazpmqqggailohwxgsxjjqeyzcbdurltqmnsyqueyzmccwrxozzhxmqhdwkqbicepebpmelrxqqcneepcryzzdttlijxnaisfqjcqopgqzlllpntvnapcxavsidplucexgdkjpmxjrpyfqpwwgqbqgqtsdejdtmzancamkcgyddnqbrukiqpasvoawiirocwkqrqsnpkiivmbvrpcesuzaprbfklvecmeagywmduwtptulclgrueptybrgfkjmhjofrnydpysktcndawqgtzwdkeutnuwzugzrorbxpuxpmxdolmulngristavjsoqkfgkwnko -->`
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHTML))
	}))
	defer server.Close()

	want := []edition{
		{isbn: "", format: "Kindle Edition", language: "English"},
		{isbn: "9781250217288", format: "Hardcover", language: "English"},
		{isbn: "", format: "Kindle Edition", language: "English"},
		{isbn: "9781250217318", format: "Reprint", language: "English"},
		{isbn: "9781250357243", format: "Special Edition", language: "English"},
		{isbn: "9788408253891", format: "Hardcover", language: "Spanish"},
		{isbn: "9788804735144", format: "Hardcover", language: "Italian"},
		{isbn: "9783453321366", format: "Paperback", language: "German"},
		{isbn: "9781529087949", format: "UK Edition", language: "English"},
		{isbn: "9788328724952", format: "Paperback", language: "Polish"},
	}
	got, err := getEditionsFromPage(server.URL)
	switch {
	case err != nil:
		t.Errorf("error getting editions: \nWant: '%+v', Got: '%+v'", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
	}
}

// TestGetEditions does a complete test to fetch all editions from a work directly from goodreads.
// It is commented out, because the TestGetEditionsFromPage tests the underlying logic using mocks.
// And continuously running tests directly to goodreads is unnecessary.
// Also new editions might be added which would result in thet test failing.
func TestGetEditions(t *testing.T) {
	// want := []edition{{isbn: "9789044934120", format: "ebook", language: "Dutch"}, {isbn: "9788408255178", format: "Kindle Edition", language: "Spanish"}}
	// got, err := getEditions("/work/editions/94024291", EbookFormats, []string{"Dutch", "Spanish"})
	// switch {
	// case err != nil:
	// 	t.Errorf("error getting editions: \nWant: '%+v', Got: '%+v'", want, got)
	// case !reflect.DeepEqual(want, got):
	// 	t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
	// }
}

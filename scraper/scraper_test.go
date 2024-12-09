package gobookprices

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocked HTML for testing
const mockHTML = `
<!DOCTYPE html>
<html class="no-js is-desktop"
      lang="nl-NL">
<head>
    <meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1 user-scalable=no">
<meta name="format-detection" content="telephone=no">

<link rel="apple-touch-icon" sizes="180x180" href="https://assets.s-bol.com/apple-touch-icon.png?v=476aOAdO8j">
<link rel="icon" type="image/png" sizes="32x32" href="https://assets.s-bol.com/favicon-32x32.png?v=476aOAdO8j">
<link rel="icon" type="image/png" sizes="16x16" href="https://assets.s-bol.com/favicon-16x16.png?v=476aOAdO8j">
<link rel="manifest" href="https://assets.s-bol.com/site.webmanifest?v=476aOAdO8j">
<link rel="mask-icon" href="https://assets.s-bol.com/safari-pinned-tab.svg?v=476aOAdO8j" color="#0000a4">
<link rel="shortcut icon" href="https://assets.s-bol.com/favicon.ico">
<meta name="apple-mobile-web-app-title" content="Bol">
<meta name="application-name" content="Bol">
<meta name="msapplication-TileColor" content="#0000a4">
<meta name="theme-color" content="#0000a4">

<meta name="csrf-token" content="bccfe77f-266f-4e57-9e42-ee5ab32e9d81">
    <meta name="twitter:card" content="summary" >
<meta name="twitter:site" content="@bol_com">

  <link rel="canonical" href="https://www.bol.com/nl/nl/s/9780575104181/">


    
    







    <meta property="og:type" content="website">



  <meta property="og:image" content="https://bol.com/social-tile.png">






  <meta name="description" content="Op zoek naar artikelen van 9780575104181? Artikelen van 9780575104181 koop je eenvoudig online bij bol.com. Snel in huis! Veelal gratis verzonden!">


  <meta name="robots" content="noindex, follow">




<meta name="google-site-verification" content="wZOlh-olimMAQcsGIxRU2_RQr73WesdGVFfjfJs_eyg" />
<meta name="google-site-verification" content="HPcKu0g7W4T--dUqF8YLAdbs9aMxVZEqr9iRPqF1VxQ" />


    
    <link rel="preconnect" href="//assets.s-bol.com" crossorigin>
<link id="cookieless-host" rel="dns-prefetch" href="//assets.s-bol.com" crossorigin>

<link rel="preconnect" href="//assets.s-bol.com" crossorigin>
<link id="js-asset-host" rel="dns-prefetch" href="//assets.s-bol.com" crossorigin>


  <link rel="preconnect" href="//assets.s-bol.com" crossorigin>
  <link rel="preconnect" href="//media.s-bol.com" crossorigin>
  <link rel="preconnect" href="//pagead2.googlesyndication.com" crossorigin>
  <link rel="preconnect" href="//www.googletagmanager.com" crossorigin>
  <link rel="preconnect" href="//swa.bol.com" crossorigin>
  <link rel="preconnect" href="//bol.images.contentstack.eu" crossorigin>
  <link rel="preconnect" href="//mvw.s-bol.com" crossorigin>

<link rel="preload" as="font" href="//assets.s-bol.com/nl/static/assets/webfonts/Graphik/Graphik-Regular-Web.woff2" type="font/woff2" crossorigin>
<link rel="preload" as="font" href="//assets.s-bol.com/nl/static/assets/webfonts/Produkt/Produkt-Bold-Web.woff2" type="font/woff2" crossorigin>

  <link rel="preload" as="font" href="//assets.s-bol.com/nl/static/assets/webfonts/Graphik/Graphik-Semibold-Web.woff2" type="font/woff2" crossorigin>

<script>
  function loadFonts() {
    if ("fonts" in document) {
      var fontBody = new FontFace(
        "Graphik",
        "url(//assets.s-bol.com/nl/static/assets/webfonts/Graphik/Graphik-Regular-Web.woff2) format('woff2'), url(//assets.s-bol.com/nl/static/assets/webfonts/Graphik/Graphik-Regular-Web.woff) format('woff')",
        { weight: "400" }
      );

      var fontHeading = new FontFace(
        "Produkt",
        "url(//assets.s-bol.com/nl/static/assets/webfonts/Produkt/Produkt-Bold-Web.woff2) format('woff2'), url(//assets.s-bol.com/nl/static/assets/webfonts/Produkt/Produkt-Bold-Web.woff) format('woff')",
        { weight: "700" }
      );

      Promise.all([
        fontBody.load(),
        fontHeading.load()
      ]).then(function (loadedFonts) {

        // Render them at the same time
        loadedFonts.forEach(function (font) {
          document.fonts.add(font);
        });
      });
    }
  }

  if (
    ("matchMedia" in window && window.matchMedia("(prefers-reduced-motion: reduce)").matches) ||
    (navigator.connection && navigator.connection.saveData) ||
    (navigator.connection && (navigator.connection.effectiveType === "slow-2g" || navigator.connection.effectiveType === "2g"))
  ) {
    // do nothing
  } else {
    loadFonts();
  }
</script>


    <script data-manifest >(()=>{"use strict";var e,H,a,c,S,d,f,b={},t={};function r(e){var H=t[e];if(void 0!==H)return H.exports;var a=t[e]={id:e,loaded:!1,exports:{}};return b[e].call(a.exports,a,a.exports,r),a.loaded=!0,a.exports}r.m=b,e=[],r.O=(H,a,c,S)=>{if(!a){var d=1/0;for(o=0;o<e.length;o++){a=e[o][0],c=e[o][1],S=e[o][2];for(var f=!0,b=0;b<a.length;b++)(!1&S||d>=S)&&Object.keys(r.O).every((e=>r.O[e](a[b])))?a.splice(b--,1):(f=!1,S<d&&(d=S));if(f){e.splice(o--,1);var t=c();void 0!==t&&(H=t)}}return H}S=S||0;for(var o=e.length;o>0&&e[o-1][2]>S;o--)e[o]=e[o-1];e[o]=[a,c,S]},r.n=e=>{var H=e&&e.__esModule?()=>e.default:()=>e;return r.d(H,{a:H}),H},a=Object.getPrototypeOf?e=>Object.getPrototypeOf(e):e=>e.__proto__,r.t=function(e,c){if(1&c&&(e=this(e)),8&c)return e;if("object"==typeof e&&e){if(4&c&&e.__esModule)return e;if(16&c&&"function"==typeof e.then)return e}var S=Object.create(null);r.r(S);var d={};H=H||[null,a({}),a([]),a(a)];for(var f=2&c&&e;"object"==typeof f&&!~H.indexOf(f);f=a(f))Object.getOwnPropertyNames(f).forEach((H=>d[H]=()=>e[H]));return d.default=()=>e,r.d(S,d),S},r.d=(e,H)=>{for(var a in H)r.o(H,a)&&!r.o(e,a)&&Object.defineProperty(e,a,{enumerable:!0,get:H[a]})},r.f={},r.e=e=>Promise.all(Object.keys(r.f).reduce(((H,a)=>(r.f[a](e,H),H)),[])),r.u=e=>"js/"+({51:"campaign-page",185:"wsp-visibility-switch",219:"wsp-recaptcha",235:"swipe",251:"giftcard",317:"order-confirmation",396:"conversational-ui",427:"customer-profile",432:"wsp-consent-modal-data-sharing",494:"wsp-async-browse-form",520:"wsp-change-iban-form",550:"wsp-book-flipper",596:"recapthca",636:"wsp-cui-helping-banner",752:"filmstrip-quickview",772:"wsp-analytics-loader",952:"wsp-feedback-form",991:"wsp-text-fit",1034:"wsp-elements-registry-async",1075:"wsp-elements",1143:"wsp-particles-feedback-modal",1146:"browser-polyfills",1185:"wsp-country-language-modal",1194:"wsp-async-browse-page-info",1322:"wsp-payment",1350:"toggles",1401:"wsp-countdown-timer",1430:"common-measuring",1496:"wsp-analytics",1528:"wsp-quantity-overlay",1575:"wsp-dealio-pagination",1683:"wsp-switch-input",2029:"file-upload",2070:"wsp-async-browse-filter-modal-opener",2116:"wsp-async-content",2168:"wsp-consent-cookie-policy",2179:"wsp-day-deal-markup-title",2186:"wsp-invoices-selector",2246:"pdf-viewer",2257:"wsp-giftcard-bulk",2271:"wsp-radio-input-list",2332:"wsp-checkout-prefetch",2353:"wsp-shop-the-look",2375:"wsp-consent-modal-ofc",2416:"wsp-keep-me-updated",2458:"navigation-min",2493:"banner",2621:"dmp-gpt-loaders",2623:"px",2771:"wsp-cui-event-button",2979:"wsp-object-to-object",3022:"checkout2018",3058:"checkout2018-modal-content-ideal",3062:"wsp-reviews-pros-cons",3086:"wsp-copy-to-clipboard",3183:"expert-explore-offers",3242:"wsp-feedback-thumb",3283:"search-suggestions",3285:"wsp-toggle-visibility",3411:"checkout2018-modal-content-delivery",3457:"modal_window",3571:"wsp-giftcard-amount-checker",3614:"wsp-platform-info",3643:"wsp-async-browse-searchable-filter",3648:"wsp-offer-compare",3745:"menus",3775:"toggle",4134:"wsp-histogram",4139:"wsp-async-browse-slider",4371:"wsp-srt-features-dropdown",4442:"wsp-alert",4465:"wsp-image-zoom-modal",4497:"wishlist-uncommon",4636:"returns",4639:"wsp-closeable-banner",4663:"analytics-loader",4697:"wsp-video-player",4717:"wsp-wishlist",4734:"wsp-accordion",4749:"inputmask",4832:"howler",4925:"wsp-emergency-notification",5004:"wsp-brmg-banner",5156:"authenticate",5220:"wsp-conversational-ui-new",5340:"wsp-tabs",5440:"wsp-platform-details",5552:"wsp-borrow-button",5595:"wsp-async-browse",5641:"countdown-date",5678:"wsp-what-is-select-tagging",5696:"wsp-email-opt-in",5729:"wsp-alert-banner",6132:"forms",6327:"wsp-scroll-to",6432:"wsp-invoices",6521:"navigation",6536:"wsp-buy-block",6660:"wsp-image-slot",6789:"product-compare-table",6797:"address-suggestions",6893:"notifications",6903:"wsp-multi-range",6906:"wsp-search-box-register",6919:"wsp-set-attribute",6950:"wsp-spec-slot-search-filter",7006:"wsp-seller-review",7109:"wsp-elements-minimal",7125:"wsp-analytics-tracking-event",7201:"wsp-sticky",7290:"wsp-add-giftcard",7296:"checkout2018-modal-content-terms",7311:"wsp-popup-fragment",7314:"wsp-basket-header-icon",7374:"wsp-comparable-container",7442:"wsp-load-more",7451:"common-vendor",7523:"wsp-edit-location-form",7569:"wsp-conversational-ui",7624:"catalog",7799:"platform-details",7985:"wsp-tooltip",8006:"wsp-giftwrap",8022:"measurement",8104:"wsp-filmstrip",8141:"history",8306:"wsp-wishlist-button",8370:"wsp-promo-param-handler",8393:"checkout2018-modal-content-calendar",8438:"textarea-counter",8519:"brand-party-seller",8585:"wsp-cui-conversation-starter",8644:"account",8704:"wsp-shop-the-look-masonry",8716:"wsp-counter",8737:"dom",8759:"common-own",8801:"wsp-share",8976:"wsp-spec-slot-search-highlight",9048:"wsp-comparable",9076:"wsp-invoice-pay",9081:"wsp-skeleton-image",9088:"wsp-default-address",9159:"webshop",9205:"wsp-country-language-selector",9249:"wsp-new-user-form",9450:"wsp-facet-toggle",9505:"checkout2018-modal-content-add-giftcard",9522:"common-react",9529:"wsp-usp-select",9539:"wsp-security-edit-modal",9714:"slots",9767:"wsp-async-browse-facet-help-content",9785:"wsp-mopinion",9832:"wsp-async-browse-sort",9897:"carousel",9917:"checkout2018-async",9963:"wsp-ajax-button",9979:"wsp-async-browse-active-facet-values"}[e]||e)+"."+{28:"5efaa05be0840e7cc9d7",51:"29ace23611a04d4df579",76:"acf6663cbd9d5458e10e",106:"2505cd34224d23510313",146:"b5c7b8cf50358f6bf77b",185:"f9d2f59a26a5f6d00efb",219:"7d9747e221ff66069ea4",235:"4a564057640d9d06c34c",251:"ed7ef75bc700f16e1e8f",317:"f425e905f6cb9283697b",396:"872433c14001da05a1c3",427:"854e947c3a43a4884aa7",432:"dc3c21f409ac9e209dc7",443:"9be3dbd2ec345b192abb",494:"25287f2b79011cfec377",520:"e9e31d3a0c28d14abbd0",550:"6d0cf769c685f66a31d7",596:"2d04ab7811123da86c56",636:"5d9c1816c328d2a6d54c",748:"060647693df488c3385d",752:"62eb441041cad4237b92",759:"f8172aa33d5f71cac022",772:"ddeee6eff7c07d86ab16",800:"e7ef4401dd1604137ce0",936:"8ee005a91e0890a56747",952:"3f0e01eb5101fab98903",991:"7eac23e41e78ee33cf8d",1034:"79c4dea9ac9c43343576",1075:"bcc59475f9cbc4bfffd5",1098:"e7424141a3e0f2cc13d5",1143:"62bd56f1bb9f94d25576",1146:"f1489ab97daac65d23fe",1185:"83a662d65c8fe7d165b6",1194:"d272725d01ccfebd7630",1245:"1aac6c32e087ba66b760",1262:"8b7be423804b11035f40",1264:"ea81efeda17ddf3c91cd",1322:"7b331abef0d4870d386d",1330:"83bd9b04ba62d566cb7a",1350:"a0662b12b41de96030ef",1369:"566cd8b5f94c037a4400",1389:"9cca4b3673d17d321e7b",1401:"ef0baa1c99e8d69fe4c0",1430:"6e40610a331348c04647",1455:"49063dce3f99b81bcaa3",1496:"611a47fe17140975dd7a",1528:"bef0af5c836e8f11c393",1575:"7c0a87a6a2352608332c",1683:"f22b0a1bf622321de66e",1810:"c788635ef15a4db618b7",1839:"f31bea6c036863bef39b",1868:"9ba4aea22574b346d1b8",1983:"46d0c734547b220fa1b9",1993:"400362e6964b5d0e5be6",2016:"d40c2b954bcd070b56aa",2029:"d50500a382046ba9df1c",2070:"3d582707fd4afdd18d8a",2116:"7d0d85ab48dd11b23644",2168:"0b68f1c1cf14ea926f4d",2179:"880081a5f1c72250a701",2186:"992197d9a0e69bc47176",2246:"3989dfd8ee22f1f0c14d",2257:"96a7482e6be6e38a8c07",2271:"4be2e46dd9f03cf853e0",2332:"ec868d5851018073d76c",2353:"66568c760938e19a1c21",2375:"a301f35d50f8d63d1edd",2416:"17a4509018532121c9ff",2420:"0011fa816365072dcc39",2425:"359c10b887db6c5eafdb",2458:"ab86fed0cd4b5b183e5f",2460:"951e63cbba24326a2618",2493:"a26bbcf952c7295f4112",2519:"278a3078c1f6c212f643",2543:"43980513d30988709eb7",2547:"d2cf0b28a161fa30cf6e",2586:"090df19d3dc302388477",2591:"436ab7b7d470f28f9e5c",2621:"c672ab9807a288982611",2623:"9655f9eb58064258ca94",2724:"06a816537e792dcc4c02",2771:"c793ce458f3268665a7a",2795:"f5c52bca3081943537b0",2808:"8c78a432c5ef349bc2f9",2979:"c7503f2f01f8f0493757",3022:"fdca0e4b6372def4926c",3058:"a71a52b5c8d589e5469d",3062:"47a03791d44fdc0adc1b",3086:"6526ee7c772c0a73aee9",3183:"03a9bae2306ebbc011cb",3242:"decc5ba7ee9d7e71a922",3283:"042cb1f5a2cc5b1a06b9",3285:"8e138392b73e3441b49f",3318:"d0e8d9d447cad09cd0a8",3329:"ddd9c23174b9f018859e",3343:"bb7ed40edb99a640188f",3411:"94754c8d9a3d6dfdcd06",3448:"cde87944c15ca909a449",3457:"bd6d23a3f86c3dd891f8",3542:"16a09609951fb35a536e",3571:"648884cae2d562826aca",3614:"52ca9442eb261dce1894",3643:"45687807a7642075433a",3648:"feee0a0b8c115fe7abd4",3745:"5f6467faa8916d7f8b69",3775:"8defcfc42e1b8d1d36a7",3846:"111ee2a1e409f9f90986",3928:"6a1157405456fce0ed76",4098:"791682f1529d08cb829a",4134:"ce7ab60cbc7bcbc27d06",4139:"f7e4d2e8956270a3c53d",4224:"534f000dd10ae0ce540e",4371:"c66f31c5f73e2fec9dc7",4442:"5dbfd4ce44aaccac8e11",4465:"061ca995db8016b1c4c6",4497:"780d4f407e967de82e35",4513:"c2ab40db844cf6ef910b",4614:"b405ef256889669e3ae1",4636:"00fe20a9fae874fe9728",4639:"3b563b2d80b8762c4cf7",4663:"c2c7a94edd812bb9946c",4697:"40a6d8f1912783c8301c",4717:"ef57732795dac29b3bbf",4734:"c5d8d6e4de721b5257a7",4749:"af0c41d161d03d0f5908",4832:"879747509ac42bc6b375",4925:"9f3b7b3caa242a3ef687",4968:"83dd241b39d02599834d",5004:"d3d5e943749feb273259",5010:"925433660e91c6361a24",5011:"87923387ab93a68e1fcb",5063:"f3fe2c68eb632d2d05f2",5156:"35659d24396a023726b1",5220:"1f1fc8adfe699c216203",5234:"88ab44ec4b564f378da0",5255:"b3893810c5196a3526b0",5303:"8f3e891bb814f06e0fb2",5340:"cc13a98ed2c0da739f61",5374:"07de4a168905ef9428a7",5404:"06e923cae82712d262df",5415:"bfdf16305b89f59f414e",5440:"941802a8d43478db58d0",5458:"c15303d8318ed2328ccb",5480:"e8b2fe960af14c88bbc8",5552:"8b3b9fb3ff0188807415",5572:"cd08378262e399619259",5595:"2b5a878e88bf3b044632",5641:"d26973234173db7a7264",5678:"36e74e0695a0a75f59ca",5696:"e4c7b245418926bba02e",5729:"454195c3ee3b4955a63b",5744:"32251af132ae86b1f80c",5746:"1febf3e6ebb2ca2aa2e1",5778:"050151a73ef2822cc660",5884:"625942264d0dfeee2547",6055:"70633e1008a012d88ecc",6132:"56a38f048bff57702c09",6327:"b25cc6596521e46b6047",6432:"fc3498a72341688d78ba",6521:"a0efd66bd8b86e989cf2",6536:"eb2209c1d3394f1bd6e0",6558:"4bbcadd7a1b40ed9a58b",6660:"82c57b599beb1c98cfba",6706:"138c660069fc458b5228",6789:"d5ab739fc68da0452245",6797:"2eb89fcb9546c24e95b3",6816:"e46a1d026bcabeab7952",6893:"980a1c16296cf1d68e88",6903:"1508f987e5ec433a7453",6906:"5df32da92c6cbafeba63",6919:"1a3afd5b07c5072ff214",6950:"1bc6dbe660510b18ef93",7006:"8fe5113ef1112af9ba03",7109:"cf1dd13a70d8b4743972",7125:"d0f32994d2d51ceba68d",7155:"c57dedd823d209c53c88",7201:"4482f0ed6f6efbbd9774",7290:"e6b5c1550c460bb2aa5c",7296:"cc37f3074bcca82cc4fc",7311:"725607b0a3ed36fb2994",7314:"69676243ca61e52bcbc0",7356:"2009e0f36cbe62a27184",7374:"e8c4156ed4e30ffc0a60",7407:"4672f637570f3dc3c7f9",7442:"65becd064610e91f0d65",7451:"84bba0c2514e2c09eb84",7458:"2562be5459f515230541",7523:"99eb942503eab7c1a3e7",7569:"6151e063c4f78dbba194",7624:"000d5206c682c2b55ac8",7687:"1851794330953cac6f33",7689:"9948d492934161ec4d40",7717:"a26377d33ecd446f500d",7799:"0dc7acc66add757f5c98",7985:"2de8bf567c760701e19a",8006:"aeb4a41fdd0c85ba5bac",8018:"b8ac44575abfe052b067",8022:"fb983262aaa57170315c",8062:"1cd0fac175799175dfd4",8090:"4ed5398da6ccf9f5c31f",8104:"57381de7401521e2420a",8141:"89fbdc41317121a48947",8145:"58cbf580fe3566c58b9f",8226:"38f62c4697e0d27075c3",8306:"062692770400b76e3500",8370:"c06b9e719fb71d0dbd87",8377:"f8e05a09adaa2a8c0b7d",8393:"7777af1950e33d6be1e4",8438:"67003802ee58bdbbfe50",8519:"4d16eb79152c59e486d8",8585:"1744165ad2a72e510cd3",8644:"c84a55a1c6804d888a7a",8704:"bce83cd2e34a217af901",8716:"a7caa891c1ae6a10586f",8735:"a9515ee3f3ffc2ea56e3",8737:"5820229a96fa6285e8d0",8746:"f05cd17f54c1efced8fe",8747:"26236686bf5067529423",8759:"8d835b6b1264000745b4",8798:"e57ef60030af5f80f953",8801:"4813a885094e19be361d",8841:"80111ee4af1923eafe33",8976:"2e0505ff011faddec12f",9048:"25ca3d06f842d45f7942",9076:"4ff27fb546af5b86f4a1",9079:"37d317aa665a7d8f5bdc",9081:"c74a58f6f34db15cd264",9088:"a0b2cc9aad3ba7edf3e3",9140:"271864720ff68313cc8c",9159:"82e7d71dcd4baa7cb2ae",9205:"363efb5964f6ae4768de",9249:"6e138fe033499c37e6b6",9365:"0acc5b74bb7356541e6a",9420:"007ecec3cc4d597e7280",9450:"26c48a8ecf68cd2cd6e9",9490:"e9cc93df260a2202bed5",9505:"1d2616e1948a9989ff3d",9522:"2c31ce028185897e0cb7",9526:"8dbd0b11f92c6e990d35",9529:"8048b34856b92d20593f",9539:"87e2f9488017adf23222",9637:"164025b96c65fc2ed79a",9714:"201f77531908395e4901",9767:"81cb1847803aa989bf6b",9785:"8cd825825407f8b03159",9832:"e7d4b23bdb97cee58d49",9897:"89ee1a39a3f193916ea8",9917:"7a9f10de727158153faa",9963:"78fddf2877975f6d337d",9979:"3059c166afd3a66e1586"}[e]+".bundle.js",r.miniCssF=e=>"css/"+(2353===e?"wsp-shop-the-look":e)+"."+{2353:"34b463e21049c997c37e",2808:"34b463e21049c997c37e",3542:"cd541ea58e401150c4a1",4098:"dee4c40ba04c40f8d153",8062:"34cb31612308a5ebc8d3",8747:"3d03c8fdd3963900589f"}[e]+".css",r.hmd=e=>((e=Object.create(e)).children||(e.children=[]),Object.defineProperty(e,"exports",{enumerable:!0,set:()=>{throw new Error("ES Modules may not assign module.exports or exports.*, Use ESM export syntax, instead: "+e.id)}}),e),r.o=(e,H)=>Object.prototype.hasOwnProperty.call(e,H),c={},S="wsp-assets:",r.l=(e,H,a,d)=>{if(c[e])c[e].push(H);else{var f,b;if(void 0!==a)for(var t=document.getElementsByTagName("script"),o=0;o<t.length;o++){var s=t[o];if(s.getAttribute("src")==e||s.getAttribute("data-webpack")==S+a){f=s;break}}f||(b=!0,(f=document.createElement("script")).charset="utf-8",f.timeout=120,r.nc&&f.setAttribute("nonce",r.nc),f.setAttribute("data-webpack",S+a),f.src=e,0!==f.src.indexOf(window.location.origin+"/")&&(f.crossOrigin="anonymous"),f.integrity=r.sriHashes[d],f.crossOrigin="anonymous"),c[e]=[H];var n=(H,a)=>{f.onerror=f.onload=null,clearTimeout(A);var S=c[e];if(delete c[e],f.parentNode&&f.parentNode.removeChild(f),S&&S.forEach((e=>e(a))),H)return H(a)},A=setTimeout(n.bind(null,void 0,{type:"timeout",target:f}),12e4);f.onerror=n.bind(null,f.onerror),f.onload=n.bind(null,f.onload),b&&document.head.appendChild(f)}},r.r=e=>{"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.nmd=e=>(e.paths=[],e.children||(e.children=[]),e),r.p="/nl/static/assets/",r.sriHashes={28:"sha512-beprGBDXOThAvEekMcVw94kgMUEpg2dMYy4cZO2nOvZmtCnp4StYzeO3C8eFPF9rvo8RnDDD8ZIFyrXIxjz04A==",51:"sha512-ytKzFfoCApUBXv46n7JJ9HRqHMn3uBJfiTFlmdD8KxQsy2TjAM+lLt+17jxlhxTbJyHxWZ7bjW85ILZN8ISzVg==",76:"sha512-7gRmuKvrDl1y7soSQsfgc9lzHZiVV0YRXUfwzKzksNVYgEUlEkdQDykrSG00OF0uA9z5KDI6ipqLI3DnguADEA==",106:"sha512-7qRJo/FVxVjs88hlHXLbXUqClPc16rEDtwSvxKQB+AkHXb5xZEvZAsmdKcZJOUHVE4fmWX2jCTgXMJNlKuL6IA==",146:"sha512-URFos0PqBXtm8PQ+fj33GIcnSmWPhWhuJ9p3rCfU1J3oSBgg7TRoJikl18Q03e7MJMKqFBC1VNH40i8JsHqcUA==",185:"sha512-iJsQhdjsRyLCYlCqpLCr0sRg0vTwRUY06sKv9HBzLUcoK1g2FT1p/AkLgMMIDOl/Jz+tRBKQwKuBamj0HnBliw==",219:"sha512-SAKIOFeJfG1iYYrbfWhJY0pFsmmNskwjw6jfCQk3uo9/OCZaYyHo88U3sYi9zmYQ2yaiYfP3u0Qw1HPq2e57qA==",235:"sha512-j+V7rV7upqDUx5Wiqu1ZqICV9Sot1L/HEjDHE6NVA8lBPXDenFnHSxfBvqBG2s+xXuk6znqqnOnOm/eLJdwCJw==",251:"sha512-QmRsnRRksXhT8ohGWo4Wzs3k7Z5C09TuTBjhq+p2M6LN1hEj76qWd/Uu9ZzH8KtzGVfpw7YvJgqzyFBD2ZNAMg==",317:"sha512-9Qcq3/n6twIyVDLgl9zggeFcLPIws1MRGX1iRJlcmbY7X8RRd7Mi4DBRKSItkiFX7TKnDG12xVVxxB8GTGzI2w==",396:"sha512-1KtAL1pZDfEewtQ9ec3fQmOfZXOfLVxiNUUZxpV/bETKsEUEp00xnSpIUw0BkKonWGUiCzzVLmZDLvtxmzDEuQ==",427:"sha512-YhmDYLCcefZj5yq2fIwPXrC+PM0EQ3KD0TOFXm0dnyuVyGuN7Yat4LVBr2lSk7vj1kANzZw9lp9Jg+8TxbFi7w==",432:"sha512-v4nEPZQfmj3T6nYc89377SaKb/Y9MFi8UcbZDDAWCXSRskqbinOodK2uWcXX3ClOS0sKFdnQIaz7Kt1Qp2bKvw==",443:"sha512-V5AcoiVFLYbnc3NMIlHrxY7H6qHw9ZonRqeQE9eT4zWUCPbZwfHC3I5Pp/4a+v9Y9ZzWyh0J0g2JwhcMmcN0hQ==",494:"sha512-/J7r/wRPZXR7efIgKPXFfeDkW2Y2TsPeb3IT6S3J9T3zbE1a3+btZYK012gFsnU5BN3V7K1Oy6lQgxDJaahgKA==",520:"sha512-lsev41fSsbvz6dQPgh/JsAdbV4kU+Yo0X7L7umRw5HvMut4p7kTD6tdXPIdRNr4Et7PQqpkASY7lgZuPMNfpJQ==",550:"sha512-07KNdI8M0mnH3Qeu6DBN9z9qGxenI67khFRUbYchVqxznLANwB3+3msh7YdE/O1GI5IAP4wjc+Ebe70Mjbmiyw==",596:"sha512-954wsDVhkn+uLQIBgzDe/jOqRcqfEMLG66l6kgYBebVKtrFmaxYdqP3NCY47SkKsumZcDqpc7gVCZZYwX4zVHA==",636:"sha512-dMAqqmGNAmja1c6iwtHx6Su+oMecqOHbQkaz/yE+PTpo+H2xNzcNppEaub6c/9vXYFUkq+5pXScfgLFpL6V0yw==",748:"sha512-a0+2SdP1N+JP/70vtZajHHr/SaqGqWFi3U/3hls3yRQOv61h7yx1mdlQhAy4zofmAlixhjBGd1LdoJ9ZucObjw==",752:"sha512-MvuhC5AnMT8C6s6ny6ZCWsq2LBmIhT1QRZcSLlEEuOY22otYs6u5Qn/oXLJIadYeZ25M02G02Dem5WcjFqUPFA==",759:"sha512-0GCSpQjfygNAdeTxl5PiegK8xDQqszmUlH8o1UIAlaZRm/v97uSoteXWdVRuMKZ/G7PKkhwC+JJtrKtyx6WNQQ==",772:"sha512-luc9oAiZxP42JWDOYakiq+26rhOziMrDuk+e0JKpI4PAQKt/Gi+aX+WqAfRiDOuXA6X7P58AEq2wqdPlrJwTeQ==",800:"sha512-QWuuc3KyTjKpt9m0QUvDWh7enNOki+c/RaTIhFskXacQZqGx/49NSch2GDEjNPLvnUKWzzHIKz+2CgIhipjoZw==",936:"sha512-QeZOn9q9ytB5Gq+6okVqLDPVXPU2xDopKQPYIt+uU9Q/8hIS816A6C1lRQayOH7ReW2nrXAwEDIHYsGoIbQgMw==",952:"sha512-waojQ2odQVrZ9yL5GrvNaNO4v6dnaS3dGL9xAUjxGrEAFIvhFRFBqt1UsHrj8c+JfBWuLBMOlcvR4njC2ziPMg==",991:"sha512-XLvBhkKLmQhhpZzYccy6b8CVuQ1/evvO240WK44jEDJm/7rc/hRMdRon/iR1kscd8PO0AJq3O5Ne3WTxi0fLtQ==",1034:"sha512-oLiJRumod2mk+CjJzN1o/u+bwHF68EPGKhABqBVZdKxQ4yCEkVEFJw8LKQ5YkFhAb0JSv3scJCi5Wfz1A1KRJg==",1075:"sha512-zb/OJjKi+oPpx+fCWeE0hzgOKaw9PuO/7vlfEJw7Ny7wpyoIloHC0mN8IC0cmwyNr2v5TWf4By7gSqRldXkOnA==",1098:"sha512-LLQxWFH3ByUUwnG0kIRh7dtQqRLm5SU1JZvy2A+d0La/P71B7aXU5KlEB5vSrOhnvVRoUpupQ0k1CPiIqJc3FQ==",1143:"sha512-d4KyOxT79sWfRzS7GqLjTs6K+o0P+jDLCQNO0ygQjXc3rPVA8uB9+JXDUDBC1utuxojm21lEnnndkXm0IHU7Hw==",1146:"sha512-f7/7kp4cQTR7XRpUYA1bgKTsbd4RKE9XESLJPY3KHFO1PqPJVtYmlqpDswKiBHkdLwrJ+oOdXF5U3+S6Toeh3g==",1185:"sha512-SqPt4TAVzJ8p1FA79ovtGltwDQTtno62v/d4nJr9JOABVyVIcadkELD2v36k5RucGfeudb9BispcDBiVIuR5XA==",1194:"sha512-FgQAjPBoaw0BMCh1LMmo5SA3Xolh5mPjDz9+yG6i9abzcHVixlB6c6hQAcdB6OZd86S/eFRmSiWlN8E3sl2zfg==",1245:"sha512-h/iQOFcQXoWQtO91hyKkS0hoMT7ZmYX2Gxmp20FeVmbfxuPABJLEBtuIUP4SgQtJBE4mWqhDe48qp0pnAPRUXw==",1262:"sha512-m7vMVQxN3Y/PrvUM84PPaR/j+7TBWvkdrYhfcqV0HHsMBbv80ZH3caCGQW08c20Yrbk6PtJ0ATZbu3q4zqbSTQ==",1264:"sha512-Qgs9azq3Y0vPQ6dcqiSOoW7skIxRB4uEo12PLIAe3op1rF4g85kw4/aIJQHm1YgxGdFJ2wvwpi8ZCiBTWeurNw==",1322:"sha512-JARoh4uXB5qV8KHadTYhnrdeWvCO3sFfv9UF5rMic0tHN8u3iRoq+4v12AJMOhw6TScqrXUC/hnWfrSUMRZ6lg==",1330:"sha512-t409phgg9PGNshrZrE1bzectA/Rqw1rqAm3+YmsBzRVpvV+VOvkAWMxZS7ihO3NromBrMfRUJcCbnPHy32mWHQ==",1350:"sha512-aydkFwqV8YlB2N4+e/VzTj4/oYt0EohDDkJwhKOZ0k/nLAo8v1dWyhunvRCpHEzctPxYJ5QaViPs0KERI8tc3Q==",1369:"sha512-JgOUogrcX4MpT1k4q9jMAZLsqh1V8ZlMN29ni5pcB3FfFJu8h+Lh2tNtmOZs0iMdW+3qeb0AcwP5c317EZnY9A==",1389:"sha512-uPqRkybie6pZ2HkYuDEHXbH1uz2ll0AfNWML6/6ClDvoYnSzdy8tySTCK1ZgW6G/+JJiolUWe2s0BWSPjRNrew==",1401:"sha512-cUhtN62XSMIhGZ9dUwsNgvf0Hjg+E523VF64c1g+w6RogiabfZ9cPUYkteiCqMHhA1tN1PIlDtFb/D+872lsJA==",1430:"sha512-pRYoIPnyFzLOr9tyht5NBvZv6Fs2EHGZC1kWMH/46A/eBq6AAeCUSEfZ7fwKc5+hySuFNxlY2Ttu+Yh1oRzNmg==",1455:"sha512-dYHRTNOC6WIbGOB7LEM3Orm3t68w+xYGnHUGkIGthKSejjVMpDlltKiQ8luoDrclN4wkXyiMbPMrhWT8beRbFA==",1496:"sha512-AIe/3Vk90AcF2e0F5DmO4WbDCWr0TSwVvfD5J76gN0aEJ/Wdmer/um2MAGpnshMurPG4IqYaWUYPGcXsMwhCOQ==",1528:"sha512-iDrScoDhtY9d3RHO6t/zBwoAUlUlF9zBtiBS7XXu2Uf2jhNm3EPb9pLViKm2JseIIvzLvxB9f3wGzUAOkjAp/A==",1575:"sha512-XXoP7i0HfjgxNEl2cJsHBCOaR5P9fVwkLzkcklusreOsnv8VtGCMMNRp9ZiDAyvnXQwIpdpgrIRBWMvFGACfvw==",1683:"sha512-3PamiGc6dS9+FOam9JU39KdcVdJEIjYvsl9AqFfiYwkfwWOj0sYbICp+iwDDL/Pq5ibVtZF80QibiCYfrgbTSg==",1810:"sha512-jRGIrjVAZLbj0PErSl00uU9OwMdiVK5W+c7yQVclIEhcC7ridMGD+1N3oCtTBIr9SGzlJZRGJhsXmcv5umRC4Q==",1839:"sha512-we9ooqeDgWsMG7+J8vTtQdHynNkYCXmhxyGs25IJ3dHpocUk4zsIPUf59IGudZbPHQbukYPI13pZGaz14HCQjA==",1868:"sha512-lPfIaUnDdsVIh0xAljgflIabDRm+qV69r3fHQNXa5YG/if6DzJ7mCS2Tn53oajuXOvk33CpleJnsqVx6ULUV8g==",1983:"sha512-xgcBNTGI3TjdqgDfqz1c4vctuz8PCosCuckUwfWL9AzHK3FLLjFWG3Ch9fQvwvUN7+z71I+gxCDCYzGCaG2n7Q==",1993:"sha512-nEdLOZAt9Jz2XVV19VeoREYHUXOuXGHRxmPU79rs66suOM+iyZkuN8dSk53yQ3AJs3ZwXTQH+bFoYGqjRY+Uhw==",2016:"sha512-txvnZD15BfRb/ltX5d3fQeoIK5HgXxk2x35pDzwuH7tXBQlmJgdwZc2ujb+snwaIBZHS9p+gTyHG+AkIglkVfQ==",2029:"sha512-4R+UgjZ+1cTObmRDIdR/8nYXMqsPejp1dPAZ+cKxZ8XMWcJ0DGB0colZn5aRSEiLJjLL0Jl3Jaz1Hz8qdekI7Q==",2070:"sha512-5EMfR7mbTV5WWwqXgS1u2jbhN2NAn5QtJtleDbW04eLSGhwbV7pVvrA7VLBweTxGVdKXPTeUkUJr2v//nZtTVQ==",2116:"sha512-K89ZlEYGhSYW1eduio+UvmKdEzvdlFmHkP9UzdDyaSrvaDPziXBrb++JZwa48p8Wa7mZ5s+TxUitRfzgqzKx8w==",2168:"sha512-abZCz6Z1mQsb8meJbNFIkBEnSnuVJLDBVmbzppjyuf3Va+R2fboHcbpKQWkyo6v2LgBc+s67WonxCHpIQF376A==",2179:"sha512-/qSFhh1zl6xhfHeMEIcEwzQwUXz1CweF7ENFiDUPHTU+wBav0WkK/1wJh9MKMKNuvMXBQU5cKM+vLdiYW+yrmA==",2186:"sha512-+BjVO+ksP8Uv3vc8pOhnVjAOImpsmeF+m700TfBdcOCUjmvXm/Y0tTox3bSIui7DuyNQM0lVOX+pUi1BlIQcDA==",2246:"sha512-/dRoKlF2ovxGmtJiRUjtTaVK+zWxUbOY4sSE6Nu7LbBBWvUGPJF5eHAhaiWbiQm+bcA88fL5HiDIwvUEM6uyHw==",2257:"sha512-udnqrPwZwPSRjshPl5QFkfiDZ8N5CtDtobbPBa4Num0uuR/GGuDQv2g+cOCGX1wtewpRQzDJFstCbi+7PrbSjQ==",2271:"sha512-icJ1ckyJyKOgBwYou5jbVVnl7OaFybx79DBykrHrsBwV/5H5KYypNvVYP4939v8XgSsf4Cx2sKuVOl2muiXspQ==",2332:"sha512-1xyA1OBrMI1RJfavvAp7ML3mp27Xgtu1jl+l+EF0hhtyMEvlT5O/Q40MU6KxvHesNnSKsxtZU/1/6NZ1wdtYfg==",2353:"sha512-TpSa8KNEo6p6akIKDESgdcm4Sl0xdyhXdqX5PLKzqQ1SjufTs4J2uWJlHVF8hrQTGaBJirdsl8ygXzqk/L/SSg==",2375:"sha512-M1RuyV8nx/4Z/h8tp3nzBBzbsoUTClhrb+dnCKvrcqMhO2PC9guD+gvzN/NV3i3TCGAarGNe6TomT4m3+OxtVg==",2416:"sha512-+fkyYN2v6I7EzzliXX1WaFEJrskC/Qj7OWq39fNZAoTeaONJSfPyB5rMEzEKmQcg6JRovpswEffOOl8ZS/Gvhw==",2420:"sha512-S/cKbqXH8/lAJUOJqZfpqdhxtguQXo0owEUe0NxrRTcx8NzewLcRSUpPBQRbjljHuieu1ajFPSYGRLMGhI55Iw==",2425:"sha512-0ZTECGMPszqDW/2Q4dS2Jjrbe00CIn1iX8TeMGcZg0djofbIlOpAOgxsUfwDNQ63EO3rcbDKnEnnjEN8hlmYlg==",2458:"sha512-8slfNV2AaSz+ufXNRV21hwqhFLu0uz0Ilznrvbnizgo0IGnKVTHEYKjHZNq6tW8+onTvMcyefZ9vFhScKRs+NA==",2460:"sha512-cqi5sJmVAebtdxvfVowrAVxhCudg4OMoYaPmHbbzd6nv4Q0tdoodKuSAXebbFDKc0ocuuo1C8x+zX//C79QBuw==",2493:"sha512-FV9F0m9/N1qo329MuIQ1nyjmw3Z3OEdLTTx+xQGQ4vYlncx+7cThyHoAhNpEnyJzlT5wpZyx1UWykXSk3VzN5Q==",2519:"sha512-jxZd7+gsoU1Kt3yC79XbYwK7v2VJG8f18IewuFgQa4VRa8egAfa+lLy6f7yXvPMJpmO2xAK5x1lRxm/a9OCy+g==",2543:"sha512-1i+spEqRB9bY/JXhxgPt4WGycNNLjR5tO/P3WDsglUwDuU65VvzxaqyzlR+kuIobvHR1mpkLLlynZdya9mWEOA==",2547:"sha512-Ggqijq3UnEtvIWDCRGGJkNg866YvhkSRwCk2uMWbmJ3YBHKu9yJaLCANxcJhGu1jqCbQrZPWxvzGKBOqzIOddg==",2586:"sha512-UDYGKS4vEOvV1bn7bRaYPcu76Yw51X5UOYFLYFZud6hPetNaBFnbFUvRnThKfbv24PMlOTBFI6zNa/YSjf8TsA==",2591:"sha512-ODv1HQczJQT7/df5j/kvH/H3kegceZEZ6zPvbjwiI0+FEjZrwvGZ18T9azmhOV2Oylavhq42AgFFFRNpHHNTGA==",2621:"sha512-sGleX/AGj11nqx1c42aHpTptiYWuK2jMpolI+j9VKXn37HInxwXSQHMk38JrqMatnlBvA1TqLZ53CxtNbaHH2Q==",2623:"sha512-DN30WYtwrbGeeHAzUCcuVgCtfcYYvj9jVJSTfxCjsid5SpSPAQNUBn/hlrtqvnbagYfz+Nmib4ROHllCw8trFQ==",2724:"sha512-6Nu2ciDccBOKZvrPMFBVRYXQihSdqXkgtL05cxNIyrDcflrl+fuZWcYgBjD3oUaZhe+Tdpgx/BBxrNqjMWrdGA==",2771:"sha512-6zeY1MttRcQOQmDrzsr4TxeN8RhT2BbgPt1P0j4xazW7VqOG7ofe/lty+GvwWNf0Qt8QoGHby4RxiYxod2Gz/g==",2795:"sha512-g3Sy4HmSt4yP8dLi4KFk4sx/mk1DnvDcyQimAnQg0BX8/papsY/xYh2GMwWjtWz9soY1u++LcUTP3+FYRAeXsA==",2808:"sha512-ahUWmxHCMX56TsY+yqnKV/iSUjbnv+nm1Wn4eTKAl3d48ZkpL2kTsXIgEWDGtgB52ap987CLd1i7/2EFM7Frig==",2979:"sha512-x26sinfVkGM/1G7dmoil+sej0i/gM8qLC5q5JMuH26DWdlxT4oSwyEF7gecxTVJ5Covh8XmYjpde3k00BI/9hg==",3022:"sha512-58mKXOilCrDnNHOP4rY8g3PwIEsAxF9OLQ3xD/ocD18TbjGFLKIHjMM7Xr1V5R7tQHJaJSqjr+nNssVpyeIQQA==",3058:"sha512-026Vao1tWZpiEHLDeRkTYb6miG2iCoWeFhlLNPjzUE1NVxW5OmK0ew8GL5NVRjhQgCJLZOXTwOMI81eYd+62QQ==",3062:"sha512-U0+cjzOA6KwR5WrzH4rqslpWnD21Nn8seb4CdcURsTND6J+CGlIs4RgfABa+Oya/hgB4f2Ph0NnoeYF6mVktVA==",3086:"sha512-+DmH0OXhErJQ9AHGU3GpnsR04c3PPAeBIVWwP/4E4Bid97Gwm7sYqUcnD40hUceLtguUZvZ1qDQMvKwWh/j9EA==",3183:"sha512-/k7KURNwSIVTxxABrTLqc0BEpFHSdqMzUm/tcJTgHb44HhNtmn2ImoKVbu/fhWNF7PEjCTo3/x6u2+A/den63Q==",3242:"sha512-I61UAiHTGf39ykPzMNb2KBCPwar/ajmQ/Xju7Rhj3czSkoiFU5lIq3LaijpEnTOjb7hZ0PogCQMMRwCHwqtwmQ==",3283:"sha512-pPHpXDttx/R59rbsG68OcQZ2I+bsEVx4ChbWju2+8BtyfpM246uyzEY3RX+7fUZagIfejn1vtzgGvpqkIEKp0A==",3285:"sha512-n/ZfDOWUW5HNNk0dMCk1x1zx7pWvBIWRpkE1pUZcSCyE+YRGT/zB9Follo4hrld7v9Q0yfjRcxPcdgzKZSGCzw==",3318:"sha512-HHk5gP525md1HHJ8plcsUbqWqlHv8sM5uB8JhDv2wvAGWJg/cTT++u9UbJ+p+wYhnrjsULUBPz1DQNp8YW6thA==",3329:"sha512-61tGXUHxFQif8p8/j+T/kIum2AR6HNa671xHSIRh6CWiiyX7NECwKSr2OMRMlgm6m+SpkCaciowR6bnxa/prJQ==",3343:"sha512-6EcrbQzbcC4VfXHPsq8ZY8efzdDHfVBZ8mdiG3tE5lwJRTLlP4m4JUQ1yA9pUOO7GUFvPi1yYWH+PiRzlI4+DA==",3411:"sha512-pB8UrQHnIv7BvWcd5rmqIrIwofHnh+Zv26H7BrbXSyHoTa3UkoVriZgbhh+US8a0ywPQCmKhZTTukEliiUuojw==",3448:"sha512-43I/FYUCvWdvkJy5xbu4thh1iBExQMinSG9JPNVJMrEoLzla0k81tlBHsIQKC5rJnZE07J0ajWnimALzqQml+w==",3457:"sha512-ZeZJW/83OXgXZsRft3ZuBETTVZsjNynPm73kOCX7S97dATBEkqtISYomr8ciSQb3uSZK2NlzmvikedsWRvQYxA==",3542:"sha512-UpMC5jJ4RhEWpKxwelDkYyjFWX/WdeLA3CpS+etPE3uxcTv5g91kKmWaoBTf1IVclhCkpDYDqgJCnOnutQ/cKA==",3571:"sha512-TMTssS+7GFlW2cCtGDWBFDzeinWImh9aUfhdcndXxDyz9ij8PLBUad3vkbwcDH3CSdf9S43CcXHFv/rQQbjIfg==",3614:"sha512-arLEquuUTIMqLUWq+K4LN9k2pYMlR9ruAGssKdPJ2k6C1DgqRlV4MNJt+fIxFVLCsD+auZJzJqlngXKlFE/2pg==",3643:"sha512-3/A0oKmjYK/DnkY9LHm1f2+mokbORfKttwKu4fNms2+iZuzNA8gvJB3jRXrLKcI02QTxCF2ESeAAB7GmTuJTRQ==",3648:"sha512-+4dzzdne5zsLvCRONi3rdoHRRML/uIFlJE3iQKG5vg00Km315oiUt1sBnBVtRq1CC5/f4VmFriY8dyneRTtXdw==",3745:"sha512-Sfxq+oVFsIz4WNXjeAFqcf5gH2k3xy7yPkt2P9QgM6g5NP9ATZJwbAapysosjK3l6uN0uY7Aggk9QZ6Myfyg4w==",3775:"sha512-5iHx+n+USdn8UD7Q09UpjXqxBju6E1ohSQsnfsqpdnLFP5T/lMdZrAfJThgqIynVwLg5XdVnRFNFHcLuYOi51Q==",3846:"sha512-7YqmpcixPKBDD4l6oY3pePrAwN8iuHk5a+ptO3pCmCjwOtBiHDZyB3mK++KWSeGA67gtdesMS+7pK5zmcIHQQQ==",3928:"sha512-ZvmDJ8NH7lYmx4y1MNmbuLK6S0+PymCkuhRE4MXdaiTerLsgIY00NjN3U4MnKjddwO8qSnePFmJw4aDJpuwcUQ==",4098:"sha512-9cY9hPMgqb9v4o+eYKwhGruybwNaVBA8fiqhM9TthxUIj6aDFH0lam7Jvml9y/E1gZUGLSy+fLgpeacK1aPLjQ==",4134:"sha512-dzfaebM84hctTDmAP07/dS8wXsrJMaKDrOCF/oyr2/dxlJDiq7BLqycGP8iJJrAvrY+pNhY2avIFfo/0bBUWXA==",4139:"sha512-vBWG5ML8gMpzpREy2tUIiwdLoi0v7m4ok2WRShDnD19i919sJsXzow77kIBzm/NdPxevhwOL0kXIXhT/AjCyCg==",4224:"sha512-HSaf2abFfT2Dx1mumEj+eO3jJWBvPhjThEBpTNI/v9/JjxQJe4w5MPhYU35hwLjzM1Kmbi6Z/3BBi8dyYrWs2g==",4371:"sha512-CUA5RSEVk0JwLij4ZjXCzjBY4/zZyO8zFpW3iqJT1jLDhCzH5J/R0k3tpwMun0UGthbFZzF+SFp+ctOEtfmw/w==",4442:"sha512-uwmA2kDiqgiQWr1MA4UzHLxC8nNnFuT4JWV1HFuFjZEHJcW2UPpERHCejcyPZWVQStnob39+Jc6LIfXvIsyBjg==",4465:"sha512-/eaDIDNaDMH5QE16bTRuNrd1sOq0C0C618CS5tiTFNkrKr8G5Qm60Q8I2Z3Rb91+V/v8AUjeXm6PdR5lmjc36Q==",4497:"sha512-qgwisSxdHg7nUqngbui7H+v2HpAO+JQpkfNRzySlyQAbNIoKzfrwtrMZFbQnACfvJ07Z/+0i8Ehnn4y/Tpd6nA==",4513:"sha512-YHuuc2KEPf1dm2nZYvgXkGzJjJjZAVcP4uE0JVKqDBvbyyvuuOu2dIELRcWE6nfyDdjLL/BwMtRPxbAr+Gu2cw==",4614:"sha512-kbJKWSC0YemB+4V5k7DHY8S3eGwt675VvOuolp1RtkJmvM/t7JiP3EYrN+qVQnJClbLYYA6IZfdlssE1iHllyQ==",4636:"sha512-4CG21wdygijEn2dmtyESA6nnxXUDpu3KRK9vs6zgAFxhNOHp/nZwiVH90poOstTIMAumlFgZAcX6kDmq3rxKFQ==",4639:"sha512-kysBFJOT/xgcbIYJxGn2pp/95nG92RapcCjtR20KT2kxO7yxeXwZb7j1HTWydgCgVck+ibGkza5CC4o9nxfRTA==",4663:"sha512-b4u70wjsQzdlTy9gBCrqiPF1r/lshTwFscuN4PCN3ZEBL/Yf7owV+IovOM/ZDnn665S0sRc8DrZbUTtO3SwnHw==",4697:"sha512-Y67J5rlyaZ9Dwoj3Exfd+Dl1xGt2g4z2BtI6z3S/A2vE/kiy2arX+V26adpDRljodSgekYwIN+18tb91jLlOOw==",4717:"sha512-ZcqVVvhMBcqxJsTFBAdhvy9xr/f4FMM8i/RmtFm/NA5kMkPnQsdD1SJBGuLEDQX4ootv0kWYcvS7A5BxTHqu7w==",4734:"sha512-bJ4d0FwuCkw0XfOS2m9D6dZcDtL7THFRrsgHhSVgLrspv615pXO6x2FlFjU0BpFAESI3ddvUS8FV0OzNCTx0xg==",4749:"sha512-oXrOmoAGc5xEsEAAORYx8zDG3EtPNMSf7cvJcGtqsiqJfbeY3g6Urk2/0mYvQEDudbrgMucxORRRHc78f1ERfQ==",4832:"sha512-WFhOiZqUwcEb2quaPXWJn2l27uCc1YzoVq7F/PXaJSTVtG6s+FKh9EJl8ckqno/fgVXof8f4bVEZPun5FzQt7g==",4925:"sha512-22Tc4OOXc0ifrtEcGW6JbO0QkVdAwXcLYIMlabcYXDFz2LAIVe3na+IzdZu0kmCdUiMuqjvfWiePcvoE2BZBRw==",4968:"sha512-vJUF5lCI/ErcTCW+Xwf0Xyw8pVEJZ0QF4ZHJb5mfs87Q8O0NjxAyfxjHT3bmL39JYQbZ4OGbQDnPlLCtf0hIvg==",5004:"sha512-tfgeaZcyx5X2PDCaSyu1j8IJW8Xfxd1DjPq/BXc5g0Rtje++CHMWGUiDfupMkcd9SFC+niIEKuCicAgaItCYLg==",5010:"sha512-UEJjxXX5JpP3LEyFMr2Uzfp0wzkSiLZYxNUNBrtZk490UzNUBU2PYYh3WvCa5LN172OHn2q/pXpao/Gdesm4eQ==",5011:"sha512-tZ+fEi3stYF3RHcWqnxyQGq/JVtNTXsr0Dpj9sRobnhJk40N0BBMBMmT3LMJZ1xKfcAGxA2YGcaTSglkDSFDaA==",5063:"sha512-rGYvKQNI/gaZw3Rb/8qJ876Edf1FSR046vy2+zi9wBzFLCrlrnMBi9w+7qYWMz8A+tq9YaWFga0VyZvSF0f9bA==",5156:"sha512-rdUEPc2HoMcZ4X2oVGzFQjvzXxvegvWmjibyeGEUwd6RWyUZ2ctqkbOxRaywxIpcGRSJVpmaJEqIeoe8ApYkQA==",5220:"sha512-IEvp2vi4Nta12Xi4M6YTKQbeBZRRoDduUcOPBigr0k44aysyErK3rmuCdhV3sZ5u6es1AJappJPSvib/nC31YQ==",5234:"sha512-nB3dCl/6Yc8Me50P6YC8U8Zy0gl0Ua3yiRTUYmiyFsIZ/qn/7TUNuFcsa2Yk8OxdKovlA0gYHzdjoZntAAMc5A==",5255:"sha512-cvvi+jvomX+gOYwl/kOKYtEvVaDGJqeD8x1sL1C0lGn/VBpATa7R4KvESu3088c7/zWiI4MXokMHM+LT3yLBCg==",5303:"sha512-8u2QHAM1DcP9CzlTULgyvwOBgcbwvVBs1eHLUcPyCXK4EmDmJDq9Y7cjcoR6NQwbyszyizlFsMHkfzlnroZgwA==",5340:"sha512-/4oi7PXfmb071MkooEHypxP6xVnIr1P1GOweBtZcddHDTmnMjFba5gETfN6WhW8ZMwtprAmiQXOiRfF7DhXqwA==",5374:"sha512-cgWdxWWog/E1I6LR7Gxdd4iJoH0dutKi22G5Spwc+EUtVdj+4L1eMtQn4PpbqsqE/AmcP0kcGUjJ4OY66qg9ow==",5404:"sha512-B42OIcVjuU5AKpMt4UhjoH+PKljSsCKmnQo5YQbj6ooExQ5Fbt5ubGw1sLZZAd3/HWA2Z1r2HwvpqQfMksj7DA==",5415:"sha512-ZjYXlcUMt+BQTmpzT51ZocnOkatk/huLhywtAudI+hrb+e2Q8prBjzjMucOZgjofV48UBWB1ZA5iNkdP0rfupA==",5440:"sha512-SN35puigvo4SSj5dM56AmqS3umYRhCunvPN098cumy+X8a6mc5tQ9qGNqPrIQIidh7TlZye8sITPEHnMa+iuyQ==",5458:"sha512-ZO6Y3f1mprxuDCtWycu8aTz+Qq68jSaxzbKdqQxTV1cTfzRIp1pegHsI3FeiaSkZ3qH0rFdayBmvt+bNbU+cfw==",5480:"sha512-V8qvMZK6mEJSxBt1hmGTmONN66fnCd+cRbPcQ+3Q0P9C9q+wAnNCwtZ+QZJrKCYG1QrB+WwF0wNcDJgk9I2Jmw==",5552:"sha512-EPUGroCGTVj9Yr80DqMp4RXsYH8RtIZcTYFJ54meRukWts0pYIoQdyDeDHjp3HSjcZQ6PqPR7F+ecn7WpL8V1Q==",5572:"sha512-JmFQMbLYSZyJ8JXrnf/hAJeleySILaI/oc+HToNIUiUl7ZdDprIRHGUR0sOVxJchuX8jVUgyt/C3866mTfp6VA==",5595:"sha512-wxd90pfVD3Dt1HTqoERiGCzbjvMG0ow7i+yZpAreKwobs6ND+dag5mrBPZ1bjna90N0QirZnBD66ktZaqBWkQQ==",5641:"sha512-3pLwRfgHV1piigJ3PBMnbPlYcpy43JLv9NXbh8HE2CU0kWDZuNFT4qgWRQz+24wi0KSXZVoXW7S3yEc90CB4Wg==",5678:"sha512-zDFInxyXHf81QbevOprj075MnZ0ozLyemEwtA8WBFq/biaGhYDo7cHlRAgNQEwmHDeQGcxBZtsdlXKtoNtixIg==",5696:"sha512-g5z8cyD3p55D/Z7L/NlPNyPMEMBwk3VDPNWHdyHDVD1/Jtbt4lzhbDXXaLb5DDulYz3FAGEYr3diB8ARZH7TfA==",5729:"sha512-+06qaKhR6+33SgXOXpcBcNrOGl6iwWLf5422UW4lCipxLdb+3lMKP+3eQUhGLiDV2TS29Dju0vP9wcwN93uk9g==",5744:"sha512-w0sOI9sdy20ewTLSpfrji/HmajFyDvR5/loi0akSLUQOaMzazNswtjbGl5oTVSL/0b9dHHaLui6h9lCZIv0ZZg==",5746:"sha512-vZxmDt/jlmkCm91rbKwai9uUAuZFsq/jrbC2kOqUz8ldZn7yL6B2DS803OaSwrm/rG2y+qF2tL6fTYUAGvc7gA==",5778:"sha512-4bYP87hQ72ljE4aqJHlDLtvgYZCqarLESMz/VvtgTCnJp5h4JAXo21km84pBhWJkCM9ZD14pX79m0og2Ug2QkQ==",5884:"sha512-vq9idq1LEl0SaZNV3iRA7WbbJSbXpUIz6IpTERcMjXO7z6CQut2uy/vEgBqAVAaKAmbqWjD1O6XXrMEx29b/Dw==",6055:"sha512-9Lj1nZTmUIDsjOf8qyE6VoTv5Zdynq107EutaxGSQ9Dj6jmFudpdGA8CPafPBgqGSqvLqcquAJBwUNPhppPmaw==",6132:"sha512-wl3CgLeSOzpF9o4KUBj05ImNXvT8ojpF9yFgWVQuOkyijAO+gxm8zXsR/2WXeUNr1h8ADKeIJ3MHFFWom+5A7A==",6327:"sha512-ixaU65rHoV3NA50O05p3EI96qMPDuHlsDErN9pWKRzMRUOIteYYUwGcL6pK9mybDrTK2LDmEIPzddZhkemEQoQ==",6432:"sha512-oTFKsjkrK+Uf65jYlHwAbJ9z25B/cEryoLsxGxcNznqWICh6loBiEdnfDhhxvAF2e1WyoM14c9Q2BY+75HEw3g==",6521:"sha512-QBYup4VFt4JM672UPdqZSZu3u+gCGSUOjn/sMbwWXYiT5orojMUEu1xTprea23tL652Lhjd8jhij0LObStZeOQ==",6536:"sha512-HeZLTJl5VjYfLSV2VXw5Zt189ke2auomh0HHfMt3WSOcKtqxGf9PRLd25UiLwOaNv0EeXmOSk/caMUv1PVF65Q==",6558:"sha512-Cfi1uU2KbWMRH2qCvSZCZRbbDT1lG3BuYHDHR+mD65tqJeFvpXgP5bR84rk1gVOsxiDh3HTb2upXSIYHQYxBUg==",6660:"sha512-bVodYwTLKPQiuIUV1kGl0tPBrqbLfHHJZ/fBEBsUEKU/eQNh0Jw0UE4hcp/kuq6lZKRimLYeIvlOfu8pJH87qA==",6706:"sha512-ILYD9/Mga4bfgbX/fn/eMRmEndTpFD0EsfHtSm3xDanBkH7pj4p1W7AG+YugzM6tDpUz3q0AA87cSWBUBIvqMw==",6789:"sha512-nowQUiLe4zz/RzKSBQt/di0y/NgGZgZ3ZzVu6MsXkwpBVwcG3RSfQEAzC/2oRzD5/IDkxkXmH6gD9QPiVyLf4w==",6797:"sha512-+CPLme0oI0mba0nJtGjvyPeVUQHHLMkKOLFr9Ug6OlHo7HcMYe7b+9W8QgKZb6okrKUvArbiAhmRUeQV0f9GqQ==",6816:"sha512-7VfXQmuyC5uEpYYG0IFVlBTFYUAeQ9ei3Y5kJ0pMEClWqZSt/DAWusGviB+8ekB3RQfFw4xNQhi9S6Ez46fuUg==",6893:"sha512-zsQooZI/j6PFXCr62UnH0c/V4b3ApWW6sA+uCTboDDfyhomvtgtTDEOD+J6Li/t3wvava+hr4Q6HkxFVrv/PNw==",6903:"sha512-leuAdDshdNZS/rSrhsNCH7bg6mNbj6+Ljf+quxoRWX3SPSKZGs+w+uhhmSjWhXK4Mwapgb13hblvqpncAcnmGQ==",6906:"sha512-aLYpj7tpF8MJ6YhukYwAYvH0Jy008vusCRqkB0dAs9AZLlsJJ9SzuyFHHSJGTJyRLgeri8imqqI7/F6aBTRe0w==",6919:"sha512-RF4Q87RZ5U6ikLYV9NZBEusJDrl8qS4daV6fsEE8C1pQGcSxMCmpIoBLDKJ1qloUZd/6hBEKHWDfN9zdHDJ6Vg==",6950:"sha512-ms6bQ8ePndTnlmqf3NwOoat0Aw7w2L5sFRFX2rDEPKqDyN2foH4evmRc+4atjdnOYlSVTbyW3cX/HwKI2/sQ/w==",7006:"sha512-EztJE4N5VDjUjEAGfs5gJlXahOEfzUPDQOd85LDc6GK115KpvzJyKNVBHPnAH/qmHMNl+kk66YO9Rti0ztUK0w==",7109:"sha512-1mHm+wP6U6oBJy2vVlFfL8Lvs5SqI3i6xsU1Ha26gtS0DSBnE9fy+mxWtQ1aGYtkj27OEDetzZsOxMUw/VPi5A==",7125:"sha512-ZwzyQ2Xks1lSDpLurer31yLjBwG+bZnBaPLQFJTAUfR3vYlDJbgoz7/0uIU7W/rFH9K+2JjTTYM23zdnkgRY+w==",7155:"sha512-rbJ2UA7fI8L80BRWDN5jEriEpMSCevpuukRWgNHKrmLTg3UKnAKeDJldfyvLI+JzanY0+sLjlxoWQzswpEVjFg==",7201:"sha512-N1YGu35uaq7DRkIpuhKQQt24frpP8WhqH3XGku2hjahHefXkbpEmYv6etm+AxtBPgmVinTstavEJMiYw5S9ayw==",7290:"sha512-mqV2RZdVB3G9WaKDJ7AFW8amtmkI9qm/KecV0MW+Q8QaYjHdoTNsE3uNpo4UT0PcF+Fqx0egXCoQVp9L7Iy1Qw==",7296:"sha512-Gkq2PnaAEfFGUAVFbYsuU35iA+DM36pniBdjSCuWdfJjFqC1X1Xg0ZtM8UpVOBNm7b/DwOXkbENTTntmGvRg8A==",7311:"sha512-LePqOev5GsBkJxu+2GlGD4L4gShctSdrTh//c4GP6z6W2meHNsSc4fcFJXd3XTEg73Qgw9s1FOM7PtS1R0z5Og==",7314:"sha512-wBO3++8gKD9oxOBrCAjeRFv/k4uAbX+Kl322ovsDJmMOyYaU15cwWo1DOEsfAim8poHLq8WMEBZCvelq98YSqA==",7356:"sha512-DKzR1kb+OsWXbvlHuCBMrOJ1Sh9Blw10oBDjbUXf85q6lLSC0VQZsicpxAwABW58VFE+OHcm0yM7ok/5tAnLug==",7374:"sha512-2/H+K/1Z8J9ncv1ETWbD1234XP3ABGg4rt65zOA5sDTARyXcL2EMKHuOI3q1dl6H4DNwh48XQVwyS4zvBVJCXA==",7407:"sha512-94NYiGq7SacF8HAfZsQ/CyxvnHoo8Fx4LHDbuey+kuBYhKCByFICfdh6bCM95wPGEh93mLtqOQS1xu5Vk6hVgA==",7442:"sha512-XEe5A7G/k+t731651Uv7GWQYfyHauLR03jRqV5Q2Udqt6iohZMOJWz5SApIcYZgJ/Z+vtBGTusJEchvkzDgCwg==",7451:"sha512-F1VPeG46PimI7dpl1c/q/FAZN2nQYAnv3nu3227UCTktHwuQU3NQI/Ljaq9jPfMd7YbnP9vPDtoT7836ZES89w==",7458:"sha512-i4Xk5DQGd+dbq2DOawU0qQYIvorluipzOzbHvrDuMLUMcrDA8rIzNI79A8j/IhjnPBaqomesIYn1Dw782+oNzg==",7523:"sha512-XDFpkHonPnzE5vxqV7d5DHh4kT/o2hzpnbZApkJQLlz4KUdZXGpXOf5VU3Y14kKPCfuLdQWfLcYzgTPdclxiqQ==",7569:"sha512-bgyJ7pu1I2GQL0D/RFvRenpiQmvPMAcA0oONGCd0jmSK3lS8dM5lDk3LrnA+BTXyIEaWEY0Hz3drbXkuHh5KmA==",7624:"sha512-ohrq+SZxguAfdjmrEWY3AQVUyiHZksoSxJrhYe4jaLIVqw5taHMPqNl7K6oGr0uNGgwQmqZPdvddB+pD/9smdA==",7687:"sha512-ESy16XSYlQ6IDq4baXTPwKWNfj6NWe17HZzGxMJICZXI52jG2n3rq65wF1E6/ebSGo6yFVDL/ULhW8VY43YeTw==",7689:"sha512-Did+mlQME+02FLq23boc9ivrh2P3ZmroCiEwL2PyhzYRTaYo7vmyf8FjpDEwIOTiU7bgGHTei+epK3+XuCyL0w==",7717:"sha512-cRWPgyyuLLNHHO3Icw6pWR4uCBCMbGOY9Lg9W0TjZq+HzZP691wyU9oiBnBeuWHeAX2QMQT9mwt2LnlEDKjEIQ==",7799:"sha512-U0+HRzxtSvfZmjYFtOEcqIKdfzKt3bXiZuVFCkxIDq18iL+5ae3B0hm3oySmSSTQ2AxvuVy2EMnZLmYYOC1iqg==",7985:"sha512-vg+p/e4q7P0Pc5Bn91tToqyr+0g4AVkLZfUCJ/sAZ8G2AN4Z74QIe9u+Hy8pAcY3U7d+QkeAR6BUFh3W89/nkQ==",8006:"sha512-PD04/PyQ8zqOJ06UYYCEjc8W54Qu67ED0lnWO7X+0TTJWK5Jt2saRv3JsGXbXe9qpIy99eCeIsD/LIb/plgvDA==",8018:"sha512-B8HnC2ajhdNnVHf3Ttb5pozmtUjXNksjlYrSlbB/4IN1+bdvnNbCb5whOf+nbP+Os+egU37yiqcWayYO0MxZPg==",8022:"sha512-HJ10fmBdpA6gZJiVdmt5gsIsGMUSbpMi1ocoCkj8MBveE59XipIU5FbTazfFvjl5GnZAlNLvgt0o7/DCrmjzBQ==",8062:"sha512-KbdWk2pxA8MoPkRSDAeQgaFvdYHNcCrxa5J4mfXDA6v10pa8G4RHN/zKenb65AWZGA2CxF+v0/QAnZmfLectUg==",8090:"sha512-y2+5WWomXk/SFMAucMcLSBuwjg77yFjYWdNWq5xKmRF2P1M5nOcXggGjczkw8SJ1FI6++q+T0vtqat9mQdy0LQ==",8104:"sha512-hereT9OYTaOiANbwppwJelT7CEwQSqE4yl8jwFzIdQi802uBZKTMHycurrEKgLCs3pzU8meqY5chBADlD8W0aw==",8141:"sha512-9LJiwOTKcbIfLcUWvUSiH7hJz22ewIWEqE0kj1EnvjvYDMgK3uAtRmueFlZtkxAwJAnj41vOccFGeKr61qfrow==",8145:"sha512-zvbuHrd7c67v4vL3MEkppHJ/Z24eicpK7DvoETwu4VoJYjrU7k7Yz3XvbrgXgFtcms9u3D8ye5LqwrszsRkbhA==",8226:"sha512-pbOy72L+5JYBgtyr07IWmB24VolLgRjp4Cr2cjTOo7JrJy8LzEJm/Kui+7UFZd5VtAQwAARh8phw2ukKkyYIag==",8306:"sha512-uWDs9KO+xet1boH+LBp5ak9n00cncCIyyjmoFSI5znCWIiteRFFJHTcJqf5TnSTLX9QpUHutJdSlgcQHGAIGhA==",8370:"sha512-yeXfPawffFiT/aUbxSD32192VTy+Ggh9dPU55ymT4zZDdcB+7lr/mXsS694MEL+HVH8tFj1/PgpZMe2mro+PMQ==",8377:"sha512-AKR7IWmhoiceoOCLqz/GBIZ6wdZZ88aDDmPM3zOvImsSpmASWOW36cSWIRdAE3O2TwY5VKb2/xhAIXwt8gwxnw==",8393:"sha512-TVH6wqOyjpljTVAlyQOCGJVdpG4lBQrScbeN4k3rW8nwgiAds/MGQj6ewtQSwYRViqdxwmPgMY7Xw6ZHSF0TSQ==",8438:"sha512-X10yRhFHFGfnj3Nk7qLwMQGLcUpljxEWDCKWCxu9mn+rwQz8itogKTpuaiCv0vqQ/Jfx92kjoPr/FM7rP5Wt3A==",8519:"sha512-N3tkFS0oW/8ZuB5j4tRSmHDPHwYWRfel2omGryPfeFtPqnBdVQ5geTVfSF5UoHQe2GS7L+RXV55/8PxMK7drJw==",8585:"sha512-j2mSF8NeXBlHZDU0qHya434va2cKt3UqchPRKCBre5RlsrgbnHYt3VfRAt2VkfxcZNt5THWdfvXdtyU4+RjApw==",8644:"sha512-RnI+VcCkji53P/Cwlf+w8/POIcaH7P4ob+rT0ATcu5G0rVjGAiCMhdXaPdLa0ApNVS6dZH3Li4vPkygJglPrPA==",8704:"sha512-EFZbOR+tfp0itYBCrmbrF9u6sDiKJZDEttsVn0UxmveRDA4ErPY91EmO7CRYhXthy/+90jxmxXClTAnt4iXu/w==",8716:"sha512-DlqQYn7nnQovOMhUl3TikS9JztrZCMgUoR0rhpZcztMxY4/tyws2n+uiu6C9sLNxbZOuupc5/IgoKyI+6YpIyQ==",8735:"sha512-54KamfXrG4SCRHn5wUY1Oy5Qy1I1RftOwFbolcbarYGvUNJowfs51xAweyETylxr1P9J43yjAy4N9pH4j/t9dQ==",8737:"sha512-3GmkxxJ/IQEBWJEWPpWbfwTZimFFM08QyNNHgGfCO88WyZ9Bh54RfkocHejARdYoqxp0ZWsLVEDECShQgVyVtw==",8746:"sha512-i04HISM/YtodqCX5i33A/zpg91a8ljVjHk4Jvyegs85utylciz7qCtHjaLCNKmjJr9ZaWK5BEuQ1/S9FdouZTA==",8747:"sha512-byo9caguuJmxVIh/Y8BzNw072qAzN6FCtAo02eIDRfkGqMNDbKhdyLJdeoIAImOFUW+6yw9idhA3fYGnTVIPug==",8759:"sha512-ZM0S+jYdgPIKHkBXbXLXMq1pJwHDpOacrhXMvJarMQ0JEFRFM05F/NuYKjA11wKfRCFFOCuPhl0CuHvNIf0fNw==",8798:"sha512-H+QzfNz51tfeLMGG97YLH7D1cPwK33dFH1GEAR4jczcheve3prbsi/tO1bgzSt8eRHVtEEIo8EUDmW3vMuzDeA==",8801:"sha512-A1HeEX/h1A4YhTc7FKMBk4Rj/kRb02f4lvfAc+SqlOfjfkL8srusepQyS1dFX1GDJ+7hWqkJWg82jEsBM6yEhw==",8841:"sha512-a5MlnKtS8fiFP+e/B8tRTpjQYMca+drHoA7aTinAMDb6qLOqSGmbzfypWARwPfGbGL4ZQPN5nSWkxdqDRflWUg==",8976:"sha512-bmJwA7FeTRYHH/0hRjDJV7Cwz9dCjKD7cG+9bugCfphOuIWOF7xf9EnhSX3qfYMlMPyEmTdGjbLy4YEReKfDqQ==",9048:"sha512-IqE771iRC1eHFIIAY/m2jXv9POMCqlEPFdK7KBFIz0D3tJCQyWDgcfagnk7c8lOVykM/AdTDVBQbK3xnNG9Ykw==",9076:"sha512-sOeU5njjpMFfcIVsCJnrCyxtPuQUbdJYV9NuLr3RSrSFaVgwtxhHyZt8yy0985jxK97oS+AUdKpzyVnF8zM/cQ==",9079:"sha512-fHFtt7hUB45DTqtWB8JrvpT0yinB4F24mPE2Vwx65ewqNfxZp9EAa0ARfH2XM8fsBA9KORoP0yByOYIV8hDlEw==",9081:"sha512-Jkgefd0BTaoLHQdbM1ovT+gwJlZPIBW4PrvB6PlrY5ynP7nWclpZu+xs0sLrhpyKJFBbkTGlp/+pe9rYONZlUw==",9088:"sha512-F/s8Aqk6G66DUAsv5GIO5fxpnAByg02jKctTdZMkMr2uWaIrD634LUWtmzMmydxnzsh6StbULrjHevd9A3Fu8w==",9140:"sha512-xZZTZdxUKMzqigzPIZzXW47VYU6DzLx0f1X+II0jcgHSbDA4VtsTXrP0pi8DbxZOycPtpngVjKTnDFGer2Jxhw==",9159:"sha512-6QASDfXT4OEFKQSFRQztYN1EpAXJjlDy5lKt+8v0RUjjr0XldsY+qJNKcS2dyUGMPAHxCwflEDWXW/1bhge+5w==",9205:"sha512-Kb4eHjCCcVZ+R947GpeMCOzHaYhYnTCGD5rrqIkXIqVUEUSnM/UGxB0VAlT/rvh4AunktUnPjZpEvgqdyVCbug==",9249:"sha512-5oLv22mToWfVnso4DtbNuTh0jb3o0aMQZtEeQ/mg+yUfTX3L5YGYhDG2i8Heu5Czlt4Hrf2qigf/TcKVE62ohA==",9365:"sha512-gt/REOIw1sO66xwts9SykQSN7iCXenG4kpUVtJZKyCLlmIgszonbY/VVvEG9aFZvAWSD1bMKdP6xmMiUcr2ndg==",9420:"sha512-Q3AS389G8CEJGIFMU5ZE0h5NUJRyi7GT4NrXCIrxWrixJiDfZK5uG/BCmLyUKQsoqB4szp7ADjRd1I1VvyGzSg==",9450:"sha512-X90LsRiLwoTB03ew+j0Rh6f1CqpX3GiJpZTkRuvl1kMi7kMi4prlSPLXnCnDS17Mq66RQTaSjVVpjDt4lD8x5Q==",9490:"sha512-4ye+mI53bAEPO2ai358PGJpoBkd5Oqco8TzVeyJ1U+IUHE9HE3wtTCG8ZT4XtHEmCT9rqkCDGSQDYQjE1rxBuQ==",9505:"sha512-Vs3tWHxve7WLi7018YVkP0Uw8wqL932vGCx25GcHDdN2go9IygzIW/JB8zicV4EMksJ8Q+csyYcFmI4HQMkltw==",9522:"sha512-I4nQ+jw8GGkeExlsx78v8A1hrrxXP6T8vRzLtgbZ25TUKNGqmc0MvY3pRZhfKFF7QHXrUki1V/beknPPoRQheA==",9526:"sha512-5MvwxRXrgGz/fUKcBGqtQi3q+/ZvlLk9fqNncBlLv1GbynaIjUJ3XwNWuTZwkDSmsG+CubTIEnW5HVvQVfatNQ==",9529:"sha512-RHsSymeGzGqGGNfU/5xCkd5sU3uysgdKyeF2D21vS0+fGKNUPbSQNmjK6JBus7iZEZEa2qp+grVBQZ0Sl8i0Xw==",9539:"sha512-lTalGcAwdh6E4dMZxODUx5aXCB5qnzptpVH1Vu7reGZlLFDD1Tqua75mYsERLP4aksHDChYPFf97Ws+0MvzuRQ==",9637:"sha512-3Xav/wIlwcQIiNcY/b306GRpDbv1Ky1UGjCEUfZ/dgdhspJb581RkeCGUmzRhY4/gdTFz7+kvBXdBTGqGddI8Q==",9714:"sha512-SxW9bYXtN6S+NVf/ItzVgH0q3+uA/SWNSU4oCJ4DIiBiBQwAlVVzQojWckVski8S7JpSOkDLKkwSWKuUrk2myA==",9767:"sha512-Ra+ZSw6oHdN9Rz3uPEbalXb4Ak19/+UM6iul4bXVoi6aNqw5ytCFcSx3+URvEsWEibEaZufCgDEmuQl4ELXUrw==",9785:"sha512-hmOFT9OddwgVkodTovzElbrgtxOYcqfAmh9oa44rOxNp+DuKZAhyponfbTTzTNEZ+oeTk13pBkH8+SaQ807j5A==",9832:"sha512-PDIFo5qimXNvwAN3cScXHAH4MhyJ/TqEDejTOSXp6zTDhslNWY8QF5ROQmt1N9N50VuJeGMNQeYzEY4KJ46dKg==",9897:"sha512-/BRNa2NXFPj0B89p8lc2ZjiyRoc78wv3yoi6rCV2XAsYhEKPOqattRosI7GtdHGr/qrRp/Xqd4JS/R/bEweCww==",9917:"sha512-CQJ2hf3RLqJ7SKhYO1cVubwD3h4RLlYbp06KG2r6gZIkQrwKAaKeYeAP1FihOqLQeTgMvz+8SCKadRK5K0ZTlw==",9963:"sha512-5wdxhRCgQC8txr19KpUv9Q4Tcxv92V2HNXo4AImld12Fn40EwOzrUVghMfJMsUnWUu6kLLBP1fe/C1Rn6LpatA==",9979:"sha512-vKXJnJQI7UzjdHIBj2JDkvgji532JoFbWdqfGrOcVOqw9rye4iNnKaZqEoav4FtZ4IWJqeRUb3evRTOLTlFRIQ=="},d=e=>new Promise(((H,a)=>{var c=r.miniCssF(e),S=r.p+c;if(((e,H)=>{for(var a=document.getElementsByTagName("link"),c=0;c<a.length;c++){var S=(f=a[c]).getAttribute("data-href")||f.getAttribute("href");if("stylesheet"===f.rel&&(S===e||S===H))return f}var d=document.getElementsByTagName("style");for(c=0;c<d.length;c++){var f;if((S=(f=d[c]).getAttribute("data-href"))===e||S===H)return f}})(c,S))return H();((e,H,a,c)=>{var S=document.createElement("link");S.rel="stylesheet",S.type="text/css",S.onerror=S.onload=d=>{if(S.onerror=S.onload=null,"load"===d.type)a();else{var f=d&&("load"===d.type?"missing":d.type),b=d&&d.target&&d.target.href||H,t=new Error("Loading CSS chunk "+e+" failed.\n("+b+")");t.code="CSS_CHUNK_LOAD_FAILED",t.type=f,t.request=b,S.parentNode.removeChild(S),c(t)}},S.href=H,0!==S.href.indexOf(window.location.origin+"/")&&(S.crossOrigin="anonymous"),document.head.appendChild(S)})(e,S,H,a)})),f={4556:0},r.f.miniCss=(e,H)=>{f[e]?H.push(f[e]):0!==f[e]&&{2353:1,2808:1,3542:1,4098:1,8062:1,8747:1}[e]&&H.push(f[e]=d(e).then((()=>{f[e]=0}),(H=>{throw delete f[e],H})))},(()=>{r.b=document.baseURI||self.location.href;var e={4556:0};r.f.j=(H,a)=>{var c=r.o(e,H)?e[H]:void 0;if(0!==c)if(c)a.push(c[2]);else if(4556!=H){var S=new Promise(((a,S)=>c=e[H]=[a,S]));a.push(c[2]=S);var d=r.p+r.u(H),f=new Error;r.l(d,(a=>{if(r.o(e,H)&&(0!==(c=e[H])&&(e[H]=void 0),c)){var S=a&&("load"===a.type?"missing":a.type),d=a&&a.target&&a.target.src;f.message="Loading chunk "+H+" failed.\n("+S+": "+d+")",f.name="ChunkLoadError",f.type=S,f.request=d,c[1](f)}}),"chunk-"+H,H)}else e[H]=0},r.O.j=H=>0===e[H];var H=(H,a)=>{var c,S,d=a[0],f=a[1],b=a[2],t=0;if(d.some((H=>0!==e[H]))){for(c in f)r.o(f,c)&&(r.m[c]=f[c]);if(b)var o=b(r)}for(H&&H(a);t<d.length;t++)S=d[t],r.o(e,S)&&e[S]&&e[S][0](),e[S]=0;return r.O(o)},a=self.webpackChunkwsp_assets=self.webpackChunkwsp_assets||[];a.forEach(H.bind(null,0)),a.push=H.bind(null,a.push.bind(a))})(),r.nc=void 0})();</script>


<script>window.Promise || document.write('<script src="//assets.s-bol.com/nl/static/assets/js/promise.7422877a152ad61057ad.js"><\/script>');</script>


<script>(self.webpackChunkwsp_assets=self.webpackChunkwsp_assets||[]).push([[8792],{24610:()=>{!function(n,e,t){-1!==n.location.hostname.indexOf(".bol.com")&&(e.domain="bol.com");var o=t.className.replace("no-js","enhanced js").split(" ");"ontouchstart"in n||navigator.MaxTouchPoints||navigator.msMaxTouchPoints||n.DocumentTouch&&e instanceof DocumentTouch?o.push("touch_supported"):o.push("no-touch"),0!==["iPhone","iPad","iPod"].filter((function(n){return-1!==navigator.userAgent.indexOf(n)})).length&&o.push("ios"),t.className=o.join(" ")}(window,document,document.documentElement)},40045:(n,e,t)=>{t(24610),t(73728)},73728:()=>{!function(){function n(e,t){if(this instanceof n)throw"px.stub() cannot be instantiated";var o={};for(var r in t)t.hasOwnProperty(r)&&(o[r]=t[r]);n.queue.push([e,o])}n.queue=[],window.px=n}()},47113:(n,e,t)=>{"use strict";t(40045);var o,r="load",i={},u=!0;function s(n){var e=function(n){return n.join("+")}(n);if(i.hasOwnProperty(e))return function(){return i[e]};var t=n.slice(),r=t.map((function(n){var e=o.retrieveChunkForModuleId(n);return e||function(n){return new Promise((function(e,t){var o=document.createElement("script");o.async=!0,n.match(/^(https:)?\/\/((static[^\/]+\.|[^\/]+\.s-)bol\.com)\//)&&(o.crossOrigin="anonymous"),o.src=n,o.onerror=o.onload=function(){return o.onerror=o.onload="",e()},document.head.insertBefore(o,document.getElementsByTagName("script")[0])}))}(n)})),u=new Promise((function(n){Promise.all(r).then((function(e){a((function(){e.forEach((function(n,e){n&&n.define&&n.define(t[e])})),n()}))}))}));return i[e]=u,function(){return u}}function c(n,e){return function(){return new Promise((function(e){a((function(){n(),e()}))}))}}function a(n){window.setTimeout(n,0)}var f={enqueue:function(n,e){var t=0===n.length,o=s,r=c;Array.isArray(e)||"string"==typeof e?(e=Array.isArray(e)?e:[e],n.push(o(e,n))):"function"==typeof e&&n.push(r(e,n)),t&&f.dequeue(n)},dequeue:function(n){n.length>0&&n[0]().then((function(){f.ack(n),f.dequeue(n)}))},ack:function(n){n.shift()}};window.addEventListener(r,(function(){u=!1}),{capture:!1,once:!0,passive:!0});var l=function(n,e,t){return new Promise((function(o){if(n()){e.addEventListener(t,(function n(){o(),e.removeEventListener(t,n,!0)}),!0)}else o()}))},d=l((function(){return!("interactive"===document.readyState||"complete"===document.readyState)}),document,"DOMContentLoaded"),h=l((function(){return u}),window,r);const m={dom:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(9140),t.e(8737)]).then(t.bind(t,24215))},modulesInChunk:["dom","tiles"]},catalog:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(7624)]).then(t.bind(t,54631))},modulesInChunk:["offcanvas","in-view","flex-tooltip","modal-window","basket"]},account:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(8644)]).then(t.bind(t,95957))},modulesInChunk:["account","account-questions-answers","reader-preview","align-columns","scroll-to-element","review","soundplayer","multi-bundle","show-more"]},"pdf-viewer":{fn:function(){return Promise.all([t.e(748),t.e(2246)]).then(t.bind(t,68403))}},"campaign-page":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(51)]).then(t.bind(t,54958))}},giftcard:{fn:function(){return Promise.all([t.e(9522),t.e(7451),t.e(8759),t.e(251)]).then(t.bind(t,1570))},modulesInChunk:["giftcard","giftcard-app","claimable","study-book-selector"]},"dmp-gpt-loaders":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(2621)]).then(t.bind(t,93080))},modulesInChunk:["mms"]},"analytics-loader":{fn:function(){return Promise.all([t.e(800),t.e(8841)]).then(t.bind(t,40352))},modulesInChunk:["analytics-loader"]},"brand-party-seller":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(8519)]).then(t.bind(t,26450))},modulesInChunk:["seller-popup","brand-page"]},"expert-explore-offers":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(3183)]).then(t.bind(t,77673))},modulesInChunk:["expertpages","offer-compare","books"]},"address-suggestions":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(6797)]).then(t.bind(t,94432))}},notifications:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(6893)]).then(t.bind(t,55352))}},"filmstrip-quickview":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(752)]).then(t.bind(t,5291))}},checkout2018:{fn:function(){return Promise.all([t.e(9522),t.e(5458),t.e(5255),t.e(8018),t.e(7458),t.e(3448),t.e(9365),t.e(7356),t.e(9420),t.e(8377),t.e(3022)]).then(t.bind(t,93367))}},measurement:{fn:function(){return t.e(8022).then(t.bind(t,29539))}},slots:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(9140),t.e(9714)]).then(t.bind(t,30449))}},returns:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(4636)]).then(t.bind(t,56951))}},"file-upload":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(2029)]).then(t.bind(t,60084))}},forms:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(6132)]).then(t.bind(t,10439))}},"textarea-counter":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(8438)]).then(t.bind(t,72867))}},"platform-details":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(7799)]).then(t.bind(t,87132))}},"countdown-date":{fn:function(){return t.e(5641).then(t.bind(t,6750))}},"product-compare-table":{fn:function(){return Promise.all([t.e(9522),t.e(7451),t.e(5458),t.e(5255),t.e(8018),t.e(2543),t.e(6558),t.e(1330),t.e(6789)]).then(t.bind(t,28152))}},"conversational-ui":{fn:function(){return Promise.all([t.e(9522),t.e(3928),t.e(759),t.e(396)]).then(t.bind(t,99995))}},inputmask:{fn:function(){return Promise.all([t.e(7451),t.e(4749)]).then(t.bind(t,94968))}},jquery:{fn:function(){return t.e(7451).then(t.bind(t,144))}},underscore:{fn:function(){return t.e(7451).then(t.bind(t,91910))}},howler:{fn:function(){return t.e(4832).then(t.bind(t,60351))}},handlebars:{fn:function(){return t.e(7451).then(t.bind(t,31110))},modulesInChunk:["handlebars","hbs4"]},swipe:{fn:function(){return t.e(235).then(t.bind(t,1988))}},backbone:{fn:function(){return t.e(7451).then(t.bind(t,70883))}},px:{fn:function(){return t.e(2623).then(t.bind(t,29435))}},webshop:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(9159)]).then(t.bind(t,79303))}},banner:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(2493)]).then(t.bind(t,32593))}},history:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(8141)]).then(t.bind(t,80233))}},authenticate:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(5156)]).then(t.bind(t,91194))}},recapthca:{fn:function(){return t.e(596).then(t.bind(t,88826))}},"customer-profile":{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(427)]).then(t.bind(t,90075))}},menus:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(3745)]).then(t.bind(t,2897))}},carousel:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(9897)]).then(t.bind(t,12221))}},toggles:{fn:function(){return Promise.all([t.e(7451),t.e(8759),t.e(1350)]).then(t.bind(t,29320))}}};var p={retrieveChunkForModuleId:function(n){var e=null;return(e=m.hasOwnProperty(n)?n:Object.keys(m).filter((function(e){return m[e].modulesInChunk&&-1!==m[e].modulesInChunk.indexOf(n)}))[0])?m[e].fn():null}};const b={init:function(){return{loader:(n=p,o=n,function(){var n=[],e=[],t=[],o={now:function(e){f.enqueue(n,e)},dom:function(n){e.push(n)},win:function(n){t.push(n)}};return d.then((function(){e.map(o.now),o.dom=o.now})),h.then((function(){t.map(o.now),o.win=o.now})),{now:function(n){return o.now(n),this},dom:function(n){return o.dom(n),this},win:function(n){return o.win(n),this}}})};var n}};t.p=document.getElementById("js-asset-host").getAttribute("href")+"/nl/static/assets/",t.e(1389).then(t.bind(t,1389)).then((function(n){n.startLogger()}));var w,P;"customElements"in window&&"function"==typeof window.CustomEvent&&Element.prototype.matches&&Element.prototype.prepend&&!(!1 in window)&&"function"==typeof Object.assign&&"function"==typeof Symbol&&"fetch"in window||Promise.all([t.e(9365),t.e(1146)]).then(t.bind(t,82321)).then((function(){var n;return null===(n=console)||void 0===n?void 0:n.log("Polyfills loaded...")}));w=b.init(),P=t.e(1075).then(t.bind(t,50554)),w.loader().dom((function(){P.then((function(n){n.initializeAsync()}))})),window.$RNWY=w},41387:n=>{"use strict";n.exports=WSP}},n=>{var e;e=47113,n(n.s=e)}]);</script>


    





  <link rel="asset:css" href="//assets.s-bol.com/nl/static/assets/css/bol-symbols.min.696139c1d6b57be86260.css" />


<script>(self.webpackChunkwsp_assets=self.webpackChunkwsp_assets||[]).push([[6975],{80427:(e,s,t)=>{t(12429)},12429:()=>{window,document},74123:()=>{!function(e,s,t){var n=e.sessionStorage.getItem("async-css")&&JSON.parse(e.sessionStorage.getItem("async-css"))||{hrefs:[]};function r(c,i){var a;a=function(){if(!s.body)return r(c);var a=-1===n.hrefs.indexOf(c)&&function(){var s="test",t=e.sessionStorage;try{return t.setItem(s,"1"),t.removeItem(s),!0}catch(e){return!1}}(),l=s.createElement("link");l.rel="stylesheet",l.href=c,l.media=a?"x-non-blocking":"all",l.onerror=l.onload=function(){l.onerror=l.onload="",a&&(l.media="all",n.hrefs.push(c),o()),function(e){if(e&&e.length){var s=n.cssCacheKeys.indexOf(e);-1!==s&&n.cssCacheKeys.splice(s,1),n.cssCacheKeys.push(e),$RNWY.loader().win("webshop").win((function(){WSP.utils.Cookie.write("cachedAssets",n.cssCacheKeys.slice(-5).join("_"))})),o()}}(i)},t.insertBefore(l,null)},e.setTimeout(a,0)}function o(){e.sessionStorage.setItem("async-css",JSON.stringify(n))}n.cssCacheKeys||(n.cssCacheKeys=[]);for(var c=s.querySelectorAll('link[rel^="asset:css"]'),i=0,a=c.length;i<a;i+=1)r(c[i].getAttribute("href"),c[i].getAttribute("rel").split(":")[2])}(window,document,document.head)},39449:(e,s,t)=>{t(74123)}},e=>{var s=s=>e(e.s=s);s(80427),s(39449)}]);</script>


  
      <link rel="stylesheet" href="//assets.s-bol.com/nl/static/assets/css/listpage.min.5d71b531d93c80bb84a5.css" />
  






    

    <title>bol.com | Zoekresultaten voor &#x27;9780575104181&#x27;</title>
    
  <meta name="apple-itunes-app" content="app-id=904894760, app-argument=https://www.bol.com/nl/nl/l/alle-artikelen/?promo=smartbanner_ios_open">



<meta name="google-play-app" content="app-id=com.bol.shop">
<link rel="alternate" href="android-app://com.bol.shop/https/www.bol.com/nl/nl/l/alle-artikelen/?promo=smartbanner_android_open">


</head>



<body  data-measurementname="webshop_searchpage-container-view_va1ei"  data-bltgc="sOloXhA4zqCekILbwytNBQ"  data-container-name="SEARCH.Search"  data-backend="future">

<div class="main" >
    <a href="#siteSearch" class="skip-link">Ga naar zoeken</a>
<a href="#mainContent" class="skip-link">Ga naar hoofdinhoud</a>


<svg width="0" height="0" class="srt" focusable="false" aria-hidden="true">
    <symbol id="logo" viewBox="0 0 1134 445">
        <path fill="currentColor" d="M602.208 334.405C642.677 334.405 670.541 307.407 670.541 268.835C670.541 230.264 642.677 203.264 602.208 203.264C561.737 203.264 533.873 230.264 533.873 268.835C533.873 307.407 561.737 334.405 602.208 334.405ZM602.208 93.3373C745.403 93.3373 800.136 186.551 800.136 268.835C800.136 358.835 737.129 444.333 603.48 444.333C460.285 444.333 404.279 351.763 404.279 268.835C404.279 181.407 463.467 93.3373 602.208 93.3373Z" />
        <path fill="currentColor" d="M826.429 0.00528719H955.889V438.429H826.429V0.00528719Z"/>
        <path fill="currentColor" d="M1133.86 365.353C1133.86 321.733 1098.5 286.372 1054.88 286.372C1011.26 286.372 975.897 321.733 975.897 365.353C975.897 408.972 1011.26 444.333 1054.88 444.333C1098.5 444.333 1133.86 408.972 1133.86 365.353Z" />
        <path fill="currentColor" d="M129.459 329.649C146.852 332.903 166.191 334.288 179.073 334.288C226.097 334.288 257.017 308.408 257.017 268.717C257.017 229.677 226.741 203.147 182.939 203.147C166.835 203.147 147.496 205.372 129.459 211.228V329.649ZM0 0.00528719H129.459V107.361C154.916 97.7186 181.197 93.2186 208.564 93.2186C310.392 93.2186 384.217 167.789 384.217 270.003C384.217 347.144 331.393 438.429 211.109 438.429H0V0.00528719Z" />
    </symbol>


    <symbol id="tagLine" viewBox="0 0 1020 64">
    <path fill="currentColor" d="M19.26,63.92C8.36,63.92,0,56.06,0,42.48v-.66c0-13.18,8.61-22,19.92-22A16.29,16.29,0,0,1,35.57,30V3.6H28.69V0H39.92V59.83H46.8v3.43H35.57v-9.9A18,18,0,0,1,19.26,63.92Zm16.4-21.77v-.66c0-11.86-6.32-18.16-15.33-18.16C11.15,23.33,4.51,30,4.51,41.66v.65c0,11.79,6.64,18.09,15.16,18.09C28.44,60.4,35.66,54,35.66,42.15Zm13.54.16v-.65c0-12.93,8.2-21.85,19.67-21.85,10.33,0,18.77,6.71,18.77,20.87v1.55H53.71c.16,11.46,5.66,18,15.82,18,7.13,0,12-3,13.44-8.93h4.35c-1.81,8.19-8.61,12.61-17.79,12.61C57.15,63.92,49.2,55.16,49.2,42.31Zm34-3.6c-.49-10.39-6-15.22-14.35-15.22s-14.18,5.89-15,15.22Zm49.85,24.55H128.4L116,24.06H111.1v-3.6h16.32v3.6h-6.81L131,57.94l11.72-37.48h3.69l11.31,37.48,10.41-33.88h-7v-3.6h16.14v3.6h-4.83L160,63.26h-4.84L144.47,27.42Zm49.46,0V59.83h6.31V24.06h-6.56v-3.6h10.9V59.83h6.56v3.43Zm4.91-55.57a3.16,3.16,0,1,1,3.2,3.2A3.13,3.13,0,0,1,187.44,7.69ZM205,63.26V59.83h6.31V24.06h-6.55v-3.6h10.9v8.59c2-4.74,7.46-9.24,15-9.24,8.93,0,15.08,4.82,15.08,16.94V59.83h6.48v3.43H234.74V59.83h6.64V36.5c0-8.92-4.27-12.93-11.72-12.93-7,0-14,4.75-14,13.67V59.83h6.64v3.43Zm68.44,0H256.37V59.83h6.31V3.6H256V0h11V40.84l18.77-17.11h-6.64V20.46h17.14V23.9h-5.82L272,40.51,291.7,59.83h5.9v3.43h-7.87L267,41V59.83h6.4ZM299,42.31v-.65c0-12.93,8.19-21.85,19.67-21.85,10.33,0,18.77,6.71,18.77,20.87v1.55H303.49c.16,11.46,5.65,18,15.82,18,7.13,0,12-3,13.44-8.93h4.34c-1.8,8.19-8.6,12.61-17.78,12.61C306.93,63.92,299,55.16,299,42.31Zm34-3.6c-.49-10.39-6-15.22-14.34-15.22s-14.18,5.89-15,15.22Zm8.45,24.55V59.83h6.31V3.6h-6.64V0h11V59.83h6.48v3.43Zm82.08-39.2-14.76,39.2h-5.41l-15.24-39.2h-5v-3.6h16.31v3.6h-6.56l13.36,35.19,13.12-35.19h-6.73v-3.6h15.91v3.6ZM444,63.92c-8.37,0-14.35-4.09-14.35-12.2,0-10,9.51-13,20-13h8.85V35c0-8.19-3.44-11.46-10.82-11.46-6.39,0-10.82,2.62-11.8,9.08h-4.34c1-9.16,8.36-12.76,16.31-12.76,8.93,0,15,4.09,15,15.14V59.83h6.31v3.43H458.48V56.47C455.86,60.81,451.11,63.92,444,63.92Zm14.5-14.49V41.9h-8.77C440.86,41.9,434,44,434,51.72c0,5.24,3.45,8.68,10.41,8.68C452.09,60.4,458.48,55.82,458.48,49.43Zm15.71,13.83V59.83h6.31V24.06h-6.56v-3.6h10.9v8.59c2-4.74,7.46-9.24,15-9.24,8.94,0,15.09,4.82,15.09,16.94V59.83h6.47v3.43H503.94V59.83h6.64V36.5c0-8.92-4.26-12.93-11.72-12.93-7,0-14,4.75-14,13.67V59.83h6.64v3.43Zm73.15-21v-.65c0-12.85,9.26-21.77,20.82-21.77s20.57,8.67,20.57,21.68v.66c0,12.93-9.18,21.77-20.74,21.77S547.34,54.83,547.34,42.23Zm36.88,0v-.65c0-11.3-7-18.09-16.06-18.09-9.27,0-16.32,7-16.32,18.09v.65c0,11,6.89,18,16.23,18C577.25,60.24,584.22,53.28,584.22,42.23Zm8.48,21V59.83H599V24.06h-6.56v-3.6h10.9v8.59c2-4.74,7.46-9.24,15-9.24,8.94,0,15.08,4.82,15.08,16.94V59.83h6.48v3.43H622.45V59.83h6.64V36.5c0-8.92-4.26-12.93-11.72-12.93-7,0-14,4.75-14,13.67V59.83H610v3.43Zm56.5-6.54v6.54h-3.69V49.76H649c1.07,5.81,5.33,10.72,13.11,10.72,6.89,0,10.58-3.44,10.58-8.51,0-5.48-3.28-7.37-11.31-8.76-10.5-2-14.84-5.07-14.84-12.27,0-6.88,5.82-11.13,13.44-11.13,5.9,0,10,2.29,11.64,6.13V20.46h3.69V31.92h-3.61c-1.14-5.65-5.41-8.59-11.39-8.59s-9.75,2.7-9.75,7.52,3,6.8,11.72,8.43C671.5,41,677,43.38,677,51.72c0,7-5.33,12.28-14.92,12.28C654.53,64,650.43,60.07,649.2,56.72Zm72,7.2c-8.36,0-14.35-4.09-14.35-12.2,0-10,9.51-13,20-13h8.85V35c0-8.19-3.44-11.46-10.82-11.46-6.39,0-10.82,2.62-11.8,9.08H708.7c1-9.16,8.36-12.76,16.31-12.76,8.93,0,15,4.09,15,15.14V59.83h6.31v3.43H735.66V56.47C733,60.81,728.29,63.92,721.16,63.92Zm14.5-14.49V41.9h-8.77c-8.85,0-15.73,2.13-15.73,9.82,0,5.24,3.44,8.68,10.41,8.68C729.27,60.4,735.66,55.82,735.66,49.43Zm15.39,13.83V59.83h6.31V3.6h-6.64V0h11V59.83h6.48v3.43Zm22.55,0V59.83h6.31V3.6h-6.64V0h11V59.83h6.48v3.43Zm19.86-20.95v-.65c0-12.93,8.2-21.85,19.68-21.85,10.32,0,18.77,6.71,18.77,20.87v1.55H798c.17,11.46,5.66,18,15.82,18,7.13,0,12-3,13.44-8.93h4.35c-1.81,8.19-8.61,12.61-17.79,12.61C801.41,63.92,793.46,55.16,793.46,42.31Zm34-3.6c-.49-10.39-6-15.22-14.34-15.22s-14.19,5.89-15,15.22Zm9.28,24.55V59.83h6.32V24.06h-6.56v-3.6h10.9v8.43a14.85,14.85,0,0,1,13.93-9.08c6.4,0,11.23,2.7,13.12,9.9a16.06,16.06,0,0,1,14.92-9.9c8.28,0,14.18,4.5,14.18,16.69V59.83H910v3.43H892.58V59.83h6.64V36.26c0-9.09-3.69-12.69-10.57-12.69-6.56,0-13.12,4.67-13.12,13.59V59.83h6.64v3.43H864.55V59.83h6.64V36.26c0-9.09-3.69-12.69-10.66-12.69-6.55,0-13.11,4.67-13.11,13.59V59.83h6.64v3.43Zm91.83.66c-8.36,0-14.35-4.09-14.35-12.2,0-10,9.51-13,20-13h8.86V35c0-8.19-3.45-11.46-10.82-11.46-6.4,0-10.82,2.62-11.81,9.08h-4.34c1-9.16,8.36-12.76,16.31-12.76,8.94,0,15,4.09,15,15.14V59.83h6.31v3.43H943.1V56.47C940.47,60.81,935.72,63.92,928.59,63.92ZM943.1,49.43V41.9h-8.77c-8.86,0-15.74,2.13-15.74,9.82C918.59,57,922,60.4,929,60.4,936.7,60.4,943.1,55.82,943.1,49.43Zm29.66,14.49c-8.36,0-14.34-4.09-14.34-12.2,0-10,9.51-13,20-13h8.85V35c0-8.19-3.44-11.46-10.82-11.46-6.39,0-10.82,2.62-11.8,9.08h-4.34c1-9.16,8.36-12.76,16.31-12.76,8.93,0,15,4.09,15,15.14V59.83h6.31v3.43H987.27V56.47C984.65,60.81,979.9,63.92,972.76,63.92Zm14.51-14.49V41.9H978.5c-8.85,0-15.74,2.13-15.74,9.82,0,5.24,3.45,8.68,10.41,8.68C980.88,60.4,987.27,55.82,987.27,49.43Zm15.6,13.83V59.83h6.31V3.6h-6.64V0h11V59.83H1020v3.43Z"/>
  </symbol>

  <symbol id="tagLineFr" viewBox="0 0 1065 78">
    <path fill="currentColor" d="M.32,62.68v-3.4H6.56V3.57H0V0H10.85V59.28h6.4v3.4ZM21.79,41.92v-.65c0-12.81,8.1-21.65,19.44-21.65,10.21,0,18.55,6.65,18.55,20.68v1.54H26.25c.16,11.35,5.58,17.84,15.63,17.84,7,0,11.82-2.92,13.28-8.84h4.3C57.67,59,51,63.33,41.88,63.33,29.65,63.33,21.79,54.65,21.79,41.92Zm33.62-3.57c-.49-10.29-5.92-15.08-14.18-15.08s-14,5.84-14.82,15.08Zm386,39.09V74h6.32V23.84h-6.48V20.27h10.77V30.41a18,18,0,0,1,16.12-10.79c10.69,0,19,8.19,19,21.41v.65c0,13.3-8.1,21.65-19,21.65A17.13,17.13,0,0,1,452.06,53V74h7v3.41Zm41.31-35.6v-.65c0-11.67-6.89-17.92-15-17.92-8.26,0-15.79,6.25-15.79,17.92v.65c0,11.52,7,17.84,15.88,17.84C476.85,59.68,482.76,53.19,482.76,41.84Zm10.94,0v-.65c0-12.73,9.15-21.57,20.57-21.57s20.33,8.6,20.33,21.49v.65c0,12.81-9.07,21.57-20.49,21.57S493.7,54.33,493.7,41.84Zm36.45,0v-.65c0-11.19-6.89-17.92-15.88-17.92-9.15,0-16.12,6.89-16.12,17.92v.65c0,10.87,6.81,17.84,16,17.84C523.26,59.68,530.15,52.79,530.15,41.84ZM558.2,63.33c-8.75,0-14.9-4.46-14.9-16.54V23.84h-6.16V20.27h10.53V47c0,8.92,4.21,12.65,11.51,12.65,6.8,0,13.6-4.62,13.6-13.54V23.84h-6.23V20.27h10.53v39h6.8v3.4h-11.1V54.33C570.84,59,565.49,63.33,558.2,63.33Zm29.89-.65v-3.4h6.24V23.84h-6.48V20.27h10.78v9.41c2-5.27,6.88-9.81,16-9.81v4.21c-9.23-.08-16,3.9-16,15.09V59.28h7v3.4Zm60.82-10.22V23.92h-6.24V20.27h6.24V10.62h4.29v9.65h11.66v3.65H653.2V52.14c0,5.11,2.43,7.46,6.8,7.46a17.14,17.14,0,0,0,5.43-1v3.65a16.16,16.16,0,0,1-5.75,1C652.55,63.25,648.91,59.28,648.91,52.46Zm21.71-10.62v-.65c0-12.73,9.15-21.57,20.57-21.57s20.33,8.6,20.33,21.49v.65c0,12.81-9.07,21.57-20.49,21.57S670.62,54.33,670.62,41.84Zm36.45,0v-.65c0-11.19-6.89-17.92-15.88-17.92-9.15,0-16.12,6.89-16.12,17.92v.65c0,10.87,6.81,17.84,16,17.84C700.18,59.68,707.07,52.79,707.07,41.84Zm29.5,21.49c-8.75,0-14.91-4.46-14.91-16.54V23.84h-6.15V20.27H726V47c0,8.92,4.21,12.65,11.5,12.65,6.8,0,13.61-4.62,13.61-13.54V23.84h-6.24V20.27h10.53v39h6.8v3.4H751.15V54.33C749.2,59,743.86,63.33,736.57,63.33Zm33.79-10.87V23.92h-6.24V20.27h6.24V10.62h4.29v9.65h11.67v3.65H774.65V52.14c0,5.11,2.43,7.46,6.81,7.46a17.19,17.19,0,0,0,5.43-1v3.65a16.17,16.17,0,0,1-5.76,1C774,63.25,770.36,59.28,770.36,52.46Zm47.75-10.54v-.65c0-12.81,8.1-21.65,19.44-21.65,10.2,0,18.55,6.65,18.55,20.68v1.54H822.56c.16,11.35,5.59,17.84,15.64,17.84,7,0,11.82-2.92,13.28-8.84h4.29C854,59,847.27,63.33,838.2,63.33,826,63.33,818.11,54.65,818.11,41.92Zm33.61-3.57c-.48-10.29-5.91-15.08-14.17-15.08s-14,5.84-14.83,15.08Zm14.56,14.11V23.92H860V20.27h6.24V10.62h4.29v9.65h11.67v3.65H870.57V52.14c0,5.11,2.43,7.46,6.81,7.46a17.13,17.13,0,0,0,5.42-1v3.65a16.16,16.16,0,0,1-5.75,1C869.92,63.25,866.28,59.28,866.28,52.46Zm48.46,0V23.92H908.5V20.27h6.24V10.62H919v9.65h11.66v3.65H919V52.14c0,5.11,2.43,7.46,6.8,7.46a17.14,17.14,0,0,0,5.43-1v3.65a16.16,16.16,0,0,1-5.75,1C918.38,63.25,914.74,59.28,914.74,52.46Zm21.7-10.62v-.65c0-12.73,9.16-21.57,20.58-21.57s20.33,8.6,20.33,21.49v.65c0,12.81-9.07,21.57-20.49,21.57S936.44,54.33,936.44,41.84Zm36.46,0v-.65C972.9,30,966,23.27,957,23.27c-9.15,0-16.12,6.89-16.12,17.92v.65c0,10.87,6.8,17.84,16,17.84C966,59.68,972.9,52.79,972.9,41.84ZM1003,63.33c-8.75,0-14.9-4.46-14.9-16.54V23.84h-6.16V20.27h10.53V47c0,8.92,4.21,12.65,11.5,12.65,6.81,0,13.61-4.62,13.61-13.54V23.84h-6.24V20.27h10.53v39h6.81v3.4h-11.1V54.33C1015.59,59,1010.24,63.33,1003,63.33Zm34.59-7.14v6.49h-3.64V49.3h3.48c1,5.76,5.26,10.62,13,10.62,6.8,0,10.45-3.4,10.45-8.43,0-5.43-3.24-7.3-11.18-8.68-10.37-1.94-14.66-5-14.66-12.16,0-6.81,5.75-11,13.28-11,5.84,0,9.89,2.27,11.5,6.08V20.27h3.65V31.62h-3.56c-1.14-5.59-5.35-8.51-11.26-8.51s-9.64,2.68-9.64,7.46,3,6.73,11.58,8.35c9.07,1.71,14.5,4.06,14.5,12.33,0,6.89-5.27,12.16-14.74,12.16C1042.81,63.41,1038.76,59.52,1037.54,56.19ZM86.22,62.76V59.35h6.23V23.91H86V20.35H96.75V28.7a14.68,14.68,0,0,1,13.77-9c6.31,0,11.09,2.67,13,9.81a15.87,15.87,0,0,1,14.74-9.81c8.18,0,14,4.46,14,16.54V59.35h6.4v3.41H141.38V59.35h6.56V36c0-9-3.65-12.57-10.45-12.57-6.48,0-13,4.62-13,13.46V59.35h6.56v3.41H113.68V59.35h6.56V36c0-9-3.65-12.57-10.53-12.57-6.48,0-13,4.62-13,13.46V59.35h6.56v3.41Zm92.34.64c-8.26,0-14.17-4-14.17-12.08,0-9.89,9.39-12.89,19.76-12.89h8.75V34.7c0-8.11-3.4-11.35-10.69-11.35-6.32,0-10.69,2.59-11.67,9h-4.29c1-9.08,8.26-12.65,16.12-12.65,8.83,0,14.82,4.05,14.82,15V59.35h6.24v3.41H192.9V56C190.31,60.32,185.61,63.4,178.56,63.4ZM192.9,49.05V41.59h-8.67c-8.74,0-15.55,2.11-15.55,9.73,0,5.19,3.4,8.6,10.29,8.6C186.58,59.92,192.9,55.38,192.9,49.05Zm16.61,16.14H214c1.06,5.43,5.59,9.16,14.5,9.16s15.31-4.13,15.31-14.59V51.08a17.62,17.62,0,0,1-16,10.3c-10.85,0-19.19-8-19.19-20.27v-.57c0-12.08,8.5-20.84,19.68-20.84,8.34,0,13.12,4.62,15.47,9.81V20.35h10.78v3.56h-6.48v36C248,72.41,239.8,78,228.54,78,216.23,78,210.56,72,209.51,65.19Zm34.34-24.33v-.65c0-10.7-6.23-16.86-15.14-16.86-9.08,0-15.64,6.48-15.64,17V41c0,10.55,6.64,16.71,15.07,16.71C236.81,57.73,243.85,51.48,243.85,40.86ZM273.2,63.4c-8.26,0-14.17-4-14.17-12.08,0-9.89,9.39-12.89,19.76-12.89h8.75V34.7c0-8.11-3.4-11.35-10.69-11.35-6.32,0-10.69,2.59-11.67,9h-4.29c1-9.08,8.26-12.65,16.12-12.65,8.83,0,14.82,4.05,14.82,15V59.35h6.24v3.41H287.54V56C285,60.32,280.25,63.4,273.2,63.4Zm14.34-14.35V41.59h-8.67c-8.74,0-15.55,2.11-15.55,9.73,0,5.19,3.4,8.6,10.29,8.6C281.22,59.92,287.54,55.38,287.54,49.05Zm20.79,7.22v6.49h-3.64V49.38h3.48c1.05,5.75,5.27,10.62,13,10.62,6.81,0,10.45-3.41,10.45-8.43,0-5.44-3.24-7.3-11.18-8.68-10.37-2-14.66-5-14.66-12.16,0-6.82,5.75-11,13.28-11,5.84,0,9.89,2.27,11.51,6.08V20.35h3.64V31.7h-3.56c-1.14-5.6-5.35-8.51-11.26-8.51s-9.64,2.67-9.64,7.46,3,6.73,11.58,8.35c9.08,1.7,14.5,4,14.5,12.32,0,6.89-5.26,12.17-14.74,12.17C313.6,63.49,309.55,59.59,308.33,56.27Zm36.21,6.49V59.35h6.24V23.91H344.3V20.35h10.77v39h6.48v3.41ZM349.4,7.7a3.12,3.12,0,1,1,3.16,3.16A3.09,3.09,0,0,1,349.4,7.7Zm18.88,55.06V59.35h6.23V23.91H368V20.35h10.78v8.51c1.94-4.7,7.37-9.16,14.82-9.16,8.83,0,14.9,4.78,14.9,16.78V59.35h6.4v3.41H397.68V59.35h6.56V36.24c0-8.84-4.21-12.81-11.58-12.81-6.89,0-13.85,4.7-13.85,13.54V59.35h6.56v3.41Z"/>
  </symbol>

  <symbol id="logoWithTagline" viewBox="0 0 1053 482">
    <use xlink:href="#logo" x="282" y="71" width="518.74" height="203.28"/>
    <use xlink:href="#tagLine" x="91.42" y="319" width="870.17" height="54.12"/>
  </symbol>

  <symbol id="logoWithTaglineFr" viewBox="0 0 1053 482">
    <use xlink:href="#logo" x="282" y="71" width="518.74" height="203.28"/>
    <use xlink:href="#tagLineFr" x="87.16" y="319" width="878.69" height="64.12"/>
  </symbol>

  <symbol id="logo-select" viewBox="0 0 305 76.51">
    <path fill="currentColor" d="M15.2,69.3v6H0V52.4H16.6c1.1,5.4,5.5,9.8,13,9.8,5.4,0,8.6-3,8.6-7,0-4.4-2.7-6.7-10.6-7.8C8.6,45.1.8,36.9.8,24,.8,11.5,9.9,2.8,23.7,2.8c8.6,0,13.6,3,16.3,6.8V3.8H55V25.3H38.5c-1.1-5.1-4.5-8.4-10.4-8.4-5.1,0-8,2.3-8,6.2,0,3.7,2.5,6,11,7.2,16.8,2.1,26.8,7.6,26.8,23.5,0,12.9-8.6,22.71-25.7,22.71C24.4,76.51,18.3,73.4,15.2,69.3ZM62.7,49.8V49c0-16.8,12.5-27.4,28.2-27.4,14.3,0,26.6,8.2,26.6,27v5H80.8c.5,7.2,4.4,10.6,10.8,10.6,6,0,8.5-2.7,9.2-6.2h16.7c-1.6,11.8-10.6,18.51-26.4,18.51C74.7,76.51,62.7,66.9,62.7,49.8Zm37.5-6.9c-.3-5.9-3.4-9.5-9.3-9.5-5.4,0-9.1,3.6-9.9,9.5Zm21.2,32.41V62.8h4.7V12.9h-4.8V0h22.6V62.8h4.8V75.31ZM152.5,49.8V49c0-16.8,12.5-27.4,28.2-27.4,14.3,0,26.6,8.2,26.6,27v5H170.6c.5,7.2,4.4,10.6,10.8,10.6,6,0,8.5-2.7,9.2-6.2h16.7c-1.6,11.8-10.6,18.51-26.4,18.51C164.5,76.51,152.5,66.9,152.5,49.8ZM190,42.9c-.3-5.9-3.4-9.5-9.3-9.5-5.4,0-9.1,3.6-9.9,9.5ZM251.6,23h13.8V43.2H249.7c-1-4.7-3.7-7.9-9.4-7.9-6.3,0-10,4.6-10,13.4v.9c0,9.1,3.4,13.8,10.3,13.8,4.9,0,8.5-3,9-8.7h16.2c-.6,13.2-9.5,21.81-26.3,21.81-15.6,0-27.5-9-27.5-27v-.9c0-16,9.7-27,24.1-27,9,0,13.6,4.2,15.5,8.8Zm24.6,35.5V35.6h-6V23h6V12.1h18V23h10.7V35.6H294.2V57.4c0,3.8,1.9,5.5,5.6,5.5A13.58,13.58,0,0,0,305,62V74.9a35.93,35.93,0,0,1-10,1.61C282.8,76.51,276.2,70.8,276.2,58.5Z"/>
  </symbol>
  <symbol id="logo-ah" viewBox="0 0 898.58 943.94">
    <path d="M334.59,447.68c-34.67,0-37.68,36.59-37.58,124.7s4.89,122.48,37.47,122.48c44.27,0,114-112.48,114-112.48S379.31,447.68,334.59,447.68Z" transform="translate(-50 -53.15)" style="fill:#00a0e2"/><path d="M937.6,504.33L719.54,98.87c-23-42.63-75.6-58.23-117.61-34.86L202.82,286.14a91.88,91.88,0,0,0-42.37,53.57L53.35,711c-12.83,44.4,12.26,91,56.06,104L738.59,993.69c43.76,13,89.68-12.49,102.53-56.95L945.2,572.65C952,551.28,948.16,524.91,937.6,504.33ZM767.18,769.45H692.65l-0.32-282.79c0-38.84-31.27-38.76-31.71-38.76-25.21,0-71.54,63.1-123.11,133.15V769.56l-75.36.08-0.09-87.26s-49.83,87.26-124.43,87.35c-84.64,0-113.36-58.08-113.57-197.25C223.92,439.89,242.77,369,333.49,368.88c69-.06,128.29,100.31,128.29,100.31l0-63.57,75.72-101.8s-0.32,166,0,166.18C572.92,419,605,368.61,664.91,368.61c66.36,0,102.32,55.42,102.55,102.09C767.46,471.17,767.18,769.45,767.18,769.45Z" transform="translate(-50 -53.15)" style="fill:#00a0e2"/>
  </symbol>
</svg>








<div class="usp-banner" role="complementary" aria-label="Onze voordelen">
    <div class="usp-banner__constrain">
        <a href="/nl/nl/sf/winkelen-zonder-zorgen/" class="usp-banner__usps  px_common_promotion_click">
            <span class="srt">lekker winkelen zonder zorgen</span>
            <ul class="usp-banner__list">
                
                
                    <li class="usp-banner__item" data-test="header-usp-0">
                        <strong>Gratis</strong> verzending vanaf 25,-
                    </li>
                
                    <li class="usp-banner__item" data-test="header-usp-1">
                        Bezorging dezelfde dag, 's avonds of in het weekend*
                    </li>
                
                    <li class="usp-banner__item" data-test="header-usp-2">
                        <strong>Gratis</strong> retourneren
                    </li>
                
            </ul>
        </a>
        <a class="usp-banner__select" data-test="header-fluid-select" data-select-header-banner-click-type="topheader:select:acquisition"  href="/nl/nl/select">
            <strong class="txt-select disable-pointer-events" data-test="select-logo">Select</strong>
<span>Ontdek nu de 4 voordelen</span>


        </a>
    </div>
</div>



<header id="header" class="wsp-header js_main_header"  data-bltgg="sOloXhA4zqCekILbwytNBQ.0" >
  <div class="wsp-header-sections  js_search_suggestion_header_pos">
    <div class="wsp-header-sections__logo">
        <button class="main-menu-btn  js_main_menu_btn" aria-expanded="false" aria-label="Menu">
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-menu bi-2x" data-test="icon-menu" focusable="false"><path fill-rule="evenodd" d="M2.51 4.5A.51.51 0 012 4c0-.276.228-.5.51-.5h8.98A.51.51 0 0112 4c0 .276-.228.5-.51.5H2.51zm0 3A.51.51 0 012 7c0-.276.228-.5.51-.5h8.98A.51.51 0 0112 7c0 .276-.228.5-.51.5H2.51zm0 3A.51.51 0 012 10c0-.276.228-.5.51-.5h8.98a.51.51 0 01.51.5c0 .276-.228.5-.51.5H2.51z"/></svg>
        </button>

        <a href="/nl/nl/" class="omniture_main_logo wsp-main-logo-link" data-test="main-logo" aria-label="Bol.com homepage">
  <svg role="img" class="wsp-main-logo" focusable="false" aria-hidden="true">
    <use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#logo" x="0" y="0"></use>
  </svg>
</a>

    </div>

    <div class="wsp-header-sections__search" data-test="search-bar-container"><wsp-search-box class="wsp-search-form" suggestions-endpoint="https://suggestions.bol.com/v3/suggestions/" bucketing-id="fl509pipaacmjnwat9pqht7tu31joh5w" beacon-id="c20689a4-fdb2-4de8-803f-cea61d9a6825"
 product-suggestions-endpoint="/nl/rnwy/search-suggestions/products"
 history-suggestions-endpoint="/nl/rnwy/search-suggestions/history"
                isMobile="false">
  <form action="/nl/nl/s/"  data-group-name="search"  data-bltgg="sOloXhA4zqCekILbwytNBQ.0_1"  method="get" data-target="top" role="search"
        aria-label="" id="siteSearch" class="js-search-form"
        data-analytics-id="px_search_button_click">
    <wsp-search-box-register>
      <wsp-search-input class="wsp-search-form__input">
        <label for="searchfor" class="srt">Zoeken</label>
        <input id="searchfor"
               name="searchtext"
               type="text"
               value="9780575104181"
               maxlength="75"
               autocorrect="off"
               spellcheck="false"
               autocomplete="off"
               autocapitalize="off"
               placeholder="Waar ben je naar op zoek?"
               class="wsp-search__input js-search-input"
               data-test="search_input_trigger"/>

        

        <wsp-search-voice class="u-hide u-hide@screen-large-up">
          <button type="button" class="wsp-search__btn js-speech-button" aria-label="Gebruik je microfoon om te zoeken" data-test="search_speech">
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-mic bi-lg" data-test="icon-mic" focusable="false"><path fill-rule="evenodd" d="M4.494 3.009v4.016a2 2 0 002 2H7.5a2 2 0 002-2V3.009h1.002V7a3 3 0 01-3 3H6.501a3 3 0 01-3-3V3.009h.993zM4 11h6v3H4v-3zm1 1v1h4v-1H5zM7 1a1.5 1.5 0 011.5 1.5v4a1.5 1.5 0 01-3 0v-4A1.5 1.5 0 017 1zm0 1a.5.5 0 00-.5.5v4a.5.5 0 001 0v-4A.5.5 0 007 2z"/></svg>
          </button>
        </wsp-search-voice>
        <button type="submit" class="wsp-search__btn" aria-label="Zoeken"
                data-analytics-id="px_search_button_click"
                data-test="search-button">
          <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-search-vi bi-lg" data-test="icon-search-vi" focusable="false"><path d="M23 24c-.3 0-.5-.1-.7-.3l-6-6C14.6 19.2 12.4 20 10 20 4.5 20 0 15.5 0 10S4.5 0 10 0s10 4.5 10 10c0 2.4-.8 4.6-2.3 6.3l6 6c.2.2.3.5.3.7s-.1.5-.3.7-.4.3-.7.3zM10 2c-4.4 0-8 3.6-8 8s3.6 8 8 8 8-3.6 8-8-3.6-8-8-8z"/></svg>
        </button>

        <wsp-search-suggestions isMobile="false" class="wsp-search-form__suggestions u-hide">
        </wsp-search-suggestions>

      </wsp-search-input>
    </wsp-search-box-register>
  </form>
</wsp-search-box>
</div>

    <div class="wsp-header-sections__account">
        <div class="o-drop-down account-dropdown  account-dropdown--canvas ">
 
    <button class="account-toggle" id="account-toggle" aria-expanded="false" data-focus="false">
    <div class="o-drop-down__label" for="account-dropdown" data-test="account-link">
      <span>Welkom</span>
      <strong data-no-translate="true">Sebastiaan</strong>
    </div>
    </button>
  
    <div class="o-drop-down__navigation" id="account-menu">
      <div class="offcanvas-container  offcanvas-container--right" data-offcanvas-id="ID2offcanvas-container--right">
        <div class="offcanvas-header">
          <div class="offcanvas-header__title">
            Welkom Sebastiaan
          </div>
          <button class="offcanvas-header__close-btn  js_offcanvas_close" data-analytics-id="px_common_click" data-analytics-tag="header menu:profile icon:close">
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-cross disable-pointer-events bi-2x u-mr--xs h-color-brand-ui" data-test="icon-cross" focusable="false"><path fill-rule="evenodd" d="M7 7.378l3.363 3.363a.496.496 0 00.703-.004c.19-.19.197-.51.004-.703L7.707 6.67l3.363-3.363a.496.496 0 00-.004-.703.502.502 0 00-.703-.004L7 5.964 3.637 2.6a.496.496 0 00-.703.004.502.502 0 00-.004.703L6.293 6.67 2.93 10.034a.496.496 0 00.004.703c.19.19.51.197.703.004L7 7.378z"/></svg>
          </button>
        </div>
          <div tabindex="0" class="srt">Prev handle</div>
        <ul class="account-nav" role="menu" id="accountNav">
  
    <li class="account-nav__item">
      <a href="/nl/rnwy/account/overzicht" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Accountoverzicht">
        <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-account svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-account" focusable="false"><path d="M12 24c-3 0-7.1-.7-9-2.5-.6-.6-1-1.4-1-2.2 0-3.8 2.6-7.2 6.3-8.6C6.9 9.6 6 7.9 6 6c0-3.3 2.7-6 6-6s6 2.7 6 6c0 1.9-.9 3.6-2.3 4.7 3.6 1.4 6.3 4.8 6.3 8.6 0 3.4-6 4.7-10 4.7zm0-12c-4.3 0-8 3.3-8 7.3 0 .1 0 .4.4.7 1.1 1.1 4.3 2 7.6 2 4.5 0 8-1.5 8-2.7 0-4-3.7-7.3-8-7.3zm0-10C9.8 2 8 3.8 8 6s1.8 4 4 4 4-1.8 4-4-1.8-4-4-4z"/></svg>
        <span class="h-block--inline  vatop">Mijn bol</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/rnwy/account/bestellingen/overzicht" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Bestellingen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-box svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-box" focusable="false"><path fill-rule="evenodd" d="M10 1H4L1 5v9h12.038L13 5l-3-4zm2 4H7.5V2h2.25L12 5zM4.25 2H6.5v3H2l2.25-3zM2 13h10V6H2v7z"/></svg>
        <span class="h-block--inline  vatop">Bestellingen</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/rnwy/account/facturen" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Facturen betalen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-payment svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-payment" focusable="false"><path fill-rule="evenodd" d="M5 9a4 4 0 114-4 4 4 0 11-4 4zm.127-1.003c.109-.423.286-.82.518-1.176a2.245 2.245 0 01-.174.006c-.392 0-.71-.107-.956-.322-.245-.215-.404-.526-.48-.932h-.37v-.374h.332l-.005-.09v-.09l.005-.16h-.332v-.375h.366c.062-.409.217-.73.466-.962.25-.233.57-.35.964-.35.326 0 .617.072.874.215l-.205.457c-.25-.124-.473-.186-.669-.186a.79.79 0 00-.562.206c-.146.136-.243.343-.292.62H5.72v.376H4.57l-.005.11v.134l.005.095h.996v.374h-.954c.104.507.4.761.888.761.205 0 .413-.039.626-.116a3.996 3.996 0 011.871-1.091 3 3 0 10-2.87 2.87zM9 12a3 3 0 100-6 3 3 0 000 6zm.461-4.342a.79.79 0 00-.562.206c-.146.136-.243.343-.292.62H9.72v.376H8.57l-.005.11v.134l.005.095h.996v.374h-.954c.104.507.4.761.888.761.233 0 .472-.05.716-.151v.495c-.213.1-.462.15-.745.15-.392 0-.71-.108-.956-.323-.245-.215-.404-.526-.48-.932h-.37v-.374h.332l-.005-.09v-.09l.005-.16h-.332v-.375h.366c.062-.409.217-.73.466-.962.25-.233.57-.35.964-.35.326 0 .617.072.874.215l-.205.457c-.25-.124-.473-.186-.669-.186z"/></svg>
        <span class="h-block--inline  vatop">Facturen betalen</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/account/retourneren/overzicht" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Retourneren">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-return svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-return" focusable="false"><path fill-rule="nonzero" d="M0 7l1.5-3H4V3a1 1 0 011-1h8v8.962L0 11V7zm2.5 6.002a1.5 1.5 0 110-3 1.5 1.5 0 010 3zm8 0a1.5 1.5 0 110-3 1.5 1.5 0 010 3zM13 3H5.5a.5.5 0 00-.5.5V10h8V3zM1 10h3V5H1.96L1 7v3zm1.5 2.002a.5.5 0 100-1 .5.5 0 000 1zm8 0a.5.5 0 100-1 .5.5 0 000 1zM12.5 4h-2a.5.5 0 100 1h2a.5.5 0 000-1zm0 4h-2a.5.5 0 100 1h2a.5.5 0 000-1zm1-2h-2a.5.5 0 100 1h2a.5.5 0 000-1z"/></svg>
        <span class="h-block--inline  vatop">Retourneren</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/account/chat/overzicht" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Chat-geschiedenis">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chat svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-chat" focusable="false"><path fill-rule="evenodd" d="M12 13s-1.471.016-2.5-1.5c-.771.284-1.643.5-2.5.5-4 0-6-3-6-5s2-5 6-5 6 3 6 5-1 3-1 3v3zm-1-1.374V9.578S12 8.88 12 7c0-1.79-1.891-4-5-4-3.328 0-5 2.44-5 4s1.667 4 5 4c.714 0 1.494-.167 2.89-.71.159.19.507.928 1.11 1.336z"/></svg>
        <span class="h-block--inline  vatop">Chat-geschiedenis</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/rnwy/account/voorkeuren" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Gegevens &amp; Voorkeuren">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-settings svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-settings" focusable="false"><g fill-rule="nonzero"><path d="M5.834 2.999a4.141 4.141 0 00-.838.347l-.49-.245c-.247-.123-.608-.07-.8.121l-.484.485c-.183.183-.246.55-.12.799l.244.49A4.141 4.141 0 003 5.834l-.52.173c-.261.087-.479.38-.479.65v.685c0 .26.215.563.48.651l.519.173c.085.294.202.575.347.838l-.245.49c-.123.247-.07.608.121.8l.485.484c.183.183.55.246.799.12l.49-.244c.263.145.544.262.838.347l.173.52c.087.261.38.479.65.479h.685c.26 0 .563-.215.651-.48l.173-.519c.294-.085.575-.202.838-.347l.49.245c.247.123.608.07.8-.121l.484-.485c.183-.183.246-.55.12-.799l-.244-.49c.145-.263.262-.544.347-.838l.52-.173c.261-.087.479-.38.479-.65v-.685c0-.26-.215-.563-.48-.651l-.519-.173a4.141 4.141 0 00-.347-.838l.245-.49c.123-.247.07-.608-.121-.8l-.485-.484c-.183-.183-.55-.246-.799-.12l-.49.244A4.141 4.141 0 008.166 3l-.173-.52C7.906 2.219 7.613 2 7.343 2h-.685c-.26 0-.563.215-.651.48l-.173.519zm-.776-.836C5.281 1.496 5.962 1 6.658 1h.684c.706 0 1.378.497 1.6 1.163l.018.055.035.015.052-.026c.63-.315 1.462-.184 1.954.308l.484.484c.499.5.622 1.326.308 1.954l-.026.052.015.035.055.018c.667.223 1.163.904 1.163 1.6v.684c0 .706-.497 1.378-1.163 1.6l-.055.018-.015.035.026.052c.315.63.184 1.462-.308 1.954l-.484.484c-.5.499-1.326.622-1.954.308l-.052-.026-.035.015-.018.055c-.223.667-.904 1.163-1.6 1.163h-.684c-.706 0-1.378-.497-1.6-1.163l-.018-.055a5.143 5.143 0 01-.035-.015l-.052.026c-.63.315-1.462.184-1.954-.308L2.515 11c-.499-.5-.622-1.326-.308-1.954l.026-.052a5.14 5.14 0 01-.015-.035l-.055-.018C1.496 8.719 1 8.038 1 7.342v-.684c0-.706.497-1.378 1.163-1.6l.055-.018a5.14 5.14 0 01.015-.035l-.026-.052c-.315-.63-.184-1.462.308-1.954L3 2.515c.5-.499 1.326-.622 1.954-.308l.052.026a5.14 5.14 0 01.035-.015l.018-.055z"/><path d="M7 8.5a1.5 1.5 0 100-3 1.5 1.5 0 000 3zm0 1a2.5 2.5 0 110-5 2.5 2.5 0 010 5z"/></g></svg>
        <span class="h-block--inline  vatop">Gegevens &amp; Voorkeuren</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/account/cadeaubonnen" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Cadeaubonnen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-wrap svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-wrap" focusable="false"><path fill-rule="evenodd" d="M2.731 2.693c-.57-.988-.69-1.928-.217-2.4C3.101-.294 4.41.03 5.634 1h2.711C9.57.03 10.877-.294 11.465.293c.469.47.354 1.402-.208 2.383L13 5.001l.038 9H1V5l1.731-2.308zM2 6v7h10V6H2zm2.61-2.548C3.461 2.304 3.307 1.248 3.388 1.167c.08-.081 1.15.083 2.254 1.186.735.735 1.05 1.407.946 2.014-.552.157-1.212-.148-1.978-.914zm4.758 0c-.766.766-1.425 1.071-1.978.914-.104-.607.211-1.279.947-2.014C9.44 1.25 10.51 1.086 10.59 1.167c.08.08-.074 1.137-1.223 2.286zm2.386 1.548H8.602c.734-.196 1.417-.685 2.05-1.466l1.102 1.466zM2.231 5c.08-.104.447-.593 1.102-1.466.633.781 1.317 1.27 2.051 1.466H2.231zm4.41-3h.685a6.761 6.761 0 00-.333.56 4.813 4.813 0 00-.352-.56zM6.489 6h1v7h-1V6zm-1 0h3v7h-3V6z"/></svg>
        <span class="h-block--inline  vatop">Cadeaubonnen</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/account/abonnement.html" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Abonnementen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-replenishment svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-replenishment" focusable="false"><g fill-rule="evenodd"><path d="M1 7a6 6 0 0110.25-4.235v1.6A5 5 0 002 7H1zm1.75 4.235v-1.6A5 5 0 0012 7h1a6 6 0 01-10.25 4.235z"/><path d="M12.124 3.775l-.002.355a.504.504 0 01-.506.5H9.21a.507.507 0 01-.504-.5c0-.276.224-.5.504-.5h1.917V1.706c0-.275.232-.498.5-.498.276 0 .5.217.5.498v2.02c0 .016 0 .033-.002.049zM1.919 10.267l.001-.354a.504.504 0 01.506-.5h2.408c.278 0 .504.232.504.5 0 .276-.225.5-.505.5H2.916v1.923a.504.504 0 01-.5.498.495.495 0 01-.5-.498v-2.019l.003-.05z"/></g></svg>
        <span class="h-block--inline  vatop">Mijn Select &amp; Kobo Plus</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/account/bookshelf.html" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Digitale producten">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-ebook svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-ebook" focusable="false"><path fill-rule="evenodd" d="M11 0H3a1 1 0 00-1 1v12a1 1 0 001 1h8a1 1 0 001-1V1a1 1 0 00-1-1zM3 8.5c-.017.999-.017 1.499 0 1.5h8V1.5a.5.5 0 00-.5-.5h-7a.5.5 0 00-.5.5v7zM3 11h8v1.5a.5.5 0 01-.5.5h-7a.5.5 0 01-.5-.5V11zm1.491-6a.5.5 0 010-1h5.018a.5.5 0 010 1H4.491zm0-2a.5.5 0 010-1h5.018a.5.5 0 010 1H4.491zm0 4.003a.5.5 0 010-1h5.018a.5.5 0 010 1H4.491zm0 1.997a.5.5 0 010-1h5.018a.5.5 0 010 1H4.491z"/></svg>
        <span class="h-block--inline  vatop">Digitale producten</span>
      </a>
  
    <li class="account-nav__item">
      <a href="//partner.bol.com/sdd" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Verkopen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-creditcard svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-creditcard" focusable="false"><path fill-rule="evenodd" d="M2 10H1a1 1 0 01-1-1V3a1 1 0 011-1h10a1 1 0 011 1v1h1a1 1 0 011 1v6a1 1 0 01-1 1H3a1 1 0 01-1-1v-1zm0-1V7H1v1.5a.5.5 0 00.5.5H2zm0-3V5a1 1 0 011-1h8v-.5a.5.5 0 00-.5-.5h-9a.5.5 0 00-.5.5v1C.983 5.499.983 5.999 1 6h1zm1 3.5v1a.5.5 0 00.5.5h9a.5.5 0 00.5-.5v-5a.5.5 0 00-.5-.5h-9a.5.5 0 00-.5.5v4zM4 6h4v1H4V6zm0 2h2v1H4V8zm3 0h2v1H7V8zm3 0h2v1h-2V8z"/></svg>
        <span class="h-block--inline  vatop">Verkopen</span>
      </a>
  
    <li class="account-nav__item">
      <a href="/nl/rnwy/logout.html" class="account-nav__link" data-analytics-id="px_common_navigation" data-analytics-tag="persoonlijk menu:Uitloggen">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-logout svg-inline--bi bi-account bi-lg  h-boxedright--xs vatop" data-test="icon-logout" focusable="false"><g fill-rule="evenodd"><path d="M12.601 6.393l.25.252a.504.504 0 01-.004.711L11.144 9.06a.507.507 0 01-.71.002.501.501 0 01.004-.71l.855-.855H6v-1h5.293l-.86-.86a.504.504 0 01.001-.706.495.495 0 01.706-.002l1.428 1.428a.485.485 0 01.033.037z"/><path fill-rule="nonzero" d="M10.348 2.636l-.039-.03a5.5 5.5 0 10-.001 8.788.5.5 0 10-.602-.798 4.5 4.5 0 110-7.19l.033.023a.5.5 0 10.61-.793z"/></g></svg>
        <span class="h-block--inline  vatop">Uitloggen</span>
      </a>
  
</ul>

          <div tabindex="0" class="srt" id="nextHandle">Next handle</div>
      </div>
    </div>
</div>

<button class="account-button  js_offcanvas_open  u-hide@screen-xl-up" data-analytics-id="px_common_click" data-analytics-tag="header menu:profile icon" data-offcanvas-target="ID2offcanvas-container--right">
    <span class="srt">Account menu</span>
    <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-account bi-lg disable-pointer-events" data-test="icon-account" focusable="false"><path d="M12 24c-3 0-7.1-.7-9-2.5-.6-.6-1-1.4-1-2.2 0-3.8 2.6-7.2 6.3-8.6C6.9 9.6 6 7.9 6 6c0-3.3 2.7-6 6-6s6 2.7 6 6c0 1.9-.9 3.6-2.3 4.7 3.6 1.4 6.3 4.8 6.3 8.6 0 3.4-6 4.7-10 4.7zm0-12c-4.3 0-8 3.3-8 7.3 0 .1 0 .4.4.7 1.1 1.1 4.3 2 7.6 2 4.5 0 8-1.5 8-2.7 0-4-3.7-7.3-8-7.3zm0-10C9.8 2 8 3.8 8 6s1.8 4 4 4 4-1.8 4-4-1.8-4-4-4z"/></svg>
</button>





        <a class="customer-lists px_common_customermenu_click" href="/nl/nl/lijstjes/overzicht-lijsten/" data-test="customer-lists-button">
            <span class="srt">Lijstjes</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-heart-default bi-lg" data-test="icon-heart-default" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M14.4612 2.49501C15.2499 2.16821 16.0953 2 16.949 2C17.8027 2 18.6481 2.16821 19.4368 2.49501C20.2254 2.82176 20.9419 3.30064 21.5454 3.90431C22.149 4.50779 22.6282 5.2246 22.955 6.01317C23.2818 6.80189 23.45 7.64726 23.45 8.50101C23.45 9.35475 23.2818 10.2001 22.955 10.9888C22.7593 11.4612 22.4407 11.8981 22.1441 12.2526C21.8405 12.6157 21.5108 12.9493 21.2543 13.2057C21.2542 13.2057 21.2543 13.2056 21.2543 13.2057L12.7555 21.7044C12.568 21.892 12.3137 21.9973 12.0484 21.9973C11.7832 21.9973 11.5289 21.892 11.3413 21.7044L2.8426 13.2057C1.73995 12.103 0.550003 10.3536 0.550003 8.50101C0.550003 6.77697 1.23487 5.12355 2.45395 3.90448C3.67302 2.6854 5.32644 2.00053 7.05047 2.00053C8.77451 2.00053 10.4279 2.6854 11.647 3.90448L11.9997 4.25721L12.3523 3.90464C12.9558 3.30082 13.6724 2.82182 14.4612 2.49501ZM16.949 4C16.358 4 15.7728 4.11645 15.2267 4.34268C14.6807 4.56892 14.1847 4.90051 13.7669 5.31853L12.7068 6.37853C12.3163 6.76906 11.6832 6.76906 11.2926 6.37853L10.2328 5.31869C9.38878 4.47469 8.24407 4.00053 7.05047 4.00053C5.85687 4.00053 4.71216 4.47469 3.86816 5.31869C3.02416 6.16269 2.55 7.30741 2.55 8.50101C2.55 9.56601 3.29639 10.8311 4.25681 11.7915L12.0484 19.5831L19.8401 11.7915C20.0945 11.5372 20.3686 11.2581 20.61 10.9695C20.8586 10.6723 21.0265 10.4183 21.1073 10.2233C21.3336 9.67725 21.45 9.09202 21.45 8.50101C21.45 7.90999 21.3336 7.32476 21.1073 6.77876C20.8811 6.23275 20.5495 5.73667 20.1315 5.31886C19.7137 4.90085 19.2173 4.56892 18.6712 4.34268C18.1252 4.11644 17.54 4 16.949 4Z" fill="currentColor"/></svg>
        </a>

        <wsp-basket-header-icon>
  <a href="/nl/nl/basket/" class="main-basket-btn  px_common_orders_click  js_basket_wrapper" id="basket"
    data-quantity="0" data-test="basket-button">
    <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-basket bi-lg" data-test="icon-basket" focusable="false"><path d="M23.9 6.5c0-.8-.7-1.5-1.5-1.5H5.8l-1-4H1c-.6 0-1 .4-1 1s.4 1 1 1h2.2L4 6.2l2.9 11.6c-.5.6-.9 1.4-.9 2.2 0 1.7 1.3 3 3 3s3-1.3 3-3c0-.4-.1-.7-.2-1h4.4c-.1.3-.2.6-.2 1 0 1.7 1.3 3 3 3s3-1.3 3-3c0-.9-.4-1.6-.9-2.2l2.7-11c0-.1.1-.2.1-.3zM10 20c0 .6-.4 1-1 1s-1-.4-1-1 .4-1 1-1 1 .4 1 1zm9 1c-.6 0-1-.4-1-1s.4-1 1-1 1 .4 1 1-.4 1-1 1zm.2-4H8.8L6.3 7h15.4l-2.5 10z"/></svg>
    <span class="srt">Winkelwagen</span>
    <span class="srt js_basket_amount" data-test="basket-button-quantity">0</span>
  </a>
</wsp-basket-header-icon>
    </div>
</div>

  <wsp-main-nav-offcanvas class="wsp-offcanvas  wsp-offcanvas--main-menu  js_main_menu">
  <div class="wsp-offcanvas__header">
    <div class="u-pl--s">
      <wsp-country-language-selector data-modal-url="/nl/rnwy/country-language"  class="o-drop-down__label c-media c-media--center hit-area u-hide@screen-small-only">
    <span class="c-country-icon c-country-icon--nl c-country-icon--s" data-test="country-flag"></span>
    <span class="u-pl--xxs" data-test="language">NL</span>
    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-down bi-md" data-test="icon-chevron-down" focusable="false"><path fill-rule="evenodd" d="M3.635 4.694a.508.508 0 00-.71 0 .497.497 0 00-.007.7l4.085 4.06 4.054-4.055a.505.505 0 000-.707.501.501 0 00-.71.002L6.997 8.046 3.635 4.694z"/></svg>
    <a class="hit-area__link" href="/nl/rnwy/country-language" aria-label="Kies je land en taal, land is Nederland" data-analytics-id="px_country_language_click" data-analytics-value="Language_modal_desktop_menu:open" data-test="country-select"></a>
</wsp-country-language-selector>

    </div>
      <button class="wsp-offcanvas-action-ab wsp-offcanvas-action-ab--back js_offcanvas_back u-hide@screen-xl-up">
        <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-back bi-lg" data-test="icon-chevron-back" focusable="false"><path fill-rule="evenodd" d="M9.306 3.644a.508.508 0 000-.71.497.497 0 00-.7-.006l-4.06 4.084 4.055 4.054c.195.195.517.19.707 0a.501.501 0 00-.002-.71L5.954 7.006l3.352-3.36z"/></svg>
      </button>
      <span class="wsp-offcanvas__title js_offcanvas_title u-hide@screen-xl-up"></span>
    <button class="wsp-offcanvas-action-ab  wsp-offcanvas-action-ab--close  js_close_offcanvas" aria-expanded="false" aria-label="ariaCloseMenu">
      <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-cross bi-lg" data-test="icon-cross" focusable="false"><path fill-rule="evenodd" d="M7 7.378l3.363 3.363a.496.496 0 00.703-.004c.19-.19.197-.51.004-.703L7.707 6.67l3.363-3.363a.496.496 0 00-.004-.703.502.502 0 00-.703-.004L7 5.964 3.637 2.6a.496.496 0 00-.703.004.502.502 0 00-.004.703L6.293 6.67 2.93 10.034a.496.496 0 00.004.703c.19.19.51.197.703.004L7 7.378z"/></svg>
    </button>
  </div>

  <div class="wsp-offcanvas--scroll-pane  js_scoll_pane">
    <div class="wsp-offcanvas--slide  js_slide_to_sub">
      <div class="wsp-offcanvas__slide-content  js_slide_content">
        <nav class="wsp-header-sections__nav" role="navigation" aria-label="ariaMainNavigation">
          <ul class="wsp-main-nav">
            <wsp-main-nav-item class="wsp-main-nav__item  wsp-main-nav__item--category  wsp-main-nav__item--icons" data-name="category-menu">
  <button class="wsp-main-nav__link  js_category_menu_button" data-test="category-menu">
    
    Categorieën
  </button>

  <div class="wsp-main-nav-constrain"  data-analytics-id="px_main_nav_panel" data-analytics-inview-id="debounced" data-analytics-inview-event-id="315" data-analytics-inview-debounce-time="1000">
    <div class="wsp-main-nav-panel   js_sub_panel">
      <wsp-main-nav-category-ab async-url="/nl/nl/menu/categories" isVariantA="false">
  <ul class="wsp-category-nav-ab">
    <li><strong class="wsp-category-nav-ab__title">Categorieën</strong></li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="1">
      <a href="/nl/nl/menu/categories/subMenu/1" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/Jl2gQN1Z61Ml/ANn2YEz/160x160.jpg" alt="Boeken" loading="lazy"/>
        <span>Boeken</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="2">
      <a href="/nl/nl/menu/categories/subMenu/2" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/Y0nA5gn9qXmA/K1O2Vln/160x160.jpg" alt="Muziek, Film &amp; Gaming" loading="lazy"/>
        <span>Muziek, Film &amp; Gaming</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="3">
      <a href="/nl/nl/menu/categories/subMenu/3" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/Y0nA5XmJXWJM/D1R2VKn/160x160.jpg" alt="Computer &amp; Elektronica" loading="lazy"/>
        <span>Computer &amp; Elektronica</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="4">
      <a href="/nl/nl/menu/categories/subMenu/4" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/Jl2gE9mq6NJD/JyZ29ko/160x160.jpg" alt="Speelgoed, Hobby &amp; Feest" loading="lazy"/>
        <span>Speelgoed, Hobby &amp; Feest</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="5">
      <a href="/nl/nl/menu/categories/subMenu/5" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/g6GjnvQ133r3/wVjAq6w/160x160.jpg" alt="Zwanger, Baby &amp; Peuter" loading="lazy"/>
        <span>Zwanger, Baby &amp; Peuter</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="6">
      <a href="/nl/nl/menu/categories/subMenu/6" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/3OpnYljOmqq4/BNg2YGJ/160x160.jpg" alt="Mooi &amp; Gezond" loading="lazy"/>
        <span>Mooi &amp; Gezond</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="7">
      <a href="/nl/nl/menu/categories/subMenu/7" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/7YBK0VMNz1ZB/16w1kRP/160x160.jpg" alt="Kleding, Schoenen &amp; Accessoires" loading="lazy"/>
        <span>Kleding, Schoenen &amp; Accessoires</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="8">
      <a href="/nl/nl/menu/categories/subMenu/8" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/qDr0AvqBmV2p/E8l2VLk/160x160.jpg" alt="Sport, Outdoor &amp; Reizen" loading="lazy"/>
        <span>Sport, Outdoor &amp; Reizen</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="12">
      <a href="/nl/nl/menu/categories/subMenu/12" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/qDr0ADLr5rP2/GvZ2VN0/160x160.jpg" alt="Kantoor &amp; School" loading="lazy"/>
        <span>Kantoor &amp; School</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="10">
      <a href="/nl/nl/menu/categories/subMenu/10" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/BQ2J8m1n4Qyk/L7g2Vm4/160x160.jpg" alt="Eten &amp; Drinken" loading="lazy"/>
        <span>Eten &amp; Drinken</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="11">
      <a href="/nl/nl/menu/categories/subMenu/11" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/qDr0MxLoyZJ7/zvmEwg8/160x160.jpg" alt="Wonen, Koken &amp; Huishouden" loading="lazy"/>
        <span>Wonen, Koken &amp; Huishouden</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="9">
      <a href="/nl/nl/menu/categories/subMenu/9" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/Rv9B6v1VMNZq/VOZpxoM/160x160.jpg" alt="Kerst, Dier, Tuin &amp; Klussen" loading="lazy"/>
        <span>Kerst, Dier, Tuin &amp; Klussen</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
    <li class="wsp-category-nav-ab__item js_cat_item" data-nav-id="13">
      <a href="/nl/nl/menu/categories/subMenu/13" class="wsp-category-nav-ab__link js_cat_link">
        <img class="wsp-category-nav-ab__image" width="80" height="80" src="https://media.s-bol.com/3OpnyxX8yg84/vlgzp6r/160x160.jpg" alt="Auto, Motor &amp; Fiets" loading="lazy"/>
        <span>Auto, Motor &amp; Fiets</span>
      </a>
      <div class="wsp-category-nav-ab-sub-panel">
        <div class="inner-panel">
          <div class="inner-panel__spacer  js_sub_panel"></div>
        </div>
      </div>
    </li>
    
  </ul>
</wsp-main-nav-category-ab>

    </div>
  </div>
</wsp-main-nav-item>
<wsp-main-nav-item class="wsp-main-nav__item  js-nav-item" omniture-value="inspiration" data-name="inspiration-menu">
  <button class="wsp-main-nav__link  js_category_menu_button" data-test="inspiration-menu">
    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-inspiration bi-lg" data-test="icon-inspiration" focusable="false"><path fill-rule="nonzero" d="M7 .5c3 0 4 2.5 4 4s-.86 2.26-1 2.5c-.667.667-1 1.5-1 2.5H5c0-1-.333-1.833-1-2.5-.17-.17-1-1-1-2.5s1-4 4-4zm0 .981c-2.25 0-3 1.875-3 3s.622 1.747.75 1.875c.567.57.953 1.29 1.157 2.156l2.186.015c.228-.938.614-1.661 1.157-2.17.159-.158.75-.75.75-1.876 0-1.125-.75-3-3-3zM5 10.5h4a1 1 0 01-1 1H6a1 1 0 01-1-1zm1 2h2a1 1 0 01-2 0z"/></svg>
    Cadeaus &amp; Inspiratie
  </button>

  <div class="wsp-main-nav-constrain" >
    <div class="wsp-main-nav-content js_sub_panel">
      
  <ul class="wsp-sub-nav  u-hide@screen-xl-up">
    
      <li class="wsp-sub-nav__item">
        <button data-test="sub-nav-link" class="wsp-sub-nav__link wsp-sub-nav__link--back js_offcanvas_back"
        >
          Naar overzicht
        </button>
      </li>
    
  </ul>


<div class="o-layout o-layout--fit-content@screen-xl-up">
  
  <div class="o-layout__item" data-test="sub-nav-column">
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Cadeaus</div></li>

      
        
        <li class="wsp-sub-nav__item u-hide u-show-block@screen-xl-up" data-test="wsp-sub-nav-list-item"
             omniture-value="Cadeaus:Cadeaus">
        
          <strong>

          
            <span>
          

            
              
              Cadeaus
            

          
            </span>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cadeaus:Cadeaus volwassenen">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/cadeauwinkel-volwassenen/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Cadeaus volwassenen
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cadeaus:Cadeaus kinderen">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/cadeauwinkel-kinderen/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Cadeaus kinderen
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cadeaus:Cadeaubonnen">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/cadeaubon-bestellen/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Cadeaubonnen
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Shop de Look</div></li>

      
        
        <li class="wsp-sub-nav__item u-hide u-show-block@screen-xl-up" data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Shop de Look">
        
          <strong>

          
            <span>
          

            
              
              Shop de Look
            

          
            </span>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Huis">
        
          

          
            <a href="https://www.bol.com/nl/nl/sdl/wonen/8" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Huis
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Damesmode">
        
          

          
            <a href="https://www.bol.com/nl/nl/sdl/mode/dames/1" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Damesmode
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Herenmode">
        
          

          
            <a href="https://www.bol.com/nl/nl/sdl/mode/heren/2" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Herenmode
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Kindermode">
        
          

          
            <a href="https://www.bol.com/nl/nl/sdl/mode/kinderen/3" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Kindermode
            

          
            </a>
          

          

        
          </li>
        

        
          <li class="wsp-sub-nav__item wsp-sub-nav__item--collapse-after js_offcanvas_show_more">
            <button class="wsp-sub-nav__link wsp-sub-nav__link--show-more">Meer categorieën</button>
          </li>
        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Shop de Look:Baby &amp; Peuter">
        
          

          
            <a href="https://www.bol.com/nl/nl/sdl/baby-peuter/42" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Baby &amp; Peuter
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
  </div>
  
  <div class="o-layout__item" data-test="sub-nav-column">
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Inspiratie</div></li>

      
        
        <li class="wsp-sub-nav__item u-hide u-show-block@screen-xl-up" data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Inspiratie">
        
          <strong>

          
            <span>
          

            
              
              Inspiratie
            

          
            </span>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Kerstinspiratie">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/kerst/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Kerstinspiratie
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Wooninspiratie">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/wooninspiratie/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Wooninspiratie
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Boeken inspiratie">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/lees/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Boeken inspiratie
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Kinderboeken inspiratie">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/kinderboeken/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Kinderboeken inspiratie
            

          
            </a>
          

          

        
          </li>
        

        
          <li class="wsp-sub-nav__item wsp-sub-nav__item--collapse-after js_offcanvas_show_more">
            <button class="wsp-sub-nav__link wsp-sub-nav__link--show-more">Meer categorieën</button>
          </li>
        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Goede keuze assortiment">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/goede-keuze-assortiment/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Goede keuze assortiment
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Refurbished via bol">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/refurbished-elektronica/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Refurbished via bol
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Smart home">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/smart-home/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Smart home
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Thuiswerken">
        
          

          
            <a href="https://www.bol.com/nl/nl/inf/wat-heb-je-nodig-voor-je-thuiswerkplek/52649/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Thuiswerken
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Puppywinkel">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/puppywinkel/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Puppywinkel
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Inspiratie:Kittenwinkel">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/kittenwinkel/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Kittenwinkel
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
  </div>
  
  <div class="o-layout__item" data-test="sub-nav-column">
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Nog meer inspiratie</div></li>

      
        
        <li class="wsp-sub-nav__item u-hide u-show-block@screen-xl-up" data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Nog meer inspiratie">
        
          <strong>

          
            <span>
          

            
              
              Nog meer inspiratie
            

          
            </span>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Parfum &amp; luxe verzorging">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/parfum-en-luxe-verzorging/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Parfum &amp; luxe verzorging
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Sporten">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/sporten/1832/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Sporten
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Gezondheidsadvies">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/gezondheidsplatform/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Gezondheidsadvies
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Zwangerschap">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/zwanger/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Zwangerschap
            

          
            </a>
          

          

        
          </li>
        

        
          <li class="wsp-sub-nav__item wsp-sub-nav__item--collapse-after js_offcanvas_show_more">
            <button class="wsp-sub-nav__link wsp-sub-nav__link--show-more">Meer categorieën</button>
          </li>
        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Nieuwe collectie mode">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/nieuwecollectie/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Nieuwe collectie mode
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Energiezuinig witgoed">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/energiezuinig-witgoed/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Energiezuinig witgoed
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Laptops">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/laptops/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Laptops
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Gaming">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/gaming/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Gaming
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Merchandise">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/merchandise/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Merchandise
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Nog meer inspiratie:Funcooking">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/funcooking-kerstfeest-in-de-keuken/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Funcooking
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
  </div>
  
</div>



    </div>
  </div>
</wsp-main-nav-item>
<wsp-main-nav-item class="wsp-main-nav__item  js-nav-item" omniture-value="deals" data-name="deals-menu">
  <button class="wsp-main-nav__link  js_category_menu_button" data-test="deals-menu">
    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-tag bi-lg" data-test="icon-tag" focusable="false"><path fill-rule="evenodd" d="M14 5.956V.996A.998.998 0 0013.005 0H8.046L.715 7.29a1.007 1.007 0 00-.002 1.422l4.576 4.577a1 1 0 001.422-.002l7.29-7.331zM1.673 8.171a.244.244 0 01.002-.344L8.501 1l4.252-.012c.134 0 .242.11.242.242l.006 4.27-6.827 6.827a.245.245 0 01-.344.003L1.672 8.17zm9.33-6.114a1 1 0 100 2 1 1 0 000-2z"/></svg>
    Aanbiedingen
  </button>

  <div class="wsp-main-nav-constrain" >
    <div class="wsp-main-nav-content js_sub_panel">
      
  <ul class="wsp-sub-nav  u-hide@screen-xl-up">
    
      <li class="wsp-sub-nav__item">
        <button data-test="sub-nav-link" class="wsp-sub-nav__link wsp-sub-nav__link--back js_offcanvas_back"
        >
          Naar overzicht
        </button>
      </li>
    
  </ul>


<div class="o-layout o-layout--fit-content@screen-xl-up">
  
  <div class="o-layout__item" data-test="sub-nav-column">
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Aanbiedingen</div></li>

      
        
        <li class="wsp-sub-nav__item u-hide u-show-block@screen-xl-up" data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Aanbiedingen">
        
          <strong>

          
            <span>
          

            
              
              Aanbiedingen
            

          
            </span>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Alle aanbiedingen">
        
          

          
            <a href="https://www.bol.com/nl/nl/deals/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Alle aanbiedingen
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Dagdeal">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/dagdeal/610/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Dagdeal
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Topdeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/topdeals/369532" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Topdeals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Voorraaddeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/voorraaddeals/1916/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Voorraaddeals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Elektronica acties">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/elektronica-acties/1887/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Elektronica acties
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Woondeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/woondeals/1912/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Woondeals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Outlet">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/outlet/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Outlet
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Klus- en tuinaanbiedingen">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/klusentuinaanbiedingen/1935/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Klus- en tuinaanbiedingen
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Retourdeals: tot 15% goedkoper *">
        
          

          
            <a href="https://www.bol.com/nl/nl/sf/retourdeals/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Retourdeals: tot 15% goedkoper *
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Speelgoeddeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/speelgoeddeals/1889/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Speelgoeddeals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Select-Deals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/select-deals/1826/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Select-Deals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Gamingdeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/gamingdeals/1906/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Gamingdeals
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Aanbiedingen:Fashion &amp; sportdeals">
        
          

          
            <a href="https://www.bol.com/nl/nl/cmp/fas/1996/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Fashion &amp; sportdeals
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
  </div>
  
  <div class="o-layout__item" data-test="sub-nav-column">
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Sint</div></li>

      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Sint:Sint">
        
          <strong>

          
            <a href="https://www.bol.com/nl/nl/cmp/sinterklaas/1992/ " class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Sint
            

          
            </a>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Sint:2e halve prijs op speelgoed">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/2e-halve-prijs-op-speelgoed/355255/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              2e halve prijs op speelgoed
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Sint:2 kinderboeken voor €30">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/2-kinderboeken-voor-30-euro/358613/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              2 kinderboeken voor €30
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Sint:Korting* op cadeaus, chocolade &amp; adventkalenders">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/korting-op-cadeaus-chocolade-adventkalenders/363553/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op cadeaus, chocolade &amp; adventkalenders
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Cyber Weekend</div></li>

      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cyber Weekend:Cyber Weekend">
        
          <strong>

          
            <a href="https://www.bol.com/nl/nl/cmp/black-friday/2002/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Cyber Weekend
            

          
            </a>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cyber Weekend:Korting* op Sony">
        
          

          
            <a href="https://www.bol.com/nl/nl/p/sony-wh-1000xm4-draadloze-over-ear-koptelefoon-met-noise-cancelling-zwart/9300000006173040/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op Sony
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cyber Weekend:Korting* op espressomachines">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/tot-100-korting-op-volautomatische-espressomachines/362256/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op espressomachines
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Cyber Weekend:Korting* op airfryers en frituurpannen">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/hoge-korting-op-airfryers-en-frituurpannen/362250/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op airfryers en frituurpannen
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
    <ul class="wsp-sub-nav">
      <li class="wsp-sub-nav__item u-hide@screen-xl-up"><div class="wsp-sub-nav__title" data-test="sub-nav-column-title">Voorraaddeals</div></li>

      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Voorraaddeals:Voorraaddeals">
        
          <strong>

          
            <a href="https://www.bol.com/nl/nl/cmp/voorraaddeals/1916/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Voorraaddeals
            

          
            </a>
          

          </strong>

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Voorraaddeals:Korting* op huishoudmiddelen">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/korting-op-huishoudmiddelen/362016/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op huishoudmiddelen
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Voorraaddeals:Korting* op luiers en billendoekjes">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/black-friday-deals-op-luiers-en-billendoekjes/344366/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op luiers en billendoekjes
            

          
            </a>
          

          

        
          </li>
        

        
      
        
        <li class="wsp-sub-nav__item " data-test="wsp-sub-nav-list-item"
             omniture-value="Voorraaddeals:Korting* op koffie">
        
          

          
            <a href="https://www.bol.com/nl/nl/ra/korting-op-koffie/357845/" class="wsp-sub-nav__link js_cat_link"
               >
          

            
              
              Korting* op koffie
            

          
            </a>
          

          

        
          </li>
        

        
      
    </ul>
    
  </div>
  
</div>


  <div class="wsp-category-nav__disclaimer small-text  h-color-subtext">* op de adviesprijs</div>


    </div>
  </div>
</wsp-main-nav-item>


            <li class="wsp-main-nav__ghost" aria-hidden="true">
              <div class="constrain"></div>
            </li>
          </ul>

          <ul class="wsp-secundary-nav">
            
              <li class="wsp-secundary-nav__item omniture_extra_services">
                <a class="wsp-secundary-nav__link" href="/nl/nl/m/zakendoen-met-bolcom/" data-test="Zakelijk">Zakelijk</a>
              </li>
            
              <li class="wsp-secundary-nav__item omniture_extra_services">
                <a class="wsp-secundary-nav__link" href="/nl/nl/l/cadeaukaarten/20639/" data-test="Cadeaukaart">Cadeaukaart</a>
              </li>
            
              <li class="wsp-secundary-nav__item omniture_extra_services">
                <a class="wsp-secundary-nav__link" href="/nl/rnwy/account/bestellingen/overzicht" data-test="Bestelstatus">Bestelstatus</a>
              </li>
            
              <li class="wsp-secundary-nav__item omniture_extra_services">
                <a class="wsp-secundary-nav__link" href="/nl/klantenservice/index.html" data-test="Klantenservice">Klantenservice</a>
              </li>
            
            <li class="wsp-secundary-nav__item omniture_customer_menu">
              <wsp-country-language-selector data-modal-url="/nl/rnwy/country-language"  class="o-drop-down__label c-media c-media--center hit-area u-hide@screen-small-only">
    <span class="c-country-icon c-country-icon--nl c-country-icon--s" data-test="country-flag"></span>
    <span class="u-pl--xxs" data-test="language">NL</span>
    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-down bi-md" data-test="icon-chevron-down" focusable="false"><path fill-rule="evenodd" d="M3.635 4.694a.508.508 0 00-.71 0 .497.497 0 00-.007.7l4.085 4.06 4.054-4.055a.505.505 0 000-.707.501.501 0 00-.71.002L6.997 8.046 3.635 4.694z"/></svg>
    <a class="hit-area__link" href="/nl/rnwy/country-language" aria-label="Kies je land en taal, land is Nederland" data-analytics-id="px_country_language_click" data-analytics-value="Language_modal_desktop_menu:open" data-test="country-select"></a>
</wsp-country-language-selector>

            </li>
          </ul>
        </nav>

        <nav class="nav-contextual" role="navigation" data-test="marketing-links" aria-label="Aanbiedingen">
  <div class="nav-contextual__inner">
    
      <a href="https://www.bol.com/nl/nl/cmp/black-friday/2002/?promo&#x3D;HEAD_907_Mnav_1_CRS_cyber_1" class="nav-contextual__link omniture_active_link px_common_marketing_click"
         data-omniture-config="nextpage,menulink"
         data-omniture-navtype="Tab clicks"
         data-omniture-prefix="Marketing Link">
        Cyber Weekend
      </a>
    
      <a href="https://www.bol.com/nl/nl/ra/het-populairste-speelgoed/356008/?promo&#x3D;HEAD_907_Mnav_1_CRS_populairspeelgoed_2" class="nav-contextual__link omniture_active_link px_common_marketing_click"
         data-omniture-config="nextpage,menulink"
         data-omniture-navtype="Tab clicks"
         data-omniture-prefix="Marketing Link">
        Populair speelgoed
      </a>
    
      <a href="https://www.bol.com/nl/nl/l/bol-com-cadeaukaarten/20639/4282200616/?promo&#x3D;HEAD_907_Mnav_1_CRS_cadeaukaart_3" class="nav-contextual__link omniture_active_link px_common_marketing_click"
         data-omniture-config="nextpage,menulink"
         data-omniture-navtype="Tab clicks"
         data-omniture-prefix="Marketing Link">
        Cadeaukaart
      </a>
    
      <a href="https://www.bol.com/nl/nl/sf/kerst/?promo&#x3D;HEAD_907_Mnav_1_CRS_kerst_4" class="nav-contextual__link omniture_active_link px_common_marketing_click"
         data-omniture-config="nextpage,menulink"
         data-omniture-navtype="Tab clicks"
         data-omniture-prefix="Marketing Link">
        Kerstinspiratie
      </a>
    

    <a href="https://www.bol.com/nl/nl/cmp/sinterklaas/1992/?promo&#x3D;HEAD_907_Mnav_1_CRS_sint_5" class="nav-contextual__flagship-link h-color-brand-ui px_common_marketing_click" omniture-value="Last minute sintcadeaus">
    <span class="bol_header" data-test="flagship-link-title">Last minute sintcadeaus</span>
    
    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-next bi-lg" data-test="icon-chevron-next" focusable="false"><path fill-rule="evenodd" d="M4.694 3.644a.508.508 0 010-.71.497.497 0 01.7-.006l4.06 4.084-4.055 4.054a.505.505 0 01-.707 0 .501.501 0 01.002-.71l3.352-3.351-3.352-3.36z"/></svg>
</a>

  </div>
</nav>


        <nav class="u-hide@screen-large-up" role="navigation">
  <wsp-main-nav-item class="wsp-category-nav  u-mb">
    
      <li class="wsp-category-nav__item">
        <a href="/nl/rnwy/account/overzicht" class="wsp-category-nav__link" omniture-value="account">
          <div class="c-media  c-media--center">
            <figure class="c-media__figure">
              <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-person bi-lg" data-test="icon-person" focusable="false"><path fill-rule="evenodd" d="M7 8c3.08 0 6 1 6 3v2H1v-2c0-1.996 2.909-3 6-3zm0 1c-3.091 0-5 .99-5 2v1h10v-1c0-1-1.92-2-5-2zm0-2a3 3 0 110-6 3 3 0 010 6zm0-1a2 2 0 100-4 2 2 0 000 4z"/></svg>
            </figure>
            <div class="c-media__body">
              <p class="u-nosp">Welkom Sebast...</p>
              <p class="small-text">Mijn account</p>
            </div>
          </div>
        </a>
        <div class="wsp-category-nav-sub-panel js_sub_panel"></div>
      </li>
    
    
      <wsp-main-nav-item class="js_cat_link">
        <li class="wsp-category-nav__item">
          <a href="/nl/select" class="wsp-category-nav__link"  omniture-value="select">
            <div class="c-media  c-media--center disable-pointer-events">
              <figure class="c-media__figure disable-pointer-events">
                <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-tag bi-lg disable-pointer-events" data-test="icon-tag" focusable="false"><path fill-rule="evenodd" d="M14 5.956V.996A.998.998 0 0013.005 0H8.046L.715 7.29a1.007 1.007 0 00-.002 1.422l4.576 4.577a1 1 0 001.422-.002l7.29-7.331zM1.673 8.171a.244.244 0 01.002-.344L8.501 1l4.252-.012c.134 0 .242.11.242.242l.006 4.27-6.827 6.827a.245.245 0 01-.344.003L1.672 8.17zm9.33-6.114a1 1 0 100 2 1 1 0 000-2z"/></svg>
              </figure>
                <div class="c-media__body disable-pointer-events">
                  <p class="u-nosp disable-pointer-events">Meer gemak en voordeel met <strong class="txt-select disable-pointer-events" data-test="select-logo">Select</strong></p>
                  <p class="small-text disable-pointer-events">Voor maar € 14,99 per jaar</p>
                </div>
            </div>
          </a>
        </li>
      </wsp-main-nav-item>
    
      <li class="wsp-category-nav__item">
        <a class="wsp-category-nav__link" href="/nl/klantenservice/index.html" omniture-value="customerservice">
          <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-headphone bi-lg" data-test="icon-headphone" focusable="false"><path d="M2 8V7h-.748A.25.25 0 001 7.256v2.488c0 .142.113.256.252.256H2V8zm1-2v5H1.001C.45 11 0 10.552 0 10V7c0-.556.447-1 1-1h1a5 5 0 1110 0h1c.556 0 1 .449 1 1.003V10.006c0 .536-.447.994-1 .994h-2V6a4 4 0 10-8 0zm9 1v3h.748A.25.25 0 0013 9.744V7.256A.254.254 0 0012.748 7H12z"/></svg> Klantenservice
        </a>
      </li>
      
          <li class="wsp-category-nav__item">
    <wsp-country-language-selector data-modal-url="/nl/rnwy/country-language" data-modal-size="full">
        <div class="wsp-category-nav__link c-media">
            <span class="c-country-icon c-country-icon--nl c-country-icon--s u-mr--m u-mt--xxs u-pr"></span>
            <div class="c-media__body">
                <p class="u-mb--xs">Nederland<span class="h-color-subtext u-ml--xs u-mr--xs">|</span>Taal: Nederlands</p>
                <p class="small-text h-color-subtext u-mr--m">Je ziet alleen de artikelen en promoties die  beschikbaar zijn in Nederland.</p>
            </div>
        </div>
    </wsp-country-language-selector>
<li>
      
  </wsp-main-nav-item>

  <wsp-main-nav-item class="wsp-category-nav  u-mb js_cat_link">
    <li class="wsp-category-nav__item">
      <a href="/nl/m/over-bol-com/" class="wsp-category-nav__link" omniture-value="over-bol-com">Over bol</a>
    </li>
    <li class="wsp-category-nav__item">
      <a href="/nl/nl/l/cadeaukaarten-bol-com/20639/4282200616/" class="wsp-category-nav__link"  omniture-value="cadeaubon">Cadeaubon</a>
    </li>
      
          <li class="wsp-category-nav__item">
              <a href="/nl/nl/m/zakendoen-met-bolcom/" class="wsp-category-nav__link"
                 omniture-value="zakelijk">Zakelijk</a>
          </li>
      
      
  </wsp-main-nav-item>
</nav>


        
      </div>
    </div>
  </div>
</wsp-main-nav-offcanvas>

</header>

    
    <main id="mainContent">
        
  <style>
  wsp-async-browse .js_content_loader_wrapper {
    position: relative;
  }

  wsp-async-browse  #js_list_view {
    transition: opacity 200ms linear;
  }

  wsp-async-browse[is-loading] #js_list_view  {
    pointer-events: none;
    opacity: 0.3;
  }

  wsp-async-browse[is-loading] .css-loader-container {
    position: absolute;
    width: 100vw;
    height: 100%;
  }

  wsp-async-browse[is-loading] .css-loader {
    display: block;
    position: sticky;
    left: calc(50% - 1rem);
    top: 50%;
  }

  .offcanvas-wrapper--open-filter_navigation .offcanvas-wrapper__main {
    min-height: 1200px;
  }

  .sort-bar__results {
    order: 1;
  }

</style>



<wsp-async-browse service-backend="Xavier" asyncUrl="/nl/nl/s/" pageMatchingPattern="/nl/nl/s/"   intent="search" >
  <div class="js_content_loader_wrapper">
    
    <script type="application/json" data-quantity-labels>
  {"pieces":"stuks","overlayButtonText":"Ok","overlayTitle":"Wat is het gewenste aantal (max. 500)?","overlayButtonTextAria":"Bevestig het gekozen aantal","overlayButtonCloseButton":"Sluiten"}
</script>


<div class="main-object-container">
    <div class="constrain constrain--light">
        
    </div>

    <div id="js_listpage_top_objects"  data-group-name="‘9780575104181’ in Alle artikelen"  data-bltgg="sOloXhA4zqCekILbwytNBQ.2_3" >
  
</div>


    <div class="constrain constrain--light">
        
        
        <div class="[ fluid-grid fluid-grid--m  ] clear_fix">

            <div class="[ fluid-grid__item u-1/5@screen-xl-up ] navigation-container">
                <wsp-async-browse-form>
                    <form>
                        <div class="offcanvas-result-status js_result_status js_offcanvas_close">
  <button class="ui-btn  ui-btn--primary ui-btn--block"
     data-test="number-of-articles-button">

  
  Filter 1 resultaat</button>


</div>

<div class="[ offcanvas-container ]" id="js_filter_navigation" data-offcanvas-id="filter_navigation" data-test="facets">

  <div class="offcanvas-wrapper">
    <div class="offcanvas-header">
      <a href="#" class="offcanvas-header__close-btn  js_offcanvas_close">
        <span class="sb sb-cross"></span>
      </a>
    </div>
    
    <div class="js_facets_heading">
      <p class="facet-notification" style="display: none">
  
</p>

<h2 class="facet-controls-header h4">
  Filters
</h2>

    </div>

    <div class="facet-controls js_facet_controls ">
      
        <div class="facet-control js_hidden_control is-hidden" data-test="hiddenselect">
  
    <input type="hidden" id="hidden_facet_1" name="N" value="0">
  
    <input type="hidden" id="hidden_query_1" name="searchtext" value="9780575104181">
  
</div>

      
        

      
        <div data-id="1" class="facet-control category-tree-control" data-test="category_tree">
  <div class="facet-control__header">
    <strong>Categorieën</strong>
  </div>

  

  
    <div class="facet-control__filter facet-control__filter--no-padding">
      <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;N&#x3D;8299" class="px_listpage_categoriesleft_click">Boeken <span class="group_amount mini_details">(1)</span></a>
      
    </div>
  
    <div class="facet-control__filter facet-control__filter--no-padding">
      <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;N&#x3D;40413" class="px_listpage_categoriesleft_click">- Fantasy <span class="group_amount mini_details">(1)</span></a>
      
    </div>
  

  
</div>

      
        <wsp-facet-toggle>
    <input id="collapse_1312" name="collapse" type="checkbox" class="is-hidden" >

        <div data-id="1312" class="facet-control js_multiselect_control " data-test="multiselect">


            <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
                <label role="heading" aria-level="2" tabindex="0" for="collapse_1312" class="h-block js_toggle_facet_heading">
                    <strong class="facet-control__heading--collaps">Taal</strong>
                </label>
            </div>


      <div class="facet-control__collapse">

        <input id="show-more_1312" name="show-more" type="checkbox" class="is-hidden js_toggle_more_less" >

        

          <div class="facet-control__items">
            
              <div class="facet-control__filter   " data-required="false">
                <div class="ui-input-checkbox ui-input-checkbox--9 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="8292" name="refine_8292" value="8292">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><div data-test="facet-value-0">
                  
                  
                  Engels
                  
                      
                          <span class="group_amount mini_details">(1)</span>
                      
                  
                </div></span>
    </div>
  </label>
</div>

              </div>
            
          </div>

          
        

        
      </div>
    </div>
</wsp-facet-toggle>
      
        <wsp-facet-toggle>
    <input id="collapse_273" name="collapse" type="checkbox" class="is-hidden" >

        <div data-id="273" class="facet-control js_multiselect_control " data-test="multiselect">


            <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
                <label role="heading" aria-level="2" tabindex="0" for="collapse_273" class="h-block js_toggle_facet_heading">
                    <strong class="facet-control__heading--collaps">Boek, ebook of luisterboek?</strong>
                </label>
            </div>


      <div class="facet-control__collapse">

        <input id="show-more_273" name="show-more" type="checkbox" class="is-hidden js_toggle_more_less" >

        

          <div class="facet-control__items">
            
              <div class="facet-control__filter   " data-required="false">
                <div class="ui-input-checkbox ui-input-checkbox--9 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="7419" name="refine_7419" value="7419">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><div data-test="facet-value-0">
                  
                  
                  Ebook
                  
                      
                          <span class="group_amount mini_details">(1)</span>
                      
                  
                </div></span>
    </div>
  </label>
</div>

              </div>
            
          </div>

          
        

        
      </div>
    </div>
</wsp-facet-toggle>
      
        <wsp-facet-toggle>
    <input id="collapse_12194" name="collapse" type="checkbox" class="is-hidden" >

    <div data-id="12194" class="facet-control js_price_slider_control" data-test="slider">
        <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
            <label role="heading" aria-level="2" tabindex="0" for="collapse_12194" class="h-block js_toggle_facet_heading">
                <strong class="facet-control__heading--collaps" data-test="facet-title">Prijs (€)</strong>
            </label>
        </div>

      <div class="facet-control__collapse">
        

        <wsp-async-browse-slider variant="price">
          <input type="hidden" name="12194" id="facet_12194" value="5-5"
            class="js_async_browse_slider_input">

          <wsp-multi-range class="multi-range" selected-min="5" selected-max="5" min="5"
            max="5" min-input-target="multiple-range_12194_selected-min"
            max-input-target="multiple-range_12194_selected-max" value-distribution="boundary">
            <div class="multi-range__inputs" data-test="slider-input">
              
                <div class="ui-input-text ui-input-text--small" autocomplete="off" data-test="minimum">
  <div class="ui-input-text__wrapper">
      <input
        class="ui-input-text__control"
        placeholder=" " id="multiple-range_12194_selected-min" name="multiple-range_12194_selected-min" value="5" maxlength="5" aria-label="minimum"/>

    <label class="ui-input-text__label" ></label>

    
    <span class="ui-input-text__meta"></span>

    </div>
</div>

                -
                <div class="ui-input-text ui-input-text--small" autocomplete="off" data-test="maximum">
  <div class="ui-input-text__wrapper">
      <input
        class="ui-input-text__control"
        placeholder=" " id="multiple-range_12194_selected-max" name="multiple-range_12194_selected-max" value="5" maxlength="5" aria-label="maximum"/>

    <label class="ui-input-text__label" ></label>

    
    <span class="ui-input-text__meta"></span>

    </div>
</div>

              

              <button class="ui-btn  ui-btn--secondary ui-btn--square u-pl--xxs" aria-label="update"
     type="slider">

  <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-next bi-lg" data-test="icon-chevron-next" focusable="false"><path fill-rule="evenodd" d="M4.694 3.644a.508.508 0 010-.71.497.497 0 01.7-.006l4.06 4.084-4.055 4.054a.505.505 0 01-.707 0 .501.501 0 01.002-.71l3.352-3.351-3.352-3.36z"/></svg>
  </button>

            </div>

            

            <div class="multi-range__slider-wrapper js_slider_wrap">
              <div class="multi-range__slider" data-test="slider-bar">
                <div class="multi-range__track">
                  <div class="multi-range__bar js_bar"></div>
                  <div class="multi-range-handle multi-range-handle--point-right js_handle_min"
                    data-test="slider-first-handle">
                    <div class="multi-range-handle__input">
                      <div class="multi-range-handle__direction">
                      </div>
                    </div>
                  </div>
                  <div class="multi-range-handle multi-range-handle--point-left js_handle_max"
                    data-test="slider-second-handle">
                    <div class="multi-range-handle__input">
                      <div class="multi-range-handle__direction">
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </wsp-multi-range>

        </wsp-async-browse-slider>

        <p style="display: none" class="multi-range__error-message js_error_message"></p>
      </div>
    </div>
</wsp-facet-toggle>

      
        <wsp-facet-toggle>
    <input id="collapse_14032" name="collapse" type="checkbox" class="is-hidden" >

        <div data-id="14032" class="facet-control js_multiselect_control " data-test="multiselect">


            <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
                <label role="heading" aria-level="2" tabindex="0" for="collapse_14032" class="h-block js_toggle_facet_heading">
                    <strong class="facet-control__heading--collaps">Conditie</strong>
                </label>
            </div>


      <div class="facet-control__collapse">

        <input id="show-more_14032" name="show-more" type="checkbox" class="is-hidden js_toggle_more_less" >

        

          <div class="facet-control__items">
            
              <div class="facet-control__filter   " data-required="false">
                <div class="ui-input-checkbox ui-input-checkbox--9 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="14033" name="refine_14033" value="14033">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><div data-test="facet-value-0">
                  
                  
                  Nieuw
                  
                      
                          <span class="group_amount mini_details">(1)</span>
                      
                  
                </div></span>
    </div>
  </label>
</div>

              </div>
            
          </div>

          
        

        
      </div>
    </div>
</wsp-facet-toggle>
      
        <wsp-facet-toggle>
    <input id="collapse_16699" name="collapse" type="checkbox" class="is-hidden" >

    <div data-id="16699" class="facet-control rating-control js_rating_control" data-test="rating">
        <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
            <label role="heading" aria-level="2" tabindex="0" for="collapse_16699" class="h-block js_toggle_facet_heading">
                <strong class="facet-control__heading--collaps">Minimale beoordeling</strong>
            </label>
        </div>

      <div class="facet-control__collapse">
        
          <div class="facet-control__filter facet-control__filter--selected" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_all" name="rating" value="all" checked="checked" autocomplete="off" />
              <label for="facet_all">
                <span class="rating-stars rating-stars--no-rating">
                  <span class="rating-stars__value"></span>
                </span>
                
                
              </label>
            </span>
          </div>
        
          <div class="facet-control__filter" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_1" name="rating" value="1" autocomplete="off" />
              <label for="facet_1">
                <span class="rating-stars rating-stars--1-star">
                  <span class="rating-stars__value"></span>
                </span>
                ster of hoger
                
              </label>
            </span>
          </div>
        
          <div class="facet-control__filter" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_2" name="rating" value="2" autocomplete="off" />
              <label for="facet_2">
                <span class="rating-stars rating-stars--2-star">
                  <span class="rating-stars__value"></span>
                </span>
                ster of hoger
                
              </label>
            </span>
          </div>
        
          <div class="facet-control__filter" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_3" name="rating" value="3" autocomplete="off" />
              <label for="facet_3">
                <span class="rating-stars rating-stars--3-star">
                  <span class="rating-stars__value"></span>
                </span>
                ster of hoger
                
              </label>
            </span>
          </div>
        
          <div class="facet-control__filter" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_4" name="rating" value="4" autocomplete="off" />
              <label for="facet_4">
                <span class="rating-stars rating-stars--4-star">
                  <span class="rating-stars__value"></span>
                </span>
                ster of hoger
                
              </label>
            </span>
          </div>
        
          <div class="facet-control__filter" data-required="false">
            <span class="radio-input">
              <input type="radio" id="facet_5" name="rating" value="5" autocomplete="off" />
              <label for="facet_5">
                <span class="rating-stars rating-stars--5-star">
                  <span class="rating-stars__value"></span>
                </span>
                
                
              </label>
            </span>
          </div>
        
      </div>
    </div>
</wsp-facet-toggle>

      
        <wsp-facet-toggle>
    <input id="collapse_1282" name="collapse" type="checkbox" class="is-hidden" >

        <div data-id="1282" class="facet-control js_multiselect_control " data-test="multiselect">


            <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
                <label role="heading" aria-level="2" tabindex="0" for="collapse_1282" class="h-block js_toggle_facet_heading">
                    <strong class="facet-control__heading--collaps">Beschikbaarheid</strong>
                </label>
            </div>


      <div class="facet-control__collapse">

        <input id="show-more_1282" name="show-more" type="checkbox" class="is-hidden js_toggle_more_less" >

        

          <div class="facet-control__items">
            
              <div class="facet-control__filter   " data-required="false">
                <div class="ui-input-checkbox ui-input-checkbox--9 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="1283" name="refine_1283" value="1283">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><div data-test="facet-value-0">
                  
                  
                  Leverbaar
                  
                      
                          <span class="group_amount mini_details">(1)</span>
                      
                  
                </div></span>
    </div>
  </label>
</div>

              </div>
            
          </div>

          
        

        
      </div>
    </div>
</wsp-facet-toggle>
      
        <wsp-facet-toggle>
    <input id="collapse_34709" name="collapse" type="checkbox" class="is-hidden" >

        <div data-id="34709" class="facet-control js_multiselect_control " data-test="multiselect">


            <div class="facet-control__header facet-control__header--collapsable js_toggle_collapse">
                <label role="heading" aria-level="2" tabindex="0" for="collapse_34709" class="h-block js_toggle_facet_heading">
                    <strong class="facet-control__heading--collaps">Studieboek of algemeen</strong>
                </label>
            </div>


      <div class="facet-control__collapse">

        <input id="show-more_34709" name="show-more" type="checkbox" class="is-hidden js_toggle_more_less" >

        

          <div class="facet-control__items">
            
              <div class="facet-control__filter   " data-required="false">
                <div class="ui-input-checkbox ui-input-checkbox--9 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="67401" name="refine_67401" value="67401">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><div data-test="facet-value-0">
                  
                  
                  Algemene boeken
                  
                      
                          <span class="group_amount mini_details">(1)</span>
                      
                  
                </div></span>
    </div>
  </label>
</div>

              </div>
            
          </div>

          
        

        
      </div>
    </div>
</wsp-facet-toggle>
      
        <div class="facet-control js_not_available_control" data-test="not_available">
  
    <div class="facet-control__filter" data-required="false">
      <div class="ui-input-checkbox ui-input-checkbox--8 ui-input-checkbox--small u-nosp">
  <label>
    <input class="srt" type="checkbox" id="non" name="availability" value="non">
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content">Toon artikelen die niet leverbaar zijn</span>
    </div>
  </label>
</div>

    </div>
  
</div>

      
        <div id="facet_feedback"></div>
      
    </div>

    

<div class="mms-banner">
      
    </div>

  </div>

  <noscript>
    <div class="offcanvas-wrapper">
      <div class="heading">
        <p class="facet-notification" style="display: none">
  
</p>

<h2 class="facet-controls-header h4">
  Filters
</h2>

      </div>

      <div class="facet-controls ">
        
          <div class="u-pb--lg">
  <h4>Categorieën</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;8299"  class="h-block">Boeken</a>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;40413"  class="h-block">- Fantasy</a>
  
</div>
        
          <div class="u-pb--lg">
  <h4>Taal</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;8292"  class="h-block">Engels</a>
  
</div>
        
          <div class="u-pb--lg">
  <h4>Boek, ebook of luisterboek?</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;7419"  class="h-block">Ebook</a>
  
</div>
        
          <div class="u-pb--lg">
  <h4>Conditie</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;14033"  class="h-block">Nieuw</a>
  
</div>
        
          <div class="u-pb--lg">
  <h4>Beschikbaarheid</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;1283"  class="h-block">Leverbaar</a>
  
</div>
        
          <div class="u-pb--lg">
  <h4>Studieboek of algemeen</h4>
  
    <a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;list&amp;filter_N&#x3D;67401"  class="h-block">Algemene boeken</a>
  
</div>
        
      </div>

      
    </div>
  </noscript>

</div>
                    </form>
                </wsp-async-browse-form>
            </div>

            <div class="[ fluid-grid__item u-4/5@screen-xl-up ] loader-control rwd-listview fright" id="js_list_view"
                 data-test="listpage">
                <div class="loader-control__content">
                    <div id="js_listpage_objects" class="u-mb--sm"  data-group-name="‘9780575104181’ in Alle artikelen"  data-bltgg="sOloXhA4zqCekILbwytNBQ.2_4" >
  
    <h1 class="h1 bol_header" data-test="page-title">‘9780575104181’ in Alle artikelen</h1>

  
    <p data-test="page-subTitle"  data-group-name="‘9780575104181’ in Alle artikelen"  data-bltgg="sOloXhA4zqCekILbwytNBQ.2_4_5" >
</p>

  
    
  
</div>


                    

                    <div class="sort-bar">
                        
                            <div class="sort-bar__results">
                                <p class="total-results js_total_results" data-test="number-of-articles">1 resultaat</p>

                            </div>
                        

                            <div class="js_sticky_placeholder sort-bar__filter-button u-mb--sm ">
                                <button class="ui-btn  ui-btn--primary ui-btn--block js_offcanvas_open tst_account-menu_offcanvas"
     data-offcanvas-target="filter_navigation" data-test="filter-menu-button">

  <svg version="1.1" viewBox="0 0 14 14" aria-label="Filter" role="img" class="svg-inline--bi bi-filter bi-lg" data-test="icon-filter" focusable="false"><path fill-rule="evenodd" d="M4.5 8.5a1.5 1.5 0 011.414 1h5.583c.247 0 .452.183.495.412L12 10c0 .276-.233.5-.503.5H5.915a1.5 1.5 0 01-2.83 0h-.58a.51.51 0 01-.497-.412L2 10c0-.276.214-.5.505-.5h.58a1.5 1.5 0 011.415-1zm0 1a.5.5 0 10.491.593L5 9.995l-.007-.085A.5.5 0 004.5 9.5zm5-4a1.5 1.5 0 011.414 1h.581a.51.51 0 01.497.412L12 7c0 .276-.214.5-.505.5h-.58a1.5 1.5 0 01-2.83 0H2.503a.509.509 0 01-.495-.412L2 7c0-.276.233-.5.503-.5h5.583a1.5 1.5 0 011.414-1zm0 1a.5.5 0 10.491.593L10 6.994l-.007-.084A.5.5 0 009.5 6.5zm-5-4a1.5 1.5 0 011.414 1h5.583c.247 0 .452.183.495.412L12 4c0 .276-.233.5-.503.5H5.915a1.5 1.5 0 01-2.83 0h-.58a.51.51 0 01-.497-.412L2 4c0-.276.214-.5.505-.5h.58a1.5 1.5 0 011.415-1zm0 1a.5.5 0 10.491.593L5 3.995l-.007-.085A.5.5 0 004.5 3.5z"/></svg>
  </button>


                            </div>
                        

                        <wsp-async-browse-sort class="sort-bar__sorting" data-test="sort-option">
  <label for="sort" class="small--is-visible--inline u-m--label">Sortering</label>
    <span class="small--is-visible--inline"><wsp-modal-opener  data-async-url="/nl/rnwy/sort-info/popup/?showRelevancySection=true" data-test="sort-info-modal">
    <a class="link-no-decoration link-cursor-pointer" data-modal-size="medium" open-modal="true" data-test="sort-info-modal-icon" data-analytics-id="px_common_popup"
        data-analytics-value="popup:listpage:sortinginfo">
        <span class="link-color-black">
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-badge-info-neg bi-lg link-icon" data-test="icon-badge-info-neg" focusable="false"><path fill-rule="evenodd" d="M7 14A7 7 0 117 0a7 7 0 010 14zm-.033-1a6 6 0 100-12 6 6 0 000 12zm-.464-6.984a.512.512 0 01.509-.516h-.024c.283 0 .514.228.515.516l.027 3.968a.512.512 0 01-.509.516h.024a.516.516 0 01-.515-.516l-.027-3.968zM7.031 4.5a.5.5 0 110-1 .5.5 0 010 1z"/></svg>
        </span>
    </a>
</wsp-modal-opener></span>
  <div class="[ form-select  form-select--medium ]">
    <select id="sort" name="sort">
      
        <option value="relevance1" selected="selected">Relevantie</option>
      
        <option value="popularity1" >Populariteit</option>
      
        <option value="price0" >Prijs laag - hoog</option>
      
        <option value="price1" >Prijs hoog - laag</option>
      
        <option value="release_date1" >Verschijningsdatum</option>
      
        <option value="rating1" >Beoordeling</option>
      
        <option value="wishListRank1" >Meest gewild</option>
      
    </select>
  </div>
    <span class="small--is-hidden"><wsp-modal-opener  data-async-url="/nl/rnwy/sort-info/popup/?showRelevancySection=true" data-test="sort-info-modal">
    <a class="link-no-decoration link-cursor-pointer" data-modal-size="medium" open-modal="true" data-test="sort-info-modal-icon" data-analytics-id="px_common_popup"
        data-analytics-value="popup:listpage:sortinginfo">
        <span class="link-color-black">
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-badge-info-neg bi-lg link-icon" data-test="icon-badge-info-neg" focusable="false"><path fill-rule="evenodd" d="M7 14A7 7 0 117 0a7 7 0 010 14zm-.033-1a6 6 0 100-12 6 6 0 000 12zm-.464-6.984a.512.512 0 01.509-.516h-.024c.283 0 .514.228.515.516l.027 3.968a.512.512 0 01-.509.516h.024a.516.516 0 01-.515-.516l-.027-3.968zM7.031 4.5a.5.5 0 110-1 .5.5 0 010 1z"/></svg>
        </span>
    </a>
</wsp-modal-opener></span>
</wsp-async-browse-sort>


                        
                            <div class="sort-bar__view-modes" id="js_view_mode_control">

                                <h2 class="srt" aria-level="2">Deze pagina is ingesteld op tegelweergave. Om de weergave van deze pagina te wijzigen van tegelweergave naar lijstweergave, klik deze button.</h2>
<a href="/nl/nl/s/?page&#x3D;1&amp;searchtext&#x3D;9780575104181&amp;view&#x3D;tile" class="viewmode tile js_view_mode" data-view-mode="tile" data-analytics-id="px_listpage_change_viewmode">
  <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-tileview" data-test="icon-tileview" focusable="false"><path fill-rule="evenodd" d="M1.5 1.629h5v5h-5v-5zm1 1v3h3v-3h-3zm5-1h5v5h-5v-5zm1 1v3h3v-3h-3zm-1 5h5v5h-5v-5zm1 1v3h3v-3h-3zm-7-1h5v5h-5v-5zm1 1v3h3v-3h-3z"/></svg>
</a>

                            </div>
                        

                    </div>

                    
                        <div id="js_active_facet_values"></div>

                    

                    <div class="results-area"  data-group-name="‘9780575104181’ in Alle artikelen"  data-bltgg="sOloXhA4zqCekILbwytNBQ.2" >
                        
<div  data-group-name="‘9780575104181’ in Alle artikelen"  data-bltgg="sOloXhA4zqCekILbwytNBQ.2_6" >
    <ul class="list-view product-list js_multiple_basket_buttons_page   " id="js_items_content" data-test="products">
        
        <li class="product-item--row js_item_root" data-id="9200000011781470"  data-measurementname="webshop_productitem-view_vmqee"  data-bltgi="sOloXhA4zqCekILbwytNBQ.2_6.7" >

  <div class="product-item__image hit-area " data-test="product-image">
    <div class="h-o-hidden ">
      <wsp-analytics-tracking-event
        data-target=".list_page_product_tracking_target"
        data-trigger="click"
        data-config='{"product_id": "9200000011781470"}'
        data-type="select-item"
      >
      <a href="/nl/nl/p/the-reckoners-steelheart/9200000011781470/" class="product-image
      product-image--list
      
       px_list_page_product_click list_page_product_tracking_target" data-list-page-product-click-location="image"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.ProductImage" >
          
            
              <img src="https://media.s-bol.com/ZzZv2plpzErw/135x210.jpg" alt="The Reckoners - Steelheart">
            
          
        </a>
      </wsp-analytics-tracking-event>
    </div>

    

    <div class="">
        <wsp-comparable class="item js-comparable-item"
                data-product-id="9200000011781470"
                data-max-number-of-products="4"
                 data-chunk-id="80007266" >
  <div class="compare-link-layout" >
    <div class="ui-input-checkbox ui-input-checkbox--small u-mr u-mb--0" data-test="compare-checkbox" id="checkbox-9200000011781470">
  <label>
    <input class="srt" type="checkbox" data-analytics-id="px_compare_widget_comparable_click"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.CompareCheckbox" >
    <div class="ui-input-checkbox__label">
      <div class="ui-input-checkbox__control" aria-hidden="true"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check ui-input__icon--check" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg></div>
      <span data-test="input-checkbox-content"><span class="js-comparable-label-on">Vergelijk met andere artikelen</span>
<span class="js-comparable-label-off hide">Je hebt al het maximum aantal artikelen in je lijst staan</span>
</span>
    </div>
  </label>
</div>

    <div class="compare-link js-compare-link links-inline">
      <a aria-label="Naar de gekozen artikelen om te vergelijken">Bekijk nu</a>
    </div>
  </div>
</wsp-comparable>

    </div>
    
        <div class="u-show-block@screen-small-only">
            <wsp-wishlist-button global-id="9200000011781470" wishlist-id="" csrf-token="bccfe77f-266f-4e57-9e42-ee5ab32e9d81"
                     data-size-pick-prevent="wishlist" new float add="Klik hier om het product toe te voegen aan je verlanglijstje"
                     saved="Dit product staat opgeslagen op je verlanglijstje" data-test="btn-wishlist">

    <button class="ui-btn ui-btn--favorite ui-btn--floating"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.AddToWishlist" >
        <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-heart-default bi-uni bi-hide" data-test="icon-heart-default" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M14.4612 2.49501C15.2499 2.16821 16.0953 2 16.949 2C17.8027 2 18.6481 2.16821 19.4368 2.49501C20.2254 2.82176 20.9419 3.30064 21.5454 3.90431C22.149 4.50779 22.6282 5.2246 22.955 6.01317C23.2818 6.80189 23.45 7.64726 23.45 8.50101C23.45 9.35475 23.2818 10.2001 22.955 10.9888C22.7593 11.4612 22.4407 11.8981 22.1441 12.2526C21.8405 12.6157 21.5108 12.9493 21.2543 13.2057C21.2542 13.2057 21.2543 13.2056 21.2543 13.2057L12.7555 21.7044C12.568 21.892 12.3137 21.9973 12.0484 21.9973C11.7832 21.9973 11.5289 21.892 11.3413 21.7044L2.8426 13.2057C1.73995 12.103 0.550003 10.3536 0.550003 8.50101C0.550003 6.77697 1.23487 5.12355 2.45395 3.90448C3.67302 2.6854 5.32644 2.00053 7.05047 2.00053C8.77451 2.00053 10.4279 2.6854 11.647 3.90448L11.9997 4.25721L12.3523 3.90464C12.9558 3.30082 13.6724 2.82182 14.4612 2.49501ZM16.949 4C16.358 4 15.7728 4.11645 15.2267 4.34268C14.6807 4.56892 14.1847 4.90051 13.7669 5.31853L12.7068 6.37853C12.3163 6.76906 11.6832 6.76906 11.2926 6.37853L10.2328 5.31869C9.38878 4.47469 8.24407 4.00053 7.05047 4.00053C5.85687 4.00053 4.71216 4.47469 3.86816 5.31869C3.02416 6.16269 2.55 7.30741 2.55 8.50101C2.55 9.56601 3.29639 10.8311 4.25681 11.7915L12.0484 19.5831L19.8401 11.7915C20.0945 11.5372 20.3686 11.2581 20.61 10.9695C20.8586 10.6723 21.0265 10.4183 21.1073 10.2233C21.3336 9.67725 21.45 9.09202 21.45 8.50101C21.45 7.90999 21.3336 7.32476 21.1073 6.77876C20.8811 6.23275 20.5495 5.73667 20.1315 5.31886C19.7137 4.90085 19.2173 4.56892 18.6712 4.34268C18.1252 4.11644 17.54 4 16.949 4Z" fill="currentColor"/></svg>
        <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-heart-active bi-uni bi-show" data-test="icon-heart-active" focusable="false"><path d="M16.949 2C16.0952 2 15.2499 2.16821 14.4611 2.49501C13.6724 2.82182 12.9558 3.30082 12.3523 3.90464L11.9997 4.25721L11.647 3.90448C10.4279 2.6854 8.77449 2.00053 7.05046 2.00053C5.32643 2.00053 3.67301 2.6854 2.45393 3.90448C1.23486 5.12355 0.549988 6.77697 0.549988 8.50101C0.549988 10.3536 1.73993 12.103 2.84258 13.2057L11.3413 21.7044C11.5289 21.892 11.7832 21.9973 12.0484 21.9973C12.3136 21.9973 12.568 21.892 12.7555 21.7044L21.2543 13.2057C21.5108 12.9493 21.8404 12.6157 22.1441 12.2526C22.4407 11.8981 22.7593 11.4612 22.955 10.9888C23.2818 10.2001 23.45 9.35475 23.45 8.50101C23.45 7.64726 23.2818 6.80189 22.955 6.01317C22.6282 5.2246 22.149 4.50779 21.5453 3.90431C20.9419 3.30064 20.2254 2.82176 19.4368 2.49501C18.6481 2.16821 17.8027 2 16.949 2Z" fill="currentColor"/></svg>
    </button>

</wsp-wishlist-button>
        </div>
    
  </div>

  <div class="product-item__content" data-test="product-content">
    <div class="product-item__info hit-area">
        
        <div data-test="list-page-sustainable-label" class="hit-area__clickable">
    <a href="/nl/nl/sf/duurzaamheid/"
       data-analytics-id="px_list-page_header_meta_click"
       data-analytics-tag="sustainable-label"
       data-test="sustainable-label"
       aria-label="Goede Keuze"
        data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.Sustainable" >
      <div class="sustainable-label" data-test="sustainable-label">
  
    <span class="sustainable-label__icon" data-test="icon"><svg version="1.1" viewBox="0 0 20 20" aria-hidden="true" class="svg-inline--bi bi-graphic-sustainable-globe bi-lg" data-test="icon-graphic-sustainable-globe" focusable="false"><g clip-path="url(#clip0_289_726)"><mask id="mask0_289_726" style="mask-type:alpha" width="20" height="20" x="0" y="0" maskUnits="userSpaceOnUse"><path fill="#fff" d="M20 10c0 5.523-4.477 10-10 10S0 15.523 0 10 4.477 0 10 0s10 4.477 10 10z"/></mask><g mask="url(#mask0_289_726)"><path fill="#1EADFF" d="M20 10c0 5.523-4.477 10-10 10S0 15.523 0 10 4.477 0 10 0s10 4.477 10 10z"/><path fill="#71E9B4" d="M4.5 3.05V1.3C5.94.47 9.45-.409 9.45-.409V.1C9.45 1.7 8 1.85 7.6 1.85c-.453 0-.8.15-.8 1.1 0 1.68-2.3 1.5-2.3.1zM10.7 4.5c-1.9.578-2.112-.562-1.6-1.65.4-.85 1.55-1.9 2.6-2.7.15-.126.585-.455.7-.55l5.5 2.25 2.85 7.4c-.557.223-1.136.061-1.75.25-1.95.6-3.15 4.457-5.15 6.75-1.596 1.83-2.25.35-2.7-.75-.45-1.1 0-4.5-2.6-4.9-2.3-.5-1.736-2.857-.5-2.9 1.236-.043 3.423-.055 4.25-.95.6-.65.7-2.95-1.6-2.25z"/></g><path fill="#1EADFF" d="M10 20c5.523 0 10-4.477 10-10S15.523 0 10 0 0 4.477 0 10s4.477 10 10 10z"/><path fill="#71E9B4" d="M11.736.15C10.523.914 9.2 2.063 8.823 3.58a.855.855 0 00.59 1.025c.2.058.49.193.795.047.8-.381 2.223-.59 2.362.487.138 1.076.243 1.77-1.84 2.222-2.084.45-3.612-.105-3.75 1.493-.132 1.507 1.354 1.493 2.048 1.91.485.29 1.355.277 1.795 3.064.374 2.362.88 3.324 2.025 3.324.52 0 3.02-3.715 4.062-5.695.923-1.753 2.091-2.225 3.07-2.093C19.69 4.72 16.23.936 11.736.15zM5.625 4.166c1.771-.173.853-1.743 1.347-2.139.444-.355 2.51-.055 2.442-2.01a9.95 9.95 0 00-5.016 1.698c-.035 1.111.127 2.559 1.227 2.45v.001z"/></g><defs><clipPath id="clip0_289_726"><path fill="#fff" d="M0 0h20v20H0z"/></clipPath></defs></svg></span>
  
  <span class="sustainable-label__label-text" data-test="label-text">Goede Keuze</span>
</div>

    </a>
</div>
      <ul class="product-creator">
        
          <li>
            <a data-test="party-link" href="/nl/nl/b/brandon-sanderson/2460483/" itemprop="name" data-no-translate="true">Brandon Sanderson</a>
          </li>
        
      </ul>

      <div class="product-title--inline ">
        <wsp-analytics-tracking-event
          data-target=".list_page_product_tracking_target"
          data-trigger="click"
          data-config='{"product_id": "9200000011781470"}'
          data-type="select-item"
        >
          <a role="heading" aria-level="2" class="product-title px_list_page_product_click list_page_product_tracking_target u-mr--xs" href="/nl/nl/p/the-reckoners-steelheart/9200000011781470/" data-test="product-title" data-no-translate="true" data-list-page-product-click-location="title"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.ProductTitle" >The Reckoners - Steelheart</a>
        </wsp-analytics-tracking-event>
      </div>
      
        <span class="labels label_product-type" data-test="product-type-label">Ebook</span>
      

      

      
        <ul class="product-small-specs "
    data-test="product-specs">
  
      
        <li>
          
            <span>Engels</span>
          

          
        </li>
      
  
      
        <li>
          
            <span>E-book</span>
          

          
        </li>
      
  
      
        <li>
          
            <span>9780575104181</span>
          

          
        </li>
      
  
      
        <li>
          
            <span>26 september 2013</span>
          

          
        </li>
      
  
      
        <li>
          
            <span>-6218</span>
          

          
        </li>
      
  
</ul>
      

      
        <div>
          <div class="u-mb--xs" aria-label="Gemiddeld 4.8 van de 5 sterren uit 16 reviews">
            <div class="star-rating" data-count="16" title="klantbeoordeling: 4,8 van de 5 sterren" aria-hidden="true" data-test="rating-stars">
            <span style="width: 96%"></span>
            </div>
          </div>
        </div>
      

      
            
              <p class="medium--is-visible" data-test="product-description">
                Ten years ago, Calamity came. It was a burst in the sky that gave ordinary men and women extraordinary&hellip;
              </p>
            
        
      

      

      <wsp-analytics-tracking-event
        data-target=".list_page_product_tracking_target"
        data-trigger="click"
        data-config='{"product_id": "9200000011781470"}'
        data-type="select-item"
      >
        <a href="/nl/nl/p/the-reckoners-steelheart/9200000011781470/" class="hit-area__link medium--is-hidden list_page_product_tracking_target"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.ProductPageUrl" ></a>
      </wsp-analytics-tracking-event>
    </div>

    <wsp-buy-block class="product-item__options hit-area" data-test="product-options">
      

      

      

    


      

      
        

        <div class="product-prices ">
          <section class="price-block
    
    
    " data-test="priceBlock"
    >
    
        <div class="price-block__price" data-test="priceBlockPrice">
            <div class="price-block__highlight
                
                "
                
                 data-test="priceWithoutDiscount"
            >
                
                <div class="srt" role="heading" aria-level="2">Prijsinformatie en bestellen</div>
<span class="srt" data-test="price-info-srt-text">De prijs van dit product is 5 euro en 49 cent.</span>
<span aria-hidden="true" class="promo-price" data-test="price">5
  <sup class="promo-price__fraction" data-test="price-fraction">49</sup>
</span>
<meta itemprop="price" content="5.49">


            </div>
        </div>
    

    
</section>


          

          

          
        </div>

        

        <div class="product-delivery" data-test="delivery-notification">
            

            Direct beschikbaar

            
            <wsp-tooltip data-test="tooltip">
                <span class="is-hidden js_tooltip_content" data-test="delivery-notification-info">
                    <p><strong>Levertijd</strong></p>
<p>We doen er alles aan om dit artikel op tijd te bezorgen. Het is echter in een enkel geval mogelijk dat door omstandigheden de bezorging vertraagd is.</p>
<p><strong>Bezorgopties</strong></p>
<p>We bieden verschillende opties aan voor het bezorgen of ophalen van je bestelling. Welke opties voor jouw bestelling beschikbaar zijn, zie je bij het afronden van de bestelling.</p>

                </span>
                <span class="medium--is-visible--inline-block" data-ref="#delivery-tooltip-content-9200000011781470" >
                    <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-badge-info-neg bi-lg" data-test="icon-badge-info-neg" focusable="false"><path fill-rule="evenodd" d="M7 14A7 7 0 117 0a7 7 0 010 14zm-.033-1a6 6 0 100-12 6 6 0 000 12zm-.464-6.984a.512.512 0 01.509-.516h-.024c.283 0 .514.228.515.516l.027 3.968a.512.512 0 01-.509.516h.024a.516.516 0 01-.515-.516l-.027-3.968zM7.031 4.5a.5.5 0 110-1 .5.5 0 010 1z"/></svg>
                </span>
            </wsp-tooltip>
            
        </div>
            <div class="product-seller
 medium--is-visible

"
data-test="plazaseller-link">
  
    Verkoop door bol
  
</div>



        <div class="medium--is-visible--flex js_product_actions">
          

          

          <div class="u-pb--xs h-flex">
            
                <a href="/nl/order/basket/addItems.html?productId&#x3D;9200000011781470&amp;offerId&#x3D;0&amp;offerUid&#x3D;6c17c7b1-2b7e-4cc3-b7a5-b8ab5a11bae6&amp;quantity&#x3D;1" role="button" class="ui-btn  ui-btn--buy ui-btn--block@screen-small js_floating_basket_btn js_btn_buy js_preventable_buy_action"
      data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.AddToCart" 
 data-product-id="9200000011781470" data-offer-uid="6c17c7b1-2b7e-4cc3-b7a5-b8ab5a11bae6" data-show-additional-page="true" data-buy-block-type="ListPage" data-test="add-to-basket" data-offer-id="0" data-relatedproductimage_id="9200000011781470" id="9200000011781470" data-button-type="buy" data-size-pick-prevent="add_to_basket"><svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-shopping-cart bi-uni bi-hide" data-test="icon-shopping-cart" focusable="false"><g clip-path="url(#clip0_2737_4226)"><path fill-rule="evenodd" clip-rule="evenodd" d="M0 1C0 0.447715 0.447715 0 1 0H5C5.47663 0 5.88701 0.336385 5.98055 0.803743L6.82043 5H23C23.298 5 23.5805 5.13293 23.7705 5.36256C23.9605 5.59218 24.0381 5.89458 23.9823 6.18733L22.3809 14.5848C22.2437 15.2754 21.868 15.8958 21.3195 16.3373C20.7738 16.7766 20.0916 17.011 19.3914 17H9.68864C8.98837 17.011 8.3062 16.7766 7.76048 16.3373C7.21225 15.8959 6.83664 15.2759 6.69933 14.5857C6.69927 14.5854 6.69939 14.586 6.69933 14.5857L5.02879 6.2392C5.02201 6.21159 5.01638 6.18353 5.01195 6.15508L4.18032 2H1C0.447715 2 0 1.55228 0 1ZM7.22073 7L8.66084 14.1952C8.70656 14.4254 8.83179 14.6322 9.01461 14.7793C9.19743 14.9265 9.42619 15.0047 9.66084 15.0002L9.68 15H19.4L19.4192 15.0002C19.6538 15.0047 19.8826 14.9265 20.0654 14.7793C20.2474 14.6329 20.3723 14.4273 20.4185 14.1984L21.7913 7H7.22073ZM7.00003 21C7.00003 19.8955 7.89546 19 9.00003 19C10.1046 19 11 19.8955 11 21C11 22.1046 10.1046 23 9.00003 23C7.89546 23 7.00003 22.1046 7.00003 21ZM18 21C18 19.8955 18.8954 19 20 19C21.1046 19 22 19.8955 22 21C22 22.1046 21.1046 23 20 23C18.8954 23 18 22.1046 18 21Z" fill="currentColor"/></g><defs><clipPath id="clip0_2737_4226"><rect width="24" height="24" fill="currentColor"/></clipPath></defs></svg>
<svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-check bi-uni bi-show" data-test="icon-check" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M20.7071 5.29289C21.0976 5.68342 21.0976 6.31658 20.7071 6.70711L9.70711 17.7071C9.31658 18.0976 8.68342 18.0976 8.29289 17.7071L3.29289 12.7071C2.90237 12.3166 2.90237 11.6834 3.29289 11.2929C3.68342 10.9024 4.31658 10.9024 4.70711 11.2929L9 15.5858L19.2929 5.29289C19.6834 4.90237 20.3166 4.90237 20.7071 5.29289Z" fill="currentColor"/></svg>
Ebook</a>

            
            <wsp-wishlist-button global-id="9200000011781470" wishlist-id="" csrf-token="bccfe77f-266f-4e57-9e42-ee5ab32e9d81"
                     data-size-pick-prevent="wishlist" new add="Klik hier om het product toe te voegen aan je verlanglijstje"
                     saved="Dit product staat opgeslagen op je verlanglijstje" data-test="btn-wishlist"><button class="ui-btn  ui-btn--favorite ui-btn--square" aria-label="Klik hier om het product toe te voegen aan je verlanglijstje"
      data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.AddToWishlist" 
 type="button">

  <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-heart-default bi-uni bi-hide" data-test="icon-heart-default" focusable="false"><path fill-rule="evenodd" clip-rule="evenodd" d="M14.4612 2.49501C15.2499 2.16821 16.0953 2 16.949 2C17.8027 2 18.6481 2.16821 19.4368 2.49501C20.2254 2.82176 20.9419 3.30064 21.5454 3.90431C22.149 4.50779 22.6282 5.2246 22.955 6.01317C23.2818 6.80189 23.45 7.64726 23.45 8.50101C23.45 9.35475 23.2818 10.2001 22.955 10.9888C22.7593 11.4612 22.4407 11.8981 22.1441 12.2526C21.8405 12.6157 21.5108 12.9493 21.2543 13.2057C21.2542 13.2057 21.2543 13.2056 21.2543 13.2057L12.7555 21.7044C12.568 21.892 12.3137 21.9973 12.0484 21.9973C11.7832 21.9973 11.5289 21.892 11.3413 21.7044L2.8426 13.2057C1.73995 12.103 0.550003 10.3536 0.550003 8.50101C0.550003 6.77697 1.23487 5.12355 2.45395 3.90448C3.67302 2.6854 5.32644 2.00053 7.05047 2.00053C8.77451 2.00053 10.4279 2.6854 11.647 3.90448L11.9997 4.25721L12.3523 3.90464C12.9558 3.30082 13.6724 2.82182 14.4612 2.49501ZM16.949 4C16.358 4 15.7728 4.11645 15.2267 4.34268C14.6807 4.56892 14.1847 4.90051 13.7669 5.31853L12.7068 6.37853C12.3163 6.76906 11.6832 6.76906 11.2926 6.37853L10.2328 5.31869C9.38878 4.47469 8.24407 4.00053 7.05047 4.00053C5.85687 4.00053 4.71216 4.47469 3.86816 5.31869C3.02416 6.16269 2.55 7.30741 2.55 8.50101C2.55 9.56601 3.29639 10.8311 4.25681 11.7915L12.0484 19.5831L19.8401 11.7915C20.0945 11.5372 20.3686 11.2581 20.61 10.9695C20.8586 10.6723 21.0265 10.4183 21.1073 10.2233C21.3336 9.67725 21.45 9.09202 21.45 8.50101C21.45 7.90999 21.3336 7.32476 21.1073 6.77876C20.8811 6.23275 20.5495 5.73667 20.1315 5.31886C19.7137 4.90085 19.2173 4.56892 18.6712 4.34268C18.1252 4.11644 17.54 4 16.949 4Z" fill="currentColor"/></svg>

  <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-heart-active bi-uni bi-show" data-test="icon-heart-active" focusable="false"><path d="M16.949 2C16.0952 2 15.2499 2.16821 14.4611 2.49501C13.6724 2.82182 12.9558 3.30082 12.3523 3.90464L11.9997 4.25721L11.647 3.90448C10.4279 2.6854 8.77449 2.00053 7.05046 2.00053C5.32643 2.00053 3.67301 2.6854 2.45393 3.90448C1.23486 5.12355 0.549988 6.77697 0.549988 8.50101C0.549988 10.3536 1.73993 12.103 2.84258 13.2057L11.3413 21.7044C11.5289 21.892 11.7832 21.9973 12.0484 21.9973C12.3136 21.9973 12.568 21.892 12.7555 21.7044L21.2543 13.2057C21.5108 12.9493 21.8404 12.6157 22.1441 12.2526C22.4407 11.8981 22.7593 11.4612 22.955 10.9888C23.2818 10.2001 23.45 9.35475 23.45 8.50101C23.45 7.64726 23.2818 6.80189 22.955 6.01317C22.6282 5.2246 22.149 4.50779 21.5453 3.90431C20.9419 3.30064 20.2254 2.82176 19.4368 2.49501C18.6481 2.16821 17.8027 2 16.949 2Z" fill="currentColor"/></svg>
</button>
</wsp-wishlist-button>

          </div>
        </div>
      

      
        

        
          <p class="small_details product-prices medium--is-visible" data-test="other-format-link">
            <a href="/nl/nl/p/the-reckoners-1-steelheart/9200000033206563/" title="The Reckoners 1 - Steelheart">Ebook</a> <span>7,99</span>
          </p>
        
      

      

      

      
        <wsp-analytics-tracking-event
          data-target=".list_page_product_tracking_target"
          data-trigger="click"
          data-config='{"product_id": "9200000011781470"}'
          data-type="select-item"
        >
          <a href="/nl/nl/p/the-reckoners-steelheart/9200000011781470/" class="hit-area__link medium--is-hidden list_page_product_tracking_target"  data-bltgh="sOloXhA4zqCekILbwytNBQ.2_6.7.ProductPageUrl" ></a>
        </wsp-analytics-tracking-event>

    </wsp-buy-block>

  </div>
</li>

        
    </ul>
</div>




                    </div>

                    <div id="js_pagination_control" class="u-pb--sm">
                        
                    </div>

                </div>
            </div>

            <div class="css-loader-container">
                <div class="css-loader"></div>
            </div>

        </div>

        <wsp-async-browse-page-info page-info="{&quot;asyncSearch&quot;:false,&quot;countryCode&quot;:&quot;NL&quot;,&quot;creatorId&quot;:&quot;&quot;,&quot;filterN&quot;:[],&quot;filterNf&quot;:[],&quot;headerVariant&quot;:&quot;&quot;,&quot;n&quot;:&quot;0&quot;,&quot;page&quot;:&quot;1&quot;,&quot;pageType&quot;:&quot;l&quot;,&quot;section&quot;:&quot;&quot;,&quot;slug&quot;:&quot;&quot;,&quot;tracking&quot;:{&quot;prop11&quot;:&quot;1&quot;,&quot;eVar14&quot;:&quot;typed search&quot;,&quot;eVar15&quot;:&quot;1&quot;,&quot;eVar19&quot;:&quot;1&quot;,&quot;eVar20&quot;:&quot;bestverkocht_11&quot;,&quot;eVar24&quot;:&quot;view:list|standard&quot;,&quot;eVar74&quot;:&quot;heterogeneous&quot;,&quot;eVar75&quot;:&quot;st:0|fb1:0|fb2:0|fb3:0|fb4:0|fb5:0|fb6:0|vce:0|bc:0|rf:0|cf:0|ih:0|se:0|hs:1|af:0|bs:0|tf:0|fr:0|qc:0&quot;,&quot;pageName&quot;:&quot;/nl/nl/s/&quot;},&quot;url&quot;:&quot;/nl/nl/s/?page\u003d1\u0026searchtext\u003d9780575104181\u0026view\u003dlist&quot;,&quot;viewMode&quot;:&quot;&quot;,&quot;tileMode&quot;:&quot;&quot;,&quot;showExplicitContent&quot;:false,&quot;dmpTaxonomy&quot;:&quot;{\&quot;breadcrumbs\&quot;:[],\&quot;searchTerms\&quot;:\&quot;9780575104181\&quot;,\&quot;filters\&quot;:{},\&quot;productInfo\&quot;:[{\&quot;productId\&quot;:\&quot;9200000011781470\&quot;,\&quot;ean\&quot;:\&quot;9780575104181\&quot;,\&quot;title\&quot;:\&quot;The Reckoners - Steelheart\&quot;,\&quot;price\&quot;:5.49,\&quot;categoryTreeList\&quot;:[[\&quot;Boeken\&quot;],[\&quot;Boeken\&quot;,\&quot;Fantasy \u0026amp; Sciencefiction\&quot;],[\&quot;Boeken\&quot;,\&quot;Thrillers \u0026amp; Spanning\&quot;],[\&quot;Boeken\&quot;,\&quot;Fantasy \u0026amp; Sciencefiction\&quot;,\&quot;Fantasy\&quot;]],\&quot;brick\&quot;:\&quot;10000926\&quot;,\&quot;chunk\&quot;:\&quot;80007266\&quot;,\&quot;brand\&quot;:null,\&quot;publisher\&quot;:\&quot;Gollancz\&quot;,\&quot;author\&quot;:\&quot;Brandon Sanderson\&quot;,\&quot;label\&quot;:null,\&quot;artist\&quot;:null,\&quot;color\&quot;:null,\&quot;size\&quot;:null,\&quot;averageReviewRating\&quot;:4.8,\&quot;priceStars\&quot;:null,\&quot;seriesList\&quot;:[],\&quot;sellerName\&quot;:\&quot;bol.com\&quot;,\&quot;uniqueProductAttribute\&quot;:\&quot;BINDING-E-book\&quot;,\&quot;additionalProductAttribute\&quot;:null}]}&quot;}"></wsp-async-browse-page-info>
    </div>
</div>

  </div>
</wsp-async-browse>


  <div class="constrain">
    <div class="email-opt-in-personalized-banner">
    <div class="email-opt-in-personalized-banner__titles">
        <div class="email-opt-in-personalized-banner__title">Geen voordeel missen, Sebastiaan</div>
        <div class="email-opt-in-personalized-banner__subtitle">Krijg (win)acties, aanbevelingen en voordeel direct in je mail</div>
    </div>
    <div class="email-opt-in-personalized-banner__button u-1/1@screen-small-only" data-test="bannerButton">
        <wsp-modal-opener data-async-url="/nl/rnwy/emailOptIn/modalView/">
            <button class="ui-btn  ui-btn--primary ui-btn--block@screen-small js_email_opt_in_opener"
     open-modal="true">

  <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-mail bi-lg" data-test="icon-mail" focusable="false"><path fill-rule="evenodd" d="M2.992 4H2v6h10V4h-.966L7.37 7.662a.502.502 0 01-.709.002l-.252-.25a.486.486 0 01-.037-.033L2.993 4zm1.414 0l2.607 2.606L9.619 4H4.406zM1.063 3H13v8H1.062V3z"/></svg>
  Ja, graag</button>

        </wsp-modal-opener>
    </div>
</div>

</div>

  <wsp-comparable-container data-product-compare-url="/nl/nl/producten-vergelijken/" data-max-number-of-products="4">
  <a class="compare-overlay js-compare-overlay" data-analytics-id="px_compare_widget_exit_click" aria-label=Minimaliseer de productvergelijker></a>
  <div class="compare-floater">
      <a class="compare-floater-header js-floater-header js-floater-header-expand" data-analytics-id="px_compare_widget_floater_click" aria-expanded="false">
          <div class="srt" role="heading" aria-level="2">Lijst met gekozen artikelen om te vergelijken</div>
          <span class="compare-floater-header--icon js-floater-header-icon">
              <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-up bi-lg" data-test="icon-chevron-up" focusable="false"><path fill-rule="evenodd" d="M3.646 8.306a.508.508 0 01-.71 0 .497.497 0 01-.007-.7l4.085-4.06 4.054 4.055a.505.505 0 010 .707.501.501 0 01-.71-.002L7.008 4.954 3.646 8.306z"/></svg>
          </span>
          <span class="compare-floater-header--title" data-test="compare-title">
              Vergelijk artikelen
          </span>
          <span class="js-floater-header-count"></span>
      </a>
      <ul class="compare-floater-items js-compare-floater-items" data-test="compare-items">
          <li class="compare-floater-explanation js-compare-explanation">
              <span class="js-compare-minimum" data-test="minimum-needed">Waar wil je dit mee vergelijken?
Je kan in totaal vier artikelen kiezen.</span>
              <span class="js-compare-maximum" data-test="maximum-explanation">
                  Er is nog plaats voor
                  <span class="js-explanation-count"></span>
                  <span class="js-explanation-items">andere artikelen.</span>
                  <span class="js-explanation-last-item">ander artikel.</span>
              </span>
          </li>
      </ul>
      <div class="compare-floater-footer js-compare-footer">
          <button class="ui-btn  ui-btn--primary ui-btn--block js-compare-go" aria-label="Naar de gekozen artikelen om te vergelijken"
     data-analytics-id="px_compare_widget_compare_click" data-test="compare-button">

  
  Vergelijk</button>

      </div>
  </div>
</wsp-comparable-container>


    </main>
    
    <footer class="wsp-footer"    data-group-name="footer"  data-bltgg="sOloXhA4zqCekILbwytNBQ.8" >
  
  
  
  <div class="wsp-footer__servicebar" data-test="service-bar">
    
    <div class="wsp-footer__services">
      <div class="wsp-footer__service  wsp-footer__service--title">
        Service &amp; contact
      </div>
      <div class="wsp-footer__service">
        <figure class="wsp-footer__figure">
          <svg version="1.1" viewBox="0 0 48 48" aria-hidden="true" class="svg-inline--bi bi-line-graphic-account bi-3x" data-test="icon-line-graphic-account" focusable="false"><path d="M23.8 48c-6.5 0-14.6-1.5-18.1-5C4.6 41.9 4 40.6 4 39.3c0-8.3 6.2-15.5 14.5-17.7-3.7-1.9-6.2-5.7-6.2-10.2C12.4 5.1 17.5 0 23.8 0s11.5 5.1 11.5 11.5c0 4.4-2.5 8.2-6.2 10.2 8.3 2.2 14.5 9.4 14.5 17.7 0 5.5-10.2 8.6-19.8 8.6zm0-25.1C14.2 22.9 6 30.4 6 39.3c0 .8.4 1.6 1.1 2.3 2.6 2.6 9.4 4.4 16.6 4.4h.1c8.6 0 17.8-2.7 17.8-6.7 0-8.9-8.1-16.4-17.8-16.4zm0-20.9c-5.2 0-9.5 4.2-9.5 9.5s4.2 9.5 9.5 9.5c5.2 0 9.5-4.2 9.5-9.5S29 2 23.8 2z"/></svg>
        </figure>
        <div class="wsp-footer__body">
          <p class="u-nosp"><strong>Snel regelen in je account</strong></p>
          <p>
            <a href="/nl/rnwy/account/bestellingen/overzicht"
               class="px_common_customerservice_footer_click" data-test="service-bar-leveringen">Volg je bestelling</a>,
            <a href="/nl/rnwy/account/facturen"
               class="px_common_customerservice_footer_click" data-test="service-bar-betalen">betaal facturen</a> of
            <a href="/nl/account/retourneren/overzicht"
               class="px_common_customerservice_footer_click" data-test="service-bar-retourneren">retourneer een artikel</a>.
          </p>
        </div>
      </div>
      <div class="wsp-footer__service">
        <figure class="wsp-footer__figure">
          <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-line-graphic-customercare bi-3x" data-test="icon-line-graphic-customercare" focusable="false"><path d="M22.52 14A9.07 9.07 0 106 7.23a8.43 8.43 0 00-3.51 2.13 8.6 8.6 0 00-1 10.92l-.71 2.93 2.93-.73a8.58 8.58 0 0010.93-1 8.62 8.62 0 002.12-3.58 8.5 8.5 0 003.12-1.27l2.74.69a.69.69 0 00.44-.15.4.4 0 00.15-.44zm-8.62 6.8a7.57 7.57 0 01-9.65.89l-.34-.25-1.81.39.44-1.76-.25-.34a7.58 7.58 0 1111.61 1.08zm6-5.19a.38.38 0 00-.4.05 7.74 7.74 0 01-2.5 1.17 8.55 8.55 0 00-2.37-7.47A8.67 8.67 0 007.12 7 7.81 7.81 0 019.2 3.33a8.16 8.16 0 0111.46 0 8.1 8.1 0 01.93 10.29.5.5 0 00-.05.39l.46 2.15zm-5.9-.28a5.2 5.2 0 01-5.29 5.14 5.24 5.24 0 01-5.29-5.14.47.47 0 01.49-.49.47.47 0 01.49.49 4.25 4.25 0 004.31 4.16A4.25 4.25 0 0013 15.33a.5.5 0 011 0z"/></svg>
        </figure>
        <div class="wsp-footer__body">
          <p class="u-nosp"><strong>Heb je ons nodig?</strong></p>
          <p>
            We helpen je graag.
            <a href="/nl/nl/m/klantenservice/" class="px_common_customerservice_footer_click" data-test="service-bar-contact">Onze klantenservice</a>
            is dag en nacht open.
          </p>
        </div>
      </div>
    </div>
    
  </div>
  

  <div class="constrain">
    
    <ul class="wsp-footer__links" data-test="footer-links">
      <li class="wsp-footer__link  px_footer_links_click">
      <h3>
       <a href="/nl/nl/m/klantenservice/">
        Klantenservice
      <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-next bi-lg" data-test="icon-chevron-next" focusable="false"><path fill-rule="evenodd" d="M4.694 3.644a.508.508 0 010-.71.497.497 0 01.7-.006l4.06 4.084-4.055 4.054a.505.505 0 01-.707 0 .501.501 0 01.002-.71l3.352-3.351-3.352-3.36z"/></svg>
       </a>
      </h3>
        <ul>
          <li><a href="/nl/nl/m/klantenservice/">Contact opnemen</a></li>
          <li><a href="/nl/nl/klantenservice/sb/5657512099446784/bestellen">Bestellen &amp; Leveren</a></li>
          <li><a href="/nl/nl/klantenservice/sb/5690774674997248/betalen">Betalen</a></li>
          <li><a href="/nl/nl/klantenservice/sb/5644784165191680/retourneren-annuleren">Retourneren</a></li>
          <li><a href="/nl/nl/klantenservice/sb/5767040140836864/garantie-reparatie">Garantie &amp; Reparatie</a></li>
        </ul>
      </li>

      <li class="wsp-footer__link  px_footer_links_click">
        <h3>
          <a href="//pers.bol.com/nl/het-verhaal/">
            Over bol 
          <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-next bi-lg" data-test="icon-chevron-next" focusable="false"><path fill-rule="evenodd" d="M4.694 3.644a.508.508 0 010-.71.497.497 0 01.7-.006l4.06 4.084-4.055 4.054a.505.505 0 01-.707 0 .501.501 0 01.002-.71l3.352-3.351-3.352-3.36z"/></svg>
          </a>
        </h3>
        <ul>
          <li><a href="/nl/nl/sf/winkelen-zonder-zorgen/">De voordelen van bol</a></li>
          <li><a href="//pers.bol.com/">Nieuws</a></li>
          <li><a href="//banen.bol.com/">Werken bij bol</a></li>
          <li><a href="/nl/nl/sf/duurzaamheid/">Bol &amp; duurzaamheid</a></li>
           <li><a href="/nl/nl/sf/b-corp/">Bol &amp; B corp</a></li>
          <li><a href="/nl/nl/sf/bollebozen/">Bollebozen</a></li>
          <li><a href="/nl/nl/sf/mobiele-app/">De bol app</a></li>
          <li><a href="/nl/nl/klantenservice/a/5715110588841984/bedrijfsgegevens">Bedrijfsgegevens</a></li>
        </ul>
      </li>

      <li class="wsp-footer__link px_footer_links_click">
        
          <a href="/nl/nl/sf/zakendoen-met-bolcom/">
            Zakendoen met bol
            <svg version="1.1" viewBox="0 0 14 14" aria-hidden="true" class="svg-inline--bi bi-chevron-next bi-lg" data-test="icon-chevron-next" focusable="false"><path fill-rule="evenodd" d="M4.694 3.644a.508.508 0 010-.71.497.497 0 01.7-.006l4.06 4.084-4.055 4.054a.505.505 0 01-.707 0 .501.501 0 01.002-.71l3.352-3.351-3.352-3.36z"/></svg>
          </a>
        
        <ul>
          <li><a href="//partnerplatform.bol.com/">Zakelijk verkopen</a></li>
          
            <li><a href="/nl/nl/sf/zakelijk-bestellen/">Zakelijk bestellen</a></li>
            <li><a href="//cadeaukaart.bol.com/">Zakelijke cadeaukaarten</a></li>
          
          <li><a href="//adverteren.bol.com/">Adverteren</a></li>
          
            <li><a href="//affiliate.bol.com/nl/">Affiliate Marketing</a></li>
          
        </ul>
      </li>

      

      <li class="wsp-footer__link wsp-footer__link--social" data-test="footer-social">
        <p><strong>Wil je ons volgen?</strong></p>
        <div class="wsp-footer__social-icons">
          <a href="//www.facebook.com/bolpuntcom" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="facebook">
            <span class="srt" id="sr-text">Volg ons op de facebook pagina van bol</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-facebook bi-lg" data-test="icon-facebook" focusable="false"><path fill-rule="evenodd" d="M23 0H1a1 1 0 00-1 1v22a1 1 0 001 1h11.75v-9h-3v-3.75h3v-3c0-3.1 1.963-4.625 4.728-4.625 1.324 0 2.463.099 2.794.142v3.24l-1.917.001c-1.504 0-1.855.715-1.855 1.763v2.479h3.75L19.5 15h-3l.06 9H23a1 1 0 001-1V1a1 1 0 00-1-1"/></svg>
          </a>
          <a href="//www.pinterest.com/bol" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="pintrest">
            <span class="srt" id="sr-text">Volg ons op de pinterest pagina van bol</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-pinterest bi-lg" data-test="icon-pinterest" focusable="false"><path fill-rule="evenodd" d="M12 0C5.372 0 0 5.372 0 12c0 5.084 3.163 9.426 7.627 11.174-.105-.95-.2-2.406.042-3.442.217-.936 1.407-5.965 1.407-5.965s-.36-.718-.36-1.781c0-1.669.968-2.915 2.172-2.915 1.024 0 1.518.769 1.518 1.69 0 1.03-.655 2.57-.993 3.996-.283 1.195.598 2.169 1.777 2.169 2.133 0 3.772-2.25 3.772-5.495 0-2.873-2.065-4.883-5.013-4.883-3.414 0-5.418 2.562-5.418 5.208 0 1.031.397 2.138.893 2.739.098.119.112.222.083.344-.091.38-.293 1.194-.333 1.361-.053.219-.174.266-.402.16-1.499-.698-2.436-2.888-2.436-4.649 0-3.785 2.75-7.261 7.93-7.261 4.162 0 7.397 2.966 7.397 6.93 0 4.136-2.607 7.464-6.227 7.464-1.216 0-2.358-.632-2.75-1.378l-.748 2.853c-.27 1.042-1.002 2.348-1.492 3.146A11.99 11.99 0 0012 24c6.628 0 12-5.373 12-12 0-6.628-5.372-12-12-12"/></svg>
          </a>
          <a href="//www.instagram.com/bol" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="instagram">
            <span class="srt" id="sr-text">Volg ons op de instagram pagina van bol</span>
            <svg version="1.1" viewBox="0 0 500 500" aria-hidden="true" class="svg-inline--bi bi-instagram bi-lg" data-test="icon-instagram" focusable="false"><path d="M250 45.05c66.75 0 74.66.25 101 1.45 24.38 1.11 37.61 5.19 46.42 8.61a77.52 77.52 0 0128.75 18.7 77.52 77.52 0 0118.7 28.75c3.42 8.81 7.5 22 8.61 46.42 1.2 26.36 1.45 34.27 1.45 101s-.25 74.66-1.45 101c-1.11 24.38-5.19 37.61-8.61 46.42a82.76 82.76 0 01-47.45 47.45c-8.81 3.42-22 7.5-46.42 8.61-26.36 1.2-34.26 1.45-101 1.45s-74.66-.25-101-1.45c-24.38-1.11-37.61-5.19-46.42-8.61a77.52 77.52 0 01-28.75-18.7 77.52 77.52 0 01-18.7-28.75c-3.42-8.81-7.5-22-8.61-46.42-1.2-26.36-1.45-34.27-1.45-101s.25-74.66 1.45-101c1.11-24.38 5.19-37.61 8.61-46.42a77.52 77.52 0 0118.7-28.75 77.52 77.52 0 0128.75-18.7c8.81-3.42 22-7.5 46.42-8.61 26.36-1.2 34.27-1.45 101-1.45m0-45c-67.9 0-76.41.29-103.07 1.5s-44.79 5.39-60.69 11.57A122.56 122.56 0 0042 42a122.56 122.56 0 00-28.88 44.24c-6.18 15.9-10.4 34.08-11.62 60.69S0 182.1 0 250s.29 76.41 1.5 103.07 5.44 44.79 11.62 60.69A122.56 122.56 0 0042 458a122.56 122.56 0 0044.28 28.84c15.9 6.18 34.08 10.4 60.69 11.62S182.1 500 250 500s76.41-.29 103.07-1.5 44.79-5.44 60.69-11.62a127.89 127.89 0 0073.12-73.12c6.18-15.9 10.4-34.08 11.62-60.69S500 317.9 500 250s-.29-76.41-1.5-103.07-5.44-44.79-11.62-60.69A122.56 122.56 0 00458 42a122.56 122.56 0 00-44.28-28.84c-15.9-6.18-34.08-10.4-60.69-11.62S317.9 0 250 0zm0 121.62A128.38 128.38 0 10378.38 250 128.38 128.38 0 00250 121.62zm0 211.71A83.33 83.33 0 11333.33 250 83.33 83.33 0 01250 333.33zM383.45 86.55a30 30 0 1030 30 30 30 0 00-30-30z"/></svg>
          </a>
          <a href="//www.youtube.com/user/bolcom" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="youtube">
            <span class="srt" id="sr-text">Volg ons op de youtube pagina van bol</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-youtube bi-lg" data-test="icon-youtube" focusable="false"><path fill-rule="evenodd" d="M9.522 15.2V8.662l6.484 3.28-6.484 3.256zM23.76 7.53s-.235-1.603-.954-2.31c-.913-.926-1.936-.93-2.405-.984C17.043 4 12.005 4 12.005 4h-.01s-5.038 0-8.396.236c-.47.053-1.492.058-2.406.985C.474 5.927.24 7.53.24 7.53S0 9.413 0 11.295v1.765c0 1.884.24 3.766.24 3.766s.234 1.603.953 2.309c.914.927 2.113.898 2.647.995 1.92.178 8.16.234 8.16.234s5.043-.008 8.401-.243c.47-.055 1.492-.06 2.405-.986.72-.706.954-2.31.954-2.31s.24-1.881.24-3.765v-1.765c0-1.882-.24-3.765-.24-3.765z"/></svg>
          </a>
          <a href="//www.linkedin.com/company/11699" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="linkedin">
            <span class="srt" id="sr-text">Volg ons op de linkedin pagina van bol</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-linkedin bi-lg" data-test="icon-linkedin" focusable="false"><path fill-rule="evenodd" d="M20.452 20.45h-3.56v-5.57c0-1.328-.022-3.036-1.85-3.036-1.851 0-2.134 1.447-2.134 2.942v5.664H9.352V8.997h3.413v1.566h.049c.475-.9 1.636-1.85 3.367-1.85 3.605 0 4.27 2.371 4.27 5.456v6.281zM5.339 7.433a2.063 2.063 0 110-4.13 2.065 2.065 0 010 4.13zM7.12 20.45H3.558V8.997H7.12V20.45zM23 0H1a1 1 0 00-1 1v22a1 1 0 001 1h22a1 1 0 001-1V1a1 1 0 00-1-1z"/></svg>
          </a>
          <a href="//www.tiktok.com/@bol" target="_blank" rel="noreferrer" class="px_common_mhp_click" data-common-mhp-click-name="tiktok">
            <span class="srt">Volg ons op de TikTok pagina van bol</span>
            <svg version="1.1" viewBox="0 0 24 24" aria-hidden="true" class="svg-inline--bi bi-tiktok bi-lg" data-test="icon-tiktok" focusable="false"><path d="M18.8911 5.40977C17.7061 4.63691 16.8506 3.40085 16.5839 1.95983C16.5262 1.64833 16.4946 1.32802 16.4946 1H12.7123L12.706 16.1582C12.6425 17.8556 11.2455 19.218 9.53267 19.218C9.00041 19.218 8.49901 19.085 8.05785 18.8525C7.04586 18.3199 6.35345 17.259 6.35345 16.0384C6.35345 14.2851 7.77978 12.8588 9.53267 12.8588C9.85996 12.8588 10.1737 12.9128 10.4708 13.0057V9.14439C10.1634 9.10252 9.85114 9.07644 9.53267 9.07644C5.69446 9.07681 2.57144 12.1998 2.57144 16.0384C2.57144 18.3937 3.74836 20.4779 5.54349 21.7382C6.67413 22.532 8.04977 23 9.53304 23C13.3716 23 16.4946 19.877 16.4946 16.0384V8.35207C17.9779 9.41695 19.7954 10.0443 21.7566 10.0443V6.26197C20.7002 6.26197 19.7161 5.94791 18.8911 5.40977Z" fill="currentColor"/></svg>
          </a>
        </div>
        <div class="links-inline">
          <a href="/nl/rnwy/account/newsletters">Meld je aan</a> voor korting en inspiratie in je mailbox
        </div>
      </li>
    </ul>
    

    <div class="wsp-footer__payoff" data-test="footer-conditions">
      <p class="wsp-footer__logo">
        <svg focusable="false" aria-hidden="true" role="img" aria-labelledby="logo-title">
          <title id="logo-title">bol de winkel van ons allemaal</title>
          <use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#logoWithTagline"></use>
        </svg>
        <span class="srt">bol de winkel van ons allemaal</span>
      </p>

      

      <div class="px_common_bottomlinks_click" data-test="footer-bottomlinks">
        <ul class="wsp-footer__terms  small-text">
          
          <li class="wsp-footer__terms--aligned">
            <a href="//www.thuiswinkel.org/leden/bol.com/certificaat" rel="external" target="_blank">
              <span class="srt">Certificaat Thuiswinkel Waarborg</span>
              <svg version="1.1" viewBox="0 0 48 48" aria-hidden="true" class="svg-inline--bi bi-thuiswinkel-waarborg bi-3x u-mr--xs" data-test="icon-thuiswinkel-waarborg" focusable="false"><path fill="#1e1e1e" d="M.52 30.77h1.26v1h.9v1h-.9v1.89c0 .72.22.87.71.87h.23v1.1a2.39 2.39 0 01-.55.05c-1.15 0-1.66-.65-1.66-2v-1.84H0v-1h.52zm7.32 3.09v2.82H6.58v-3a.82.82 0 00-.81-.88.86.86 0 00-.88.92v2.93H3.62v-6.72h1.27v2.44a1.49 1.49 0 011.28-.69c1.39 0 1.67 1.14 1.67 2.18zm1.08.75V31.8h1.27v3a.82.82 0 00.81.88.85.85 0 00.86-.91V31.8h1.27v4.88h-1.25v-.57a1.52 1.52 0 01-1.29.69c-1.38 0-1.67-1.15-1.67-2.19zm5.94-4.72a.72.72 0 01.75.73.75.75 0 11-.75-.73zm-.63 1.91h1.28v4.88h-1.28zm2.1 4.09l.75-.67a1.38 1.38 0 001.11.53c.36 0 .59-.16.59-.44 0-.74-2.28-.45-2.28-2.19 0-.9.77-1.44 1.76-1.44a1.85 1.85 0 011.65.86l-.84.64a1.09 1.09 0 00-.87-.47c-.29 0-.5.15-.5.39 0 .78 2.3.4 2.3 2.15a1.66 1.66 0 01-1.86 1.55 2.07 2.07 0 01-1.81-.91zm7.82-2.34l-.9 3.13h-1.42l-1.45-4.88h1.39l.85 3.46.93-3.46h1.2l.92 3.46.86-3.46h1.38l-1.45 4.88H25zm5.09-3.66a.72.72 0 01.75.73.75.75 0 11-.75-.73zm-.62 1.91h1.27v4.88h-1.27zm6.64 2.06v2.82H34v-3a.82.82 0 00-.81-.88.85.85 0 00-.86.92v2.93H31V31.8h1.25v.57a1.5 1.5 0 011.29-.69c1.46 0 1.72 1.14 1.72 2.18zm1.1-3.93h1.27v3.92l1.56-2.05h1.43l-1.7 2.21 1.83 2.67h-1.5l-1.62-2.52v2.52h-1.27zm7.16 1.75c1.5 0 2.35 1 2.35 2.6a2.62 2.62 0 010 .28h-3.4a1.19 1.19 0 001.26 1.12 1.4 1.4 0 001.17-.6l.84.64a2.27 2.27 0 01-2 1.08 2.562 2.562 0 01-.18-5.12zm1.06 2a1 1 0 00-1-1 1 1 0 00-1 1zm2.14-3.79H48v6.75h-1.28zM8.77 41.76l-1.2 3.58h-.74l-1.51-4.51h.74l1.17 3.67 1.2-3.71h.67c1.19 3.71 1.2 3.71 1.2 3.71l1.17-3.71h.74l-1.51 4.55H10zm5.64.71a1.7 1.7 0 011.18.41v-.56a1 1 0 00-1.09-1 2.54 2.54 0 00-1.31.39l-.28-.51a3.19 3.19 0 011.65-.49 1.5 1.5 0 011.7 1.65v3h-.58l-.08-.38a1.54 1.54 0 01-1.21.49 1.5 1.5 0 110-3zm.09 2.41a1 1 0 001.11-.94c0-.56-.45-.92-1.12-.92s-1.11.39-1.11.92a1 1 0 001.12.94zm4.42-2.38a1.69 1.69 0 011.17.41v-.56a1 1 0 00-1.08-1 2.54 2.54 0 00-1.31.39l-.28-.51a3.16 3.16 0 011.65-.49 1.5 1.5 0 011.7 1.65v3h-.58l-.08-.38a1.54 1.54 0 01-1.21.49 1.5 1.5 0 110-3zm.08 2.38a1 1 0 001.11-.94c0-.56-.45-.92-1.12-.92s-1.11.39-1.11.92a1 1 0 001.12.94zm5-3.48c-1 0-1.32.82-1.32 1.78v2.16H22v-4.51h.67v.76a1.39 1.39 0 011.33-.87zm.89-2.33h.67v2.49a1.87 1.87 0 011.58-.84 2.16 2.16 0 012.13 2.37 2.19 2.19 0 01-2.18 2.41 1.84 1.84 0 01-1.55-.82v.71h-.65zm2.16 5.76a1.57 1.57 0 001.53-1.73 1.54 1.54 0 10-3 0 1.57 1.57 0 001.47 1.73zm5.26-4.11A2.37 2.37 0 1130 43.09a2.28 2.28 0 012.31-2.37zm0 4.11a1.75 1.75 0 10-1.63-1.74 1.62 1.62 0 001.62 1.74zm5.33-3.43c-1 0-1.32.82-1.32 1.78v2.16h-.68v-4.51h.67v.76a1.38 1.38 0 011.33-.87zM39 46.01a2.67 2.67 0 001.34.38c.85 0 1.58-.36 1.58-1.36v-.4a1.83 1.83 0 01-1.55.82 2.38 2.38 0 010-4.73 1.84 1.84 0 011.58.84v-.73h.65v3.88A2 2 0 0140.36 47a3.23 3.23 0 01-1.59-.45zm1.41-4.67a1.76 1.76 0 101.53 1.75 1.59 1.59 0 00-1.53-1.75zM24 0L10.65 7.77v15.59h26.7V7.77zm11.93 21.92H24.71v-6l11.22-6.49zm-12.64 0H12.07V9.43l11.22 6.53zm.71-7.2L12.79 8.19 24 1.66l11.21 6.53z"/><path fill="#bdea18" d="M35.92 9.43l-11.21 6.53v5.96h11.21V9.43z"/><path fill="#fc3f4d" d="M12.07 21.92h11.22v-5.96L12.07 9.43v12.49z"/><path fill="#fff" d="M12.79 8.19L24 14.72l11.21-6.53L24 1.66 12.79 8.19z"/></svg>
            </a>

            Geldt voor aankopen bij <a href="/nl/nl/sf/thuiswinkel-waarborg/">Thuiswinkel-leden</a>
          </li>
          
          <li><a href="/nl/nl/tc/">Algemene voorwaarden</a></li>
          <li><a href="/nl/nl/tc/privacybeleid/">Privacy</a></li>
          <li><a href="/nl/nl/tc/cookiebeleid/">Cookies</a></li>
          <li><a href="/nl/nl/m/copyright/">&copy;</a> 1999-2024 bol.com b.v.</li>
        </ul>

        <p class="wsp-footer__btw  px_common_disclaimer_click">
          
          * De voordelen van bol.com gelden niet voor het gehele assortiment.
          <a href="/nl/nl/sf/winkelen-zonder-zorgen/">Bekijk de voorwaarden</a>
          <br class="medium--is-hidden">
          

          Alle prijzen zijn inclusief BTW en andere heffingen en exclusief eventuele
          <a href="/nl/nl/klantenservice/a/5711226101301248/bezorgopties-en-kosten">verzenden en servicekosten</a>
        </p>
      </div>
    </div>
  </div>
  
    
        <wsp-conversational-ui  data-css="//assets.s-bol.com/nl/static/assets/css/chat.min.03ddabfc55ec65197b91.css" >
    
     <div id="cui-root" >
       
     </div>
     
     <button class="o-chat__button js_cui_notifier_button is-hidden" id="cui-chat-bubble" data-test="cui-bubble">
       <svg version="1.1" viewBox="0 0 14 14" aria-label="Chatbot Billie" role="img" class="svg-inline--bi bi-chat svg-inline--bi bi-chat bi-3x" data-test="icon-chat" focusable="false"><path fill-rule="evenodd" d="M12 13s-1.471.016-2.5-1.5c-.771.284-1.643.5-2.5.5-4 0-6-3-6-5s2-5 6-5 6 3 6 5-1 3-1 3v3zm-1-1.374V9.578S12 8.88 12 7c0-1.79-1.891-4-5-4-3.328 0-5 2.44-5 4s1.667 4 5 4c.714 0 1.494-.167 2.89-.71.159.19.507.928 1.11 1.336z"/></svg>
     </button>
     

     <script type="application/json" data-js-payload data-payload-id="conversational_ui">
        {"isNewBillieGraphicExpActive":true,"textInputComponentMigrationEnabled":false,"bodyContainerComponentMigrationEnabled":false,"messageComponentMigrationEnabled":false,"windowContainerComponentMigrationEnabled":false,"headerComponentMigrationEnabled":false,"conversationalProtocolV4":false,"authRedirectEnabled":false,"isLiveChatRegisterFixEnabled":true,"isLocalStorageAndroidEnabled":true,"isFocusTrapEnabled":true,"pageName":"/nl/nl/s/","chatrUrl":"https://chatr.bol.com/v1/p3/converse","chatrImageUploadUrl":"https://chatr.bol.com/v1/p3/image","chatrIsTypingUrl":"https://chatr.bol.com/v1/converse/isTyping","countryCode":"NL","language":"nl-NL","deviceType":"DESKTOP","onLoginPage":false,"authState":"RECOGNIZED","labels":[{"key":"chosenLiveChat","value":"Gekozen voor live chat"},{"key":"headerAgentConnected","value":"In gesprek met $agent"},{"key":"submitButtonTitle","value":"Typ een bericht in het tekstveld, en druk dan op deze knop om het bericht te zenden. (Of druk op ENTER)"},{"key":"alwaysOnline","value":"Altijd online"},{"key":"proactiveHeader","value":"Stel je vraag, krijg meteen antwoord"},{"key":"billieInitProductText","value":"Hallo, ik ben Billie, de chatbot van bol. Ik kan je snel helpen. En anders stuur ik je door naar de juiste persoon bij bol of onze partners. Wat wil je over dit artikel vragen?"},{"key":"downloadingImageFromBackend","value":"Afbeelding wordt opgehaald"},{"key":"headerAgentTyping","value":"$agent is aan het typen"},{"key":"headerSearchingExpert","value":"Zoeken naar een expert"},{"key":"askYourQuestion","value":"Stel je vraag"},{"key":"attachmentError401","value":"Log in om bijlages te kunnen verzenden"},{"key":"plazaPhoneErrorText","value":"Sorry er ging iets mis. Probeer het nog een keer."},{"key":"proactiveWelcome","value":"Op zoek naar antwoorden? Die hebben wij. Stel gerust je vraag."},{"key":"addAttachment","value":"Voeg een bijlage toe"},{"key":"noConnectionToChat","value":"Ik kan je niet verbinden met de chat. Ik verzamel nu de andere contactopties voor jou."},{"key":"attachmentErrorDefault","value":"Er is iets mis gegaan, je bijlage kan niet worden verzonden. Probeer het later nog eens."},{"key":"askNewQuestion","value":"Nieuwe vraag stellen"},{"key":"thereAreNewMessages","value":"Er staan $count nieuwe berichten voor je klaar"},{"key":"menuMinimize","value":"Gesprek minimaliseren"},{"key":"loginAttachmentLink","value":" om een bijlage te sturen."},{"key":"plazaPhoneNumberDescription","value":"Je kunt ze op dit nummer bereiken:"},{"key":"plazaShowNumber","value":"Telefoonnummer tonen"},{"key":"searchingExpert","value":"Zoeken naar een expert"},{"key":"noConnectionContactBot","value":" Ik kan je niet verbinden. Ik verzamel nu de andere contactopties voor jou."},{"key":"noConnectionToChatOtherContactOptions","value":"Er ging iets mis ... Door een onverwachte storing kun je niet verder chatten.\r\n      We zijn wel bereikbaar op een andere manier. Ik zoek even de contactopties voor jou."},{"key":"somethingWentWrong","value":"Er ging iets mis..."},{"key":"cannotSendMessage","value":"Kan bericht niet versturen"},{"key":"billieInitTextSub","value":"Hoe kan ik je helpen?"},{"key":"menuClose","value":"Gesprek afsluiten"},{"key":"inputTitle","value":"Typ hier je bericht."},{"key":"msgLongerThanExpected","value":"Sorry, het is nu wat drukker dan verwacht"},{"key":"thereIsANewMessage","value":"Er staat $count nieuw bericht voor je klaar"},{"key":"loginAttachment","value":"Log in om een bijlage te sturen."},{"key":"srUserChatMessage","value":"Gebruiker zegt:"},{"key":"liveChatStarted","value":"Je live chat is gestart"},{"key":"connectWithAgent","value":"We verbinden je daarna met een expert"},{"key":"billieError","value":"Ik kan je niet verbinden met onze chatbot Billie. Ik verzamel nu de andere contactopties voor jou."},{"key":"cannotSendImage","value":"Kan bijlage niet versturen"},{"key":"groupedChannelHeader","value":"Kom je er niet uit met de partner?"},{"key":"submitButtonName","value":"versturen"},{"key":"contactNow","value":"Ik wil contact opnemen"},{"key":"typeSomething","value":"Typ eerst je bericht"},{"key":"imagePending","value":"Bijlage wordt verstuurd"},{"key":"billieInitTextNew","value":"Hallo, ik ben Billie, de chatbot van bol. Ik kan je snel helpen. En anders weet ik wie bij bol je verder kan helpen."},{"key":"attachmentError500","value":"De server geeft een error,je bijlage kan niet worden verzonden. Probeer het later nog eens."},{"key":"billieInitProductTextSub","value":"Wat wil je over dit artikel vragen?"},{"key":"welcomeWaiting","value":"Typ gerust vast een bericht ..."},{"key":"chattingWithOtherAgent","value":"Je chat nu met een andere $agent"},{"key":"login","value":"Log in"},{"key":"yourCustomerNumberIs","value":"Je klantnummer is"},{"key":"noPartnerSolution","value":"Als je er met de partner niet uitkomt, zoeken we samen naar een oplossing."},{"key":"contactPlazaHeader","value":"Contact opnemen met onze partner"},{"key":"inputPlaceholder","value":"Je vraag of bericht ..."},{"key":"errorTimeoutRetry","value":"Er ging iets mis. We verbinden je opnieuw."},{"key":"noConnectionToBillie","value":"Er ging iets mis ... Door een onverwachte storing kun je niet verder chatten met onze chatbot Billie. We zijn wel bereikbaar op een andere manier. We zoeken even de contactopties voor jou."},{"key":"chooseContactOptions","value":"Kies een contactoptie"},{"key":"bolCustomerService","value":"De bol klantenservice is er voor je"},{"key":"attachmentError415","value":"Je bijlage is het verkeerde type en kan niet worden verzonden. Kies een .jpg, .png of .gif afbeelding"},{"key":"attachmentError413","value":"Je bijlage is te groot en kan niet worden verzonden"},{"key":"billieInitText","value":"Hallo, ik ben Billie, de chatbot van bol. Ik kan je snel helpen. En anders stuur ik je door naar de juiste persoon bij bol of onze partners. Hoe kan ik je helpen?"},{"key":"chatBotBillie","value":"Chatbot Billie"},{"key":"tryNow","value":"direct proberen"},{"key":"srBillieChatMessage","value":"Chatbot Billie zegt:"},{"key":"noConnectionPossible","value":"Geen verbinding mogelijk"},{"key":"cantPreview","value":"Kan afbeelding niet tonen"},{"key":"errorTimeoutCountdown","value":"Er ging iets mis. We verbinden je opnieuw over $countdown seconden."},{"key":"msgKeepOnSearching","value":"We blijven voor je zoeken"},{"key":"chattingWithAgent","value":"Je chat nu met $agent"},{"key":"menuRestart","value":"Gesprek opnieuw starten"}],"isCuiFromPackage":false,"isBillieBannerImprovements":false,"isBillieIntroductionExpActive":false,"isCUIResetSwitchActive":true,"isUrlMessageParamsFixEnabled":false,"themeName":"","recaptchaSiteKey":"6LfvbwkTAAAAANOIO8BXcuW1TzqOPfdZ3aUPhuPY","isFullScreen":false,"isStandalone":false,"csHomeUrl":"https://www.bol.com/nl/m/klantenservice/","conversationalUIAsyncMeasurementEndpoint":"/nl/conversational-ui-m2-data/send-conversation-id","initializeConversationType":"ContactBot","abTests":["BLSM_CUI_RESET_APP_BUG","CUI-BASSIE-V2_a","CUI-CHITOS-COPY-RETURN_control","CUI-CHITOS-PAY_a","CUI-CHITOS-RETURN_a","CUI-CHITOS-WARREP_a","CUI-CHITOS-WIMP_a"],"unleashSwitches":[],"bucketingId":"fl509pipaacmjnwat9pqht7tu31joh5w","fallbackChannels":[{"type":"phone","titleShort":"Bel ons","title":"Bel met onze klantenservice","responseTime":"Binnen 1 minuut reactie","description":"We zijn 24 uur per dag, 7 dagen per week bereikbaar","buttonText":"Bel 030 310 49 91","buttonUrl":"tel:0303104991"},{"type":"mail","titleShort":"Mail ons","title":"","responseTime":"We reageren binnen 24 uur","description":"Mogelijk om een foto mee te sturen","buttonText":"Ga naar mail formulier","buttonUrl":"/nl/contactformulier"},{"type":"liveChat","titleShort":"Chat met ons","title":"Chat met onze klantenservice","responseTime":"Zo snel mogelijk reactie","description":"We zijn dag en nacht bereikbaar. Ook in het weekend.","buttonText":"Start live chat","buttonUrl":"","errorType":"ParleyError"}],"conversationalUIM2Wrapper":{"m2":{"group":{"group-name":"conversationalUI","bltgg":"sOloXhA4zqCekILbwytNBQ.8_9"}}}}
     </script>
      
      </wsp-conversational-ui>
      



</footer>

</div>


  <script type="application/json" data-component-id="bltg-measurement">
    { 
      "PERCENTAGE_ITEM_IN_VIEW": 50,
      "TIME_IN_VIEW_MS": 1000,
      "SEND_THROTTLE_MS": 250,
      "CHECK_FOR_METRICS_MS": 250,
      "BLTG_TIMEOUT_IN_MS": 100,
      "URL_BLTG_END_POINT": "https://mvw.s-bol.com/v1/client-interaction"
    }
  </script>
<noscript>
    <div class="srt">
        <img src="https://swa.bol.com/b/ss/advbolprod2/1/H.24.2-NS/0" height="1" width="1" alt=""/>
    </div>
</noscript>
<wsp-mopinion>
    <script type="application/json" data-payload data-payload-id="mopinion">{"deploymentId":"iylvyb75XjFdC2iUCNMO18sUE4lEXmnQCPCwOozu","proxyDomain":"spoor.bol.com/app/v1"}</script>
</wsp-mopinion>
<script data-test="performance-script">
  !function(){if('PerformanceLongTaskTiming' in window){var g=window.__tti={e:[]};
    g.o=new PerformanceObserver(function(l){g.e=g.e.concat(l.getEntries())});
    g.o.observe({entryTypes:['longtask']})}}();

  var perfMetrics = (function (w, p) {
    return w[p] = w[p] || function(){(w[p].q=w[p].q||[]).push(arguments)};
  })(window, '@wsp/performance-metrics');

  perfMetrics('init', {"apiUrl":"https://firefly.bol.com","apiKey":"WEBSHOP","softwareVersion":"20241128091440","whiteList":["^https:\\/\\/[^\\/]+\\.(s-)?bol\\.com\\/"],"blackList":["^https:\\/\\/[^\\/]+\\.bol\\.com\\/(nl\\/bltg|tracking)\\/","^https:\\/\\/swa\\.bol\\.com\\/"]});
</script>
<script>
  (function (w, d, u) {
    function l() {
      var t = d.createElement("script");
      t.async = true;
      t.src = u;
      d.getElementsByTagName("head")[0].appendChild(t);
    }
    w.addEventListener('pageshow', function () {
      if ('requestIdleCallback' in w) {
        w.requestIdleCallback(l);
      } else {
        w.setTimeout(l, 200);
      }
    });
  })(window, document, "https://assets.s-bol.com/0.4.0/performance-metrics.js");
</script>
<wsp-analytics><script type="application/json" data-payload data-payload-id="js_analytics_data_ga">{"waitForConsent":true,"fbPixId":"903451823449917","analyticsObj":{"pageInfo":{"pageType":"LPSEARCH","country":"NL","shoppingChannelContextTypeAndDeviceType":"www.bol.com,DESKTOP","canonicalUrl":"https://www.bol.com/nl/nl/l/alle-artikelen/","shortURL":"/s/","countryLanguage":"nl-nl","external":true,"subscriptionInfo":{"selectMembership":false},"shouldSendUserProps":true,"renderedInApp":false},"products":[{"productId":"9200000011781470","title":"The Reckoners - Steelheart","category":"0","brand":"Brandon Sanderson","brick":"10000926","seller":"0_","price":5.49,"discount":0.0,"categoryNumbersFlattened":["8299","2510","2551","40413"],"chunk":"80007266","categories":["Boeken","Fantasy \u0026 Sciencefiction","Fantasy"],"orderable":"available","quantity":1}],"searchTerms":"9780575104181"}}</script></wsp-analytics>
<script>(function(){
//rnwy_comp: device_profile/v1/default
$RNWY.loader()
    .now('webshop')
    .now(function() {
      WSP.utils.deviceProfile.setDeviceProfile('DESKTOP');
    });

//rnwy_comp: customer_menu/v1/customer_menu
$RNWY.loader()
  .win('webshop')
  .win('offcanvas')
  .win(function () {
    
      WSP.data.user.set({
        firstName: 'Sebastiaan'
      });
    
  
      const accountToggle = document.getElementById('account-toggle');

      if (!accountToggle) return null
      const accountMenu = document.getElementById('account-menu');
      const menuItems =  Array.from(document.querySelectorAll('#accountNav > li a'));
      const nextHandle = document.getElementById('nextHandle');

      let preventFocusTrap = false;
      let currentFocus = -1;
      let mouseDownTarget = null;

      const isFocusable = (element) => {
        if (element.tabIndex < 0) return false;
        if (element.disabled) return false;

        switch (element.nodeName) {
          case 'A':
            return !!element.href && element.rel !== 'ignore';
          case 'INPUT':
            return element.type !== 'hidden';
          case 'BUTTON':
          case 'SELECT':
          case 'TEXTAREA':
            return true;
          default:
            return false;
        }
      }

      const attemptFocus = (element) => {
        if (!isFocusable(element)) return false
        preventFocusTrap = true
        element.focus();
        preventFocusTrap = false;

        return document.activeElement === element;
      }


      const focusFirstDescendant = (element) => {
        for (let i = 0; i < element.childNodes.length; i++) {
          const child = element.childNodes[i];
          if (attemptFocus(child) || focusFirstDescendant(child)) return true;
        }      
      }

      const focusLastDescendant = (element) => {
        for (let i = element.childNodes.length - 1; i >= 0; i--) {
          const child = element.childNodes[i];
          if (attemptFocus(child) || focusLastDescendant(child)) return true;
        }      
      }

      // Trap focus within the menu
      const trapFocus = (e) => {
        if (preventFocusTrap) return;
      
        const isDialogBlurred = !accountMenu.contains(e.target);
      
        if (isDialogBlurred) {
          if (document.activeElement === nextHandle) {
            focusFirstDescendant(accountMenu);
            currentFocus = 0;
          } else {
            focusLastDescendant(accountMenu);
            currentFocus = menuItems.length - 1;
          }
        } else {
          currentFocus = menuItems.indexOf(document.activeElement);
        }
      };

      function setFocusAttribute(isFocused) {
        accountToggle.setAttribute('data-focus', isFocused ? 'true' : 'false');
      }
      
      accountToggle.addEventListener('focus', () => setFocusAttribute(true));
      accountToggle.addEventListener('blur', () => setFocusAttribute(false));
      
      const accountToggleIsFocused = document.activeElement === accountToggle;
      setFocusAttribute(accountToggleIsFocused);

      const showMenu = () => {
        document.addEventListener('focus', trapFocus, true);
        focusFirstDescendant(accountMenu);
        currentFocus = 0;
      }

      function toggleMenu() {
        const isExpanded = accountToggle.getAttribute('aria-expanded') === 'true';

        accountToggle.setAttribute('aria-expanded', !isExpanded);
        if (!isExpanded) {
          accountMenu.classList.add('is-visible');
          accountMenu.setAttribute('aria-hidden', 'false');
          showMenu()
        } else {
          closeMenu()
        }
      }
    
      function closeMenu() {
        document.removeEventListener('focus', trapFocus, true);
        accountMenu.classList.remove('is-visible');
        accountToggle.setAttribute('aria-expanded', 'false');
        accountMenu.setAttribute('aria-hidden', 'true');
        accountToggle.setAttribute('data-focus', 'false')
        currentFocus = -1;
      }
    
      accountToggle?.addEventListener('click', (e) => {
        // This prevents click from bubbling to document
        e.stopPropagation(); 
        toggleMenu();
      });
    
      document?.addEventListener('click', (e) => {
        if (!accountMenu?.contains(e.target) && e.target !== accountToggle) {
          closeMenu();
        }
      });
    
      document?.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
          closeMenu();
        }
      });

      function handleNavigation(event) {
        const key = event.key;
        if (key === 'ArrowDown' || key === 'ArrowUp' || key === 'Tab') {
          event.preventDefault();
          
          if (key === 'ArrowDown' || (key === 'Tab' && !event.shiftKey)) {
            currentFocus = (currentFocus + 1) % menuItems.length;
          } else if (key === 'ArrowUp' || (key === 'Tab' && event.shiftKey)) {
            currentFocus = (currentFocus - 1 + menuItems.length) % menuItems.length;
          }
      
          menuItems[currentFocus].focus();
        }
      }
      
      function handleMenuItemMouseDown(event) {
        const targetItem = event.target.closest('a');
        if (targetItem && menuItems.includes(targetItem)) {
          mouseDownTarget = targetItem;
          currentFocus = menuItems.indexOf(targetItem);
        }
      }

      document.getElementById('accountNav').addEventListener('keydown', handleNavigation);
      document.getElementById('accountNav').addEventListener('mousedown', handleMenuItemMouseDown);
});

//rnwy_comp: basket_button/v9/basket_button
$RNWY.loader()
  .now('jquery')
  .win('basket')
  .win(function () {
    WSP.core.mediator.subscribe('wsp:accessories_helper_ensure_load_accessories_helper_dependencies', function(callback) {
      if(typeof WSP.Basket === 'undefined') {
        WSP.Basket = new WSP.core.Basket({});
      }
      if(typeof WSP.BasketButtons === 'undefined') {
        WSP.BasketButtons = new WSP.core.BasketButtons({});

        WSP.core.mediator.subscribe(WSP.events.Basket.basketUpdated, function(data) {
          WSP.AccessoriesHelper.checkSubserviceButtons(data);
        });
      }
      callback();
    });

    WSP.StickyBasket = new WSP.core.StickyBasket({
      addToCartElements   : '.js_floating_basket_btn',
      basketUrl       : '/nl/nl/basket/',
      loadAccessoriesInline : true,
      stickyBasketTargetEl  : jQuery('.js_basket_wrapper'),
      showBasketOnHover     : false
    });

    if(typeof WSP.AccessoriesHelper === 'undefined') {
      WSP.AccessoriesHelper = new WSP.core.AccessoriesHelper();
    }

  });

//rnwy_comp: search_bar/v5/default
$RNWY.loader()
  .now('jquery')
  .dom(function () {
    document.querySelector('.js_main_menu_btn').addEventListener('click', function () {
      var mainNavOffcanvas = document.querySelector('wsp-main-nav-offcanvas.js_main_menu');
      if (mainNavOffcanvas) {
        mainNavOffcanvas.toggleVisibility(true);

        $RNWY.loader()
          .win('px')
          .win(function () {
            px('common.mobilemenu.open', { fromHeader: true });
          });
      }
    });
  });

//rnwy_comp: price_block/v3/price_block_promo_param


//rnwy_comp: main_object_container/v4/default_ce
$RNWY.loader()
  .win('offcanvas');

//rnwy_comp: footer_default/v4/footer
$RNWY.loader().dom(function () {
  const icons = Array.from(document.querySelectorAll('.wsp-footer__social-icons a svg'));
  const srTexts = Array.from(document.querySelectorAll('#sr-text'));

  const iconTypes = ['facebook', 'pinterest', 'instagram', 'youtube', 'linkedin', 'tiktok'];

  icons.forEach((svg, index) => {
    // Skip if we don't have a defined icon type
    if (index >= iconTypes.length) return;

    const iconType = iconTypes[index];
    const srText = srTexts[index] ? srTexts[index].textContent : "Follow us on ${iconType}";

    const iconTitle = document.createElement('title');
    iconTitle.id = "${iconType}-icon-title";
    iconTitle.textContent = srText;

    const iconDesc = document.createElement('desc');
    iconDesc.id = "${iconType}-icon-desc";
    iconDesc.textContent = "Icon for ${iconType}";

    svg.setAttribute('role', 'img');
    svg.setAttribute('aria-labelledby', "${iconType}-icon-title");
    svg.setAttribute('aria-describedby', "${iconType}-icon-desc");

    svg.insertBefore(iconDesc, svg.firstChild);
    svg.insertBefore(iconTitle, svg.firstChild);

    svg.setAttribute('data-icon', iconType);
  });
});

//rnwy_comp: foot/v6/foot
window['@wsp/analytics-loader'] = window['@wsp/analytics-loader'] || {};
window['@wsp/analytics-loader'].waitForConsent = true;

$RNWY.loader()
  
  
    .win('jquery')
    .win('flex-tooltip')
    .win(function() {
    if (WSP) {
       new WSP.core.TooltipController({
         tooltipElemsSelector: '.info-link'
       });
    }

  });
    




//rnwy_comp: analytics_common/v1/analytics_common
window['@wsp/analytics-loader'] = window['@wsp/analytics-loader'] || {};
window['@wsp/analytics-loader'].waitForConsent = true;

$RNWY.loader()
  .win('analytics-loader');

//rnwy_comp: measurement/v3/bltg

  $RNWY.loader()
    .win('measurement');

//rnwy_comp: tracking/v1/pixel
$RNWY.loader()
    .win('px')
    .win(function(){
        px('common.page.view', {
  "channel" : "main",
  "eVar11" : "select_nl:no|select_be:no|kp:no",
  "eVar12" : "recognized",
  "eVar14" : "typed search",
  "eVar15" : "1",
  "eVar16" : "customer_loyalty_status:0",
  "eVar18" : "9780575104181",
  "eVar19" : "1",
  "eVar20" : "bestverkocht_11",
  "eVar21" : "consumer",
  "eVar24" : "view:list|standard",
  "eVar31" : "9002219944",
  "eVar51" : "www.bol.com,DESKTOP",
  "eVar52" : "typed search",
  "eVar55" : "79327cab-fc31-4deb-8abb-7f13948535d7",
  "eVar67" : "ALL",
  "eVar74" : "heterogeneous",
  "eVar75" : "st:0|fb1:0|fb2:0|fb3:0|fb4:0|fb5:0|fb6:0|vce:0|bc:0|rf:0|cf:0|ih:0|se:0|hs:1|af:0|bs:0|tf:0|fr:0|qc:0",
  "eVar93" : "external-behavioural=no|internal-behavioural=no|internal-transactional=no|external-transactional=no",
  "list3" : "CUI-CHITOS-RETURN_a,CUI-CHITOS-PAY_a,CUI-CHITOS-WIMP_a,wsp-search-WSP1-12767_a,CUI-BASSIE-V2_a,CUI-CHITOS-COPY-RETURN_control,OXYGEN-1755_a,ATLANTIS-1387_control,WSP1-14681_a,CUI-CHITOS-WARREP_a,WSP1-14406_a,wsp-search-mopinion_a",
  "pageName" : "\/nl\/nl\/s\/",
  "prop1" : "main:quickSearch",
  "prop11" : "1",
  "prop13" : "9780575104181",
  "prop16" : "ALL",
  "prop2" : "no cms page",
  "prop34" : "search",
  "prop37" : "select:no|nl",
  "prop56" : "unknown:nl:NL:nl-NL_bolcom_media_www:",
  "prop68" : "\/s\/",
  "server" : "wspc-deployment-77748977d4-mpg5x"
});
    });


}());</script>
<script type="text/javascript"  src="/qpgFclBoxchhIObcn3rSa0gC/VX3NDmf7EEiN/CEdLAQ/SH0a/AlwVUlUB?v=8e05e47b-fd69-0b26-b3ef-b046a04f0210" defer></script></body>
</html>
`

// Test for successful scraping
func TestScrapeData_Success(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHTML))
	}))
	defer server.Close()

	// Define scraper config
	config := ScraperConfig{
		SearchURL:      server.URL + "?isbn=%s",
		TitleSelector:  ".product-title",
		PriceSelector:  "meta[itemprop='price']",
		LinkSelector:   ".product-title",
	}
	
	// Perform scraping
	isbn := "9780575104181"
	results, err := ScrapeData(config, isbn)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}
	if results[0].Title != "The Reckoners - Steelheart" {
		t.Errorf("Expected title 'Test Book', got '%s'", results[0].Title)
	}
	if results[0].Price != "€5.49" {
		t.Errorf("Expected price '€12.34', got '%s'", results[0].Price)
	}
	if results[0].Link != "/nl/nl/p/the-reckoners-steelheart/9200000011781470/" {
		t.Errorf("Expected link '/product/12345', got '%s'", results[0].Link)
	}
}

// Test for missing elements
func TestScrapeData_NoMatch(t *testing.T) {
	// Create a test server with empty body
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body></body></html>"))
	}))
	defer server.Close()

	config := ScraperConfig{
		SearchURL:      server.URL + "?isbn=%s",
		TitleSelector:  ".product-title",
		PriceSelector:  ".promo-price",
		LinkSelector:   ".product-title",
	}

	isbn := "9780316769488"
	results, err := ScrapeData(config, isbn)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

// Test for HTTP error
func TestScrapeData_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	config := ScraperConfig{
		SearchURL: server.URL + "?isbn=%s",
	}

	isbn := "9780316769488"
	_, err := ScrapeData(config, isbn)

	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
package gobookprices

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGetBooksFromPage(t *testing.T) {

	t.Run("basic test", func(t *testing.T) {
		mockHTML := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage">
<body class="">
<div class="content" id="bodycontainer" style="">
<div class='siteHeader'>
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D2%26shelf%3Dto-read%26view%3Dtable&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".utrh2fji6g" data-react-checksum="-875785090"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".utrh2fji6g.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".utrh2fji6g.0.0"><picture data-reactid=".utrh2fji6g.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".utrh2fji6g.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".utrh2fji6g.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".utrh2fji6g.1"><div class="siteHeader__contents" data-reactid=".utrh2fji6g.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".utrh2fji6g.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".utrh2fji6g.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".utrh2fji6g.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".utrh2fji6g.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".utrh2fji6g.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".utrh2fji6g.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".utrh2fji6g.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".utrh2fji6g.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".utrh2fji6g.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".utrh2fji6g.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.2.0.2.0.0"><span data-reactid=".utrh2fji6g.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".utrh2fji6g.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".utrh2fji6g.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".utrh2fji6g.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.2.0.3.0.0"><span data-reactid=".utrh2fji6g.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".utrh2fji6g.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".utrh2fji6g.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".utrh2fji6g.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".utrh2fji6g.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".utrh2fji6g.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".utrh2fji6g.1.0.4"><ul class="personalNav" data-reactid=".utrh2fji6g.1.0.4.0"><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.4.0.0"><div data-reactid=".utrh2fji6g.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".utrh2fji6g.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".utrh2fji6g.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".utrh2fji6g.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".utrh2fji6g.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".utrh2fji6g.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".utrh2fji6g.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".utrh2fji6g.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.0"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D2%26shelf%3Dto-read%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".utrh2fji6g.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".utrh2fji6g.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".utrh2fji6g.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".utrh2fji6g.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".utrh2fji6g.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".utrh2fji6g.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".utrh2fji6g.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".utrh2fji6g.1.0.6"><div class="modal__content" data-reactid=".utrh2fji6g.1.0.6.0"><div class="modal__close" data-reactid=".utrh2fji6g.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".utrh2fji6g.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".utrh2fji6g.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".utrh2fji6g.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".utrh2fji6g.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".utrh2fji6g.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".utrh2fji6g.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".utrh2fji6g.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".utrh2fji6g.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".utrh2fji6g.1.0.7"><div data-reactid=".utrh2fji6g.1.0.7.0"><div class="modal__close" data-reactid=".utrh2fji6g.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".utrh2fji6g.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".utrh2fji6g.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".utrh2fji6g.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".utrh2fji6g.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".utrh2fji6g.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1"><ul data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D2%26shelf%3Dto-read%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".utrh2fji6g.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".utrh2fji6g.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".utrh2fji6g.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".utrh2fji6g.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".utrh2fji6g.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".utrh2fji6g.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".utrh2fji6g.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".utrh2fji6g.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".utrh2fji6g.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".utrh2fji6g.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".utrh2fji6g.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".utrh2fji6g.2.0.0.0.2.0.0"><span data-reactid=".utrh2fji6g.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".utrh2fji6g.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".utrh2fji6g.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".utrh2fji6g.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".utrh2fji6g.2.0.0.0.3.0.0"><span data-reactid=".utrh2fji6g.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".utrh2fji6g.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
        <a href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read&amp;view=table">My Books</a>: 
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
<a id="batchEditLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this, {afterOpen: startEditing, afterClose: stopEditing});; return false;">Batch Edit</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;print=true&amp;shelf=to-read&amp;view=table">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
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
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read%2Cread&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=read&amp;view=table">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read%2Ccurrently-reading&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=currently-reading&amp;view=table">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=&amp;view=table">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;view=table">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
                <a class="actionLink" href="#" onclick="$$(&#39;#paginatedShelfList .multiLink&#39;).invoke(&#39;toggle&#39;); return false;">select multiple</a>
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
      <form id="fieldsForm" class="edit_user_shelf" action="/shelf/update/222519907" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="patch" />        <table>
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
            <input type="submit" name="commit" value="Save Current Settings to Your &quot;to-read&quot; shelf" id="save_curr_sett_submit" class="gr-button gr-button--small" style="margin-right: 10px" />
            <span class="loading" style="display: none"><img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Saving...</span>
              <span class="greyText inter uitext">shelf settings customized</span>
              <input type="checkbox" name="reset_display_fields" id="reset_display_fields" value="true" style="display:none" />
            <span class="greyText status inter"></span>
          </div>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div id="batchEdit" style="display: none;" class="controlBody">
        <div id="shelfTools" class="toolset">
          <form name="reviewEditForm" id="reviewEditForm" action="/review/update_list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="pmCttGBfKs4miLevGkO0XGKeCcfaYWVXYFwdyoqfOEm1G9GMIm00X48TQm7//rawItW0u5+vdfKq2Bpi91gTuw==" />
            <input type="hidden" name="view" id="view" value="table" />
            <label>
              Shelf:
              <select name="edit[shelf]" id="edit_shelf"><option value="read">read</option>
<option value="currently-reading">currently-reading</option>
<option value="to-read">to-read</option></select>
              &nbsp;
            </label>
            <a id="add_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;add_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_59122&#39;);$(&#39;add_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;add_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_59122&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_59122&#39;);Element.hide(&#39;add_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_59122&#39;);Element.show(&#39;add_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=add_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;N1t+zRQL2bzGbpQBBQ6qaWclpllOJJkqXy6b8GlBFPkkIAL1VjnHLW/1YcDgs6iFJ24bJQvqiY+VqpxYFIY/Cw==&#39;)}); return false;">add books to this shelf</a><img style="display:none" id="loading_anim_59122" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_340350&#39;);$(&#39;remove_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_340350&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_340350&#39;);Element.hide(&#39;remove_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_340350&#39;);Element.show(&#39;remove_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=remove_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;/DKGn7qESWcMjrqe23WjPXcRu7ighviRc0PsNLhLAnTvSfqn+LZX9qUVT18+yKHRN1oGxOVI6DS5x+ucxYwphg==&#39;)}); return false;">remove books from this shelf</a><img style="display:none" id="loading_anim_340350" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_books_link" class="actionLinkLite" href="#" onclick="if (confirm(&#39;This will completely remove the selected books from your shelves.&#39;)) { new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_books_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_459535&#39;);$(&#39;remove_books_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_books_link&#39;).show();;Element.hide(&#39;loading_anim_459535&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_459535&#39;);Element.hide(&#39;remove_books_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_459535&#39;);Element.show(&#39;remove_books_link&#39;);              for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {
                var r = request.responseJSON.reviews[i];
                $(&#39;review_&#39;+r.object.id).fade();
              }
}, parameters:(&#39;form_action=remove_books&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;+6DRuAYS+VpNrgJbqF4kxUAnyParYEhVSAuG46EPN+Xo262ARCDny+Q195pN4yYpAGx1iu6uWPCCj4FL3MgcFw==&#39;)}); }; return false;">remove books from all shelves</a><img style="display:none" id="loading_anim_459535" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
</form>        </div>
        <div id="otherTools" class="toolset greyText">
          <div class="right">
                  <a class="actionLinkLite smallText" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=2&amp;shelf=to-read&amp;sort=position&amp;view=table">sort by position to re-order</a>
                  <span class="greyText">|</span>
                <a method="post" id="loading_link_39431700" class="actionLinkLite smallText" href="#" onclick="if (confirm(&#39;Are you sure you want to disable sorting?  This will remove the custom order you\&#39;ve applied to this shelf.&#39;)) { $(this).hide(); $(&#39;loading_anim_39431700&#39;).show(); $(&#39;hidden_link_39431700&#39;).simulate(&#39;click&#39;); }; return false;">disable sorting</a><img style="display:none" id="loading_anim_39431700" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /><a id="hidden_link_39431700" class="actionLinkLite smallText" style="display: none" rel="nofollow" data-method="post" href="/shelf/disable_sorting/222519907">disable sorting</a>
                <span class="greyText">|</span>
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
        <a id="loading_link_447803" class="button" href="#" onclick="new Ajax.Request(&#39;/shelf/move_batch/68156753&#39;, {asynchronous:true, evalScripts:true, onComplete:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;hide&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;show&#39;);}, onFailure:function(request){alert(&#39;Something went wrong re-ordering those shelves.&#39;);;Element.hide(&#39;loading_anim_447803&#39;);}, onLoading:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;show&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;hide&#39;);;Element.show(&#39;loading_anim_447803&#39;);Element.hide(&#39;loading_link_447803&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_447803&#39;);Element.show(&#39;loading_link_447803&#39;);$(&#39;booksBody&#39;).update(request.responseJSON.html);toggleFieldsToMatchHeader();startEditing();$(&#39;reorderConfirm&#39;).hide();$(&#39;booksBody&#39;).highlight();}, parameters:Form.serializeElements($$(&#39;#books .position input&#39;)) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;9xnGiI4JDU/wfw4T2nHlTYOluJEo4kTYNiXH6r5w0lrkYrqwzDsT3lnk+9I/zOehw+4F7W0sVH38ocBCw7f5qA==&#39;)}); return false;">apply position changes?</a><img style="display:none" id="loading_anim_447803" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          &nbsp;
          <a href="#" onclick="$(&#39;reorderConfirm&#39;).hide(); return false;">Not yet</a>
      </div>
      <div class="right uitext">
        <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=2&amp;shelf=to-read&amp;view=table">2</a> <a href="/review/list/68156753-sebastiaan?page=3&amp;per_page=2&amp;shelf=to-read&amp;view=table">3</a> <a href="/review/list/68156753-sebastiaan?page=4&amp;per_page=2&amp;shelf=to-read&amp;view=table">4</a> <a href="/review/list/68156753-sebastiaan?page=5&amp;per_page=2&amp;shelf=to-read&amp;view=table">5</a> <a href="/review/list/68156753-sebastiaan?page=6&amp;per_page=2&amp;shelf=to-read&amp;view=table">6</a> <a href="/review/list/68156753-sebastiaan?page=7&amp;per_page=2&amp;shelf=to-read&amp;view=table">7</a> <a href="/review/list/68156753-sebastiaan?page=8&amp;per_page=2&amp;shelf=to-read&amp;view=table">8</a> <a href="/review/list/68156753-sebastiaan?page=9&amp;per_page=2&amp;shelf=to-read&amp;view=table">9</a> <span class="gap">&hellip;</span> <a href="/review/list/68156753-sebastiaan?page=21&amp;per_page=2&amp;shelf=to-read&amp;view=table">21</a> <a href="/review/list/68156753-sebastiaan?page=22&amp;per_page=2&amp;shelf=to-read&amp;view=table">22</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=2&amp;shelf=to-read&amp;view=table">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="display: none">
              </th>
              <th alt="position" class="header field position" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=position&amp;view=table">#</a>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=cover&amp;view=table">cover</a>
              </th>
              <th alt="title" class="header field title" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=title&amp;view=table">title</a>
              </th>
              <th alt="author" class="header field author" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=author&amp;view=table">author</a>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=isbn&amp;view=table">isbn</a>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=isbn13&amp;view=table">isbn13</a>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=asin&amp;view=table">asin</a>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=num_pages&amp;view=table">num pages</a>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=avg_rating&amp;view=table">avg rating</a>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=num_ratings&amp;view=table">num ratings</a>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=date_pub&amp;view=table">date pub</a>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=date_pub_edition&amp;view=table">date pub (ed.)</a>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=rating&amp;view=table">rating</a>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=shelves&amp;view=table">shelves</a>
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=review&amp;view=table">review</a>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=notes&amp;view=table">notes</a>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=comments&amp;view=table">comments</a>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=votes&amp;view=table">votes</a>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=read_count&amp;view=table">read count</a>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=date_started&amp;view=table">date started</a>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=date_read&amp;view=table">date read</a>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=2&amp;shelf=to-read&amp;sort=date_added&amp;view=table">date</a>
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=2&amp;shelf=to-read&amp;sort=date_added&amp;view=table">
                      <nobr>
                        added <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                      </nobr>
</a>              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=owned&amp;view=table">owned</a>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=2&amp;shelf=to-read&amp;sort=format&amp;view=table">format</a>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_7064093266" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7064093266]" id="checkbox_review_7064093266" value="7064093266" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6736083077" name="positions[6736083077]" value="44">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6736083077'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6736083077&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6736083077').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6736083077').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    745,467
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
                <a id="loading_link_87105" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/45047384&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_87105&#39;);$(&#39;loading_link_87105&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_87105&#39;).show();;Element.hide(&#39;loading_anim_87105&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_87105&#39;);Element.hide(&#39;loading_link_87105&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_87105&#39;);Element.show(&#39;loading_link_87105&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;dR/3ZNTuPYxgt3s6CfiM4yzedf+gl/hZjhPQBEPNM3hmZItcltwjHcksjvvsRY4PbJXIg+VZ6PxEl9esPgoYig==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_87105" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7064093266">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The House in the Cerulean Sea from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/45047384?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D2%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_7009239588" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7009239588]" id="checkbox_review_7009239588" value="7009239588" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6679332637" name="positions[6679332637]" value="43">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6679332637'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6679332637&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6679332637').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6679332637').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    14,159
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
                <a id="loading_link_276240" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/127278666&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_276240&#39;);$(&#39;loading_link_276240&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_276240&#39;).show();;Element.hide(&#39;loading_anim_276240&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_276240&#39;);Element.hide(&#39;loading_link_276240&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_276240&#39;);Element.show(&#39;loading_link_276240&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;t0/B/mBTbCI9Fj2l5mhg1G3UmOCHfp+3l0N6ePr+3SekNL3GImFys5SNyGQD1WI4LZ8lnMKwjxJdx33Qhzn21Q==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_276240" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7009239588">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Fox Wife from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/127278666?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D2%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>
</tbody></table>    </div>`
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML))
		}))
		defer server.Close()

		want := []Book{{
			Title:       "The House in the Cerulean Sea (Cerulean Chronicles, #1)",
			Author:      "Klune, T.J.",
			WorkUrl:     "/book/show/45047384",
			EditionsUrl: "/work/editions/62945242",
		}, {
			Title:       "The Fox Wife",
			Author:      "Choo, Yangsze",
			Isbn:        "1250266017",
			WorkUrl:     "/book/show/127278666",
			EditionsUrl: "/work/editions/148387285",
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
<html class="desktop withSiteHeaderTopFullImage">
<body class="">
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
		want := []Book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books: \nWant: '%+v', Got: '%+v'", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
		}
	})
}

func TestGetBooks(t *testing.T) {
	mockHTML1 := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage">
<body class="">
<div id="leadercol">
  <div id="review_list_error_message" class="review_list_error_message" style="display: none;">
  </div>
  <div id="header" style="float: left">
    <h1>
        <a href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=table">My Books</a>: 
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
<a id="batchEditLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this, {afterOpen: startEditing, afterClose: stopEditing});; return false;">Batch Edit</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;print=true&amp;shelf=to-read&amp;view=table">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
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
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read%2Cread&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=read&amp;view=table">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read%2Ccurrently-reading&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=currently-reading&amp;view=table">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=&amp;view=table">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;view=table">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
                <a class="actionLink" href="#" onclick="$$(&#39;#paginatedShelfList .multiLink&#39;).invoke(&#39;toggle&#39;); return false;">select multiple</a>
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
      <form id="fieldsForm" class="edit_user_shelf" action="/shelf/update/222519907" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="patch" />        <table>
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
            <input type="submit" name="commit" value="Save Current Settings to Your &quot;to-read&quot; shelf" id="save_curr_sett_submit" class="gr-button gr-button--small" style="margin-right: 10px" />
            <span class="loading" style="display: none"><img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Saving...</span>
              <span class="greyText inter uitext">shelf settings customized</span>
              <input type="checkbox" name="reset_display_fields" id="reset_display_fields" value="true" style="display:none" />
            <span class="greyText status inter"></span>
          </div>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div id="batchEdit" style="display: none;" class="controlBody">
        <div id="shelfTools" class="toolset">
          <form name="reviewEditForm" id="reviewEditForm" action="/review/update_list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="8XLZ2qP/FAsDudwwkx4cYtlwlNpBQeQRCN/ZaGf3OyniCaXi4c0KmqoiKfF2ox6OmTsppgSP9LTCW97AGjAQ2w==" />
            <input type="hidden" name="view" id="view" value="table" />
            <label>
              Shelf:
              <select name="edit[shelf]" id="edit_shelf"><option value="read">read</option>
<option value="currently-reading">currently-reading</option>
<option value="to-read">to-read</option></select>
              &nbsp;
            </label>
            <a id="add_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;add_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_430827&#39;);$(&#39;add_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;add_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_430827&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_430827&#39;);Element.hide(&#39;add_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_430827&#39;);Element.show(&#39;add_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=add_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;DTEbLBARa/PXtoYSXwJKbfDr5MBVkmJ9jmPcbj2ijeIeSmcUUiN1Yn4tc9O6v0iBsKBZvBBccthE59vGQGWmEA==&#39;)}); return false;">add books to this shelf</a><img style="display:none" id="loading_anim_430827" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_883497&#39;);$(&#39;remove_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_883497&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_883497&#39;);Element.hide(&#39;remove_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_883497&#39;);Element.show(&#39;remove_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=remove_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;jqFh+P+GXEOMmCLDAop5ptk+F5XX4GzvyQPAD15Fq4Od2h3AvbRC0iUD1wLnN3tKmXWq6ZIufEoDh8enI4KAcQ==&#39;)}); return false;">remove books from this shelf</a><img style="display:none" id="loading_anim_883497" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_books_link" class="actionLinkLite" href="#" onclick="if (confirm(&#39;This will completely remove the selected books from your shelves.&#39;)) { new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_books_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_954980&#39;);$(&#39;remove_books_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_books_link&#39;).show();;Element.hide(&#39;loading_anim_954980&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_954980&#39;);Element.hide(&#39;remove_books_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_954980&#39;);Element.show(&#39;remove_books_link&#39;);              for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {
                var r = request.responseJSON.reviews[i];
                $(&#39;review_&#39;+r.object.id).fade();
              }
}, parameters:(&#39;form_action=remove_books&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;a2xEop68x42TLVN8na45cO5BnDvNoJ7CAjwON55zG/54Fzia3I7ZHDq2pr14EzucrgohR4hujmfIuAmf47QwDA==&#39;)}); }; return false;">remove books from all shelves</a><img style="display:none" id="loading_anim_954980" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
</form>        </div>
        <div id="otherTools" class="toolset greyText">
          <div class="right">
                  <a class="actionLinkLite smallText" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;sort=position&amp;view=table">sort by position to re-order</a>
                  <span class="greyText">|</span>
                <a method="post" id="loading_link_50091780" class="actionLinkLite smallText" href="#" onclick="if (confirm(&#39;Are you sure you want to disable sorting?  This will remove the custom order you\&#39;ve applied to this shelf.&#39;)) { $(this).hide(); $(&#39;loading_anim_50091780&#39;).show(); $(&#39;hidden_link_50091780&#39;).simulate(&#39;click&#39;); }; return false;">disable sorting</a><img style="display:none" id="loading_anim_50091780" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /><a id="hidden_link_50091780" class="actionLinkLite smallText" style="display: none" rel="nofollow" data-method="post" href="/shelf/disable_sorting/222519907">disable sorting</a>
                <span class="greyText">|</span>
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
        <a id="loading_link_64486" class="button" href="#" onclick="new Ajax.Request(&#39;/shelf/move_batch/68156753&#39;, {asynchronous:true, evalScripts:true, onComplete:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;hide&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;show&#39;);}, onFailure:function(request){alert(&#39;Something went wrong re-ordering those shelves.&#39;);;Element.hide(&#39;loading_anim_64486&#39;);}, onLoading:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;show&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;hide&#39;);;Element.show(&#39;loading_anim_64486&#39;);Element.hide(&#39;loading_link_64486&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_64486&#39;);Element.show(&#39;loading_link_64486&#39;);$(&#39;booksBody&#39;).update(request.responseJSON.html);toggleFieldsToMatchHeader();startEditing();$(&#39;reorderConfirm&#39;).hide();$(&#39;booksBody&#39;).highlight();}, parameters:Form.serializeElements($$(&#39;#books .position input&#39;)) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;ro7PDYUk3gwV/iMiY1gGc2fOdCrqelA63UYTmN7J61m99bM1xxbAnbxl1uOG5QSfJ4XJVq+0QJ8XwhQwow7Aqw==&#39;)}); return false;">apply position changes?</a><img style="display:none" id="loading_anim_64486" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          &nbsp;
          <a href="#" onclick="$(&#39;reorderConfirm&#39;).hide(); return false;">Not yet</a>
      </div>
      <div class="right uitext">
        <div id="reviewPagination"><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=20&amp;shelf=to-read&amp;view=table">2</a> <a href="/review/list/68156753-sebastiaan?page=3&amp;per_page=20&amp;shelf=to-read&amp;view=table">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=20&amp;shelf=to-read&amp;view=table">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="display: none">
              </th>
              <th alt="position" class="header field position" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=position&amp;view=table">#</a>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=cover&amp;view=table">cover</a>
              </th>
              <th alt="title" class="header field title" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=title&amp;view=table">title</a>
              </th>
              <th alt="author" class="header field author" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=author&amp;view=table">author</a>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=isbn&amp;view=table">isbn</a>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=isbn13&amp;view=table">isbn13</a>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=asin&amp;view=table">asin</a>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=num_pages&amp;view=table">num pages</a>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=avg_rating&amp;view=table">avg rating</a>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=num_ratings&amp;view=table">num ratings</a>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_pub&amp;view=table">date pub</a>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_pub_edition&amp;view=table">date pub (ed.)</a>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=rating&amp;view=table">rating</a>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=shelves&amp;view=table">shelves</a>
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=review&amp;view=table">review</a>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=notes&amp;view=table">notes</a>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=comments&amp;view=table">comments</a>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=votes&amp;view=table">votes</a>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=read_count&amp;view=table">read count</a>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_started&amp;view=table">date started</a>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_read&amp;view=table">date read</a>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=20&amp;shelf=to-read&amp;sort=date_added&amp;view=table">date</a>
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=20&amp;shelf=to-read&amp;sort=date_added&amp;view=table">
                      <nobr>
                        added <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                      </nobr>
</a>              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=owned&amp;view=table">owned</a>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=format&amp;view=table">format</a>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_7064093266" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7064093266]" id="checkbox_review_7064093266" value="7064093266" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6736083077" name="positions[6736083077]" value="44">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6736083077'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6736083077&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6736083077').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6736083077').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    745,467
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
                <a id="loading_link_788070" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/45047384&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_788070&#39;);$(&#39;loading_link_788070&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_788070&#39;).show();;Element.hide(&#39;loading_anim_788070&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_788070&#39;);Element.hide(&#39;loading_link_788070&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_788070&#39;);Element.show(&#39;loading_link_788070&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;NwItDvfpM7mcTr+RhDmJ9xivh5MIgdpEcpA0IXioDYMkeVE2tdstKDXVSlBhhIsbWOQ6701PyuG4FDOJBW8mcQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_788070" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7064093266">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The House in the Cerulean Sea from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/45047384?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_7009239588" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7009239588]" id="checkbox_review_7009239588" value="7009239588" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6679332637" name="positions[6679332637]" value="43">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6679332637'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6679332637&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6679332637').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6679332637').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    14,159
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
                <a id="loading_link_722136" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/127278666&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_722136&#39;);$(&#39;loading_link_722136&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_722136&#39;).show();;Element.hide(&#39;loading_anim_722136&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_722136&#39;);Element.hide(&#39;loading_link_722136&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_722136&#39;);Element.show(&#39;loading_link_722136&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;P4Bympa3CEz6g0GF03I50Z/0/z17CBrsHR8lSfBgRhMs+w6i1IUW3VMYtEQ2zzs9379CQT7GCknXmyLhjadt4Q==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_722136" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7009239588">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Fox Wife from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/127278666?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_7008916305" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[7008916305]" id="checkbox_review_7008916305" value="7008916305" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6678993824" name="positions[6678993824]" value="42">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6678993824'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6678993824&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6678993824').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6678993824').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    40,450
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 13, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 13, 2024
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="156009464" data-user-id="68156753" data-submit-url="/review/rate/156009464?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage156009464_68156753"></span>
        <span id="successMessage156009464_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_156009464"><span id="shelf_6678993824"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 156009464, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/156009464?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 156009464, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 156009464, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 156009464, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 16, 2024">
    Nov 16, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/169348179">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_557713" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/156009464&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_557713&#39;);$(&#39;loading_link_557713&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_557713&#39;).show();;Element.hide(&#39;loading_anim_557713&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_557713&#39;);Element.hide(&#39;loading_link_557713&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_557713&#39;);Element.show(&#39;loading_link_557713&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;onAbQf0OHXQ+NtTuvy0K4Jiyc0DJPN0hlAx977DdnMKxC2d5vzwD5ZetIS9akAgM2PnOPIzyzYReiHpHzRq3MA==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_557713" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/7008916305">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Book of Doors from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/156009464?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6953601371" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6953601371]" id="checkbox_review_6953601371" value="6953601371" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6621329541" name="positions[6621329541]" value="41">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6621329541'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6621329541&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6621329541').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6621329541').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    12,550
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 04, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 09, 2024
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="195608705" data-user-id="68156753" data-submit-url="/review/rate/195608705?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage195608705_68156753"></span>
        <span id="successMessage195608705_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_195608705"><span id="shelf_6621329541"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 195608705, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/195608705?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 195608705, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 195608705, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 195608705, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="October 25, 2024">
    Oct 25, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/91916832">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_297002" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/195608705&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_297002&#39;);$(&#39;loading_link_297002&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_297002&#39;).show();;Element.hide(&#39;loading_anim_297002&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_297002&#39;);Element.hide(&#39;loading_link_297002&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_297002&#39;);Element.show(&#39;loading_link_297002&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;O4kYIKJhN64ewFVc8mIY2ABM6ubrHzMqMTOrQ6dd4q0o8mQY4FMpP7dboJ0X3xo0QAdXmq7RI4/7t6zr2prJXw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_297002" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6953601371">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Argylle from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/195608705?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6857169809" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6857169809]" id="checkbox_review_6857169809" value="6857169809" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6522465055" name="positions[6522465055]" value="40">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6522465055'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6522465055&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6522465055').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6522465055').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    186,940
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Sep 24, 2013
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 24, 2013
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="17182126" data-user-id="68156753" data-submit-url="/review/rate/17182126?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage17182126_68156753"></span>
        <span id="successMessage17182126_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_17182126"><span id="shelf_6522465055"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 17182126, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/17182126?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 17182126, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 17182126, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 17182126, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 18, 2024">
    Sep 18, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/21366540">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_44849" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/17182126&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_44849&#39;);$(&#39;loading_link_44849&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_44849&#39;).show();;Element.hide(&#39;loading_anim_44849&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_44849&#39;);Element.hide(&#39;loading_link_44849&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_44849&#39;);Element.show(&#39;loading_link_44849&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;gfi7OT7quwLivMi3ga2xSo995MzlJIYhQtllZThAMiGSg8cBfNilk0snPXZkELOmzzZZsKDqloSIXWLNRYcZ0w==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_44849" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6857169809">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Steelheart from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/17182126?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6824061184" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6824061184]" id="checkbox_review_6824061184" value="6824061184" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6488398786" name="positions[6488398786]" value="39">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6488398786'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6488398786&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6488398786').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6488398786').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="28493290" data-user-id="68156753" data-submit-url="/review/rate/28493290?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage28493290_68156753"></span>
        <span id="successMessage28493290_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_28493290"><span id="shelf_6488398786"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 28493290, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/28493290?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 28493290, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 28493290, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 28493290, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 06, 2024">
    Sep 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        ebook
            <a class="smallText" href="/work/editions/48643348">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_325551" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/28493290&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_325551&#39;);$(&#39;loading_link_325551&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_325551&#39;).show();;Element.hide(&#39;loading_anim_325551&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_325551&#39;);Element.hide(&#39;loading_link_325551&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_325551&#39;);Element.show(&#39;loading_link_325551&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;FY4Eb7t2oAiFMiagpuQU1PQ1bzSzfDyVaY3L2gvtUwsG9XhX+US+mSyp02FDWRY4tH7SSPayLDCjCcxydip4+Q==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_325551" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6824061184">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Lusty Argonian Maid Vol 1 from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/28493290?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6824059304" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6824059304]" id="checkbox_review_6824059304" value="6824059304" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6488396790" name="positions[6488396790]" value="38">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6488396790'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6488396790&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6488396790').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6488396790').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    499,846
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 26, 1985
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      2001
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="3187658" data-user-id="68156753" data-submit-url="/review/rate/3187658?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage3187658_68156753"></span>
        <span id="successMessage3187658_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_3187658"><span id="shelf_6488396790"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 3187658, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/3187658?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 3187658, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 3187658, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 3187658, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 06, 2024">
    Sep 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/2977727">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_634926" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/3187658&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_634926&#39;);$(&#39;loading_link_634926&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_634926&#39;).show();;Element.hide(&#39;loading_anim_634926&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_634926&#39;);Element.hide(&#39;loading_link_634926&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_634926&#39;);Element.show(&#39;loading_link_634926&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;hfCyHGZDvSBL+rDHS3880XJBDIlVbH9PIAE4ka/sqQOWi84kJHGjseJhRQauwj49Mgqx9RCib+rqhT850iuC8Q==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_634926" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6824059304">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Het parfum from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/3187658?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6797714160" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6797714160]" id="checkbox_review_6797714160" value="6797714160" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6461344092" name="positions[6461344092]" value="37">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6461344092'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6461344092&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6461344092').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6461344092').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
        1,344
        <span class="greyText">pp</span>
      </nobr>
</div></td>  <td class="field avg_rating"><label>avg rating</label><div class="value">    4.74
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    4,620
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Dec 06, 2024
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Dec 06, 2024
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="203578847" data-user-id="68156753" data-submit-url="/review/rate/203578847?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage203578847_68156753"></span>
        <span id="successMessage203578847_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_203578847"><span id="shelf_6461344092"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 203578847, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/203578847?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 203578847, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 203578847, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 203578847, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 28, 2024">
    Aug 28, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/23840276">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_803875" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/203578847&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_803875&#39;);$(&#39;loading_link_803875&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_803875&#39;).show();;Element.hide(&#39;loading_anim_803875&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_803875&#39;);Element.hide(&#39;loading_link_803875&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_803875&#39;);Element.show(&#39;loading_link_803875&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;z1ndLbDYD6+1cRsN5b5AycV+mt6Rbk3j1dZqWKSJD1DcIqEV8uoRPhzq7swAA0IlhTUnotSgXUYfUm3w2U4kog==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_803875" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6797714160">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Wind and Truth from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/203578847?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6734110148" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6734110148]" id="checkbox_review_6734110148" value="6734110148" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6397015054" name="positions[6397015054]" value="36">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6397015054'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6397015054&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6397015054').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6397015054').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    124,392
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 20, 2016
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 17, 2017
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="34703445" data-user-id="68156753" data-submit-url="/review/rate/34703445?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage34703445_68156753"></span>
        <span id="successMessage34703445_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_34703445"><span id="shelf_6397015054"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 34703445, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/34703445?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 34703445, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 34703445, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 34703445, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2024">
    Aug 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/54097500">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_664443" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/34703445&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_664443&#39;);$(&#39;loading_link_664443&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_664443&#39;).show();;Element.hide(&#39;loading_anim_664443&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_664443&#39;);Element.hide(&#39;loading_link_664443&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_664443&#39;);Element.show(&#39;loading_link_664443&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;bIIuMd0/5LGQr4sKToKaRl+NG3vfj+w+fl2aHNUHF1l/+VIJnw36IDk0fsurP5iqH8amB5pB/Ju02Z20qMA8qw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_664443" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6734110148">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Edgedancer from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/34703445?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6649338122" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6649338122]" id="checkbox_review_6649338122" value="6649338122" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6310407744" name="positions[6310407744]" value="35">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6310407744'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6310407744&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6310407744').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6310407744').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    531,117
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 28, 2014
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 28, 2014
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="15839976" data-user-id="68156753" data-submit-url="/review/rate/15839976?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage15839976_68156753"></span>
        <span id="successMessage15839976_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_15839976"><span id="shelf_6310407744"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 15839976, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/15839976?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 15839976, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 15839976, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 15839976, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="July 07, 2024">
    Jul 07, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/21580644">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_965922" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/15839976&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_965922&#39;);$(&#39;loading_link_965922&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_965922&#39;).show();;Element.hide(&#39;loading_anim_965922&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_965922&#39;);Element.hide(&#39;loading_link_965922&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_965922&#39;);Element.show(&#39;loading_link_965922&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;fBg0iUDmF38T5TIAhYw/IYNODTGrV5Z3JJCtlC9wCc1vY0ixAtQJ7rp+x8FgMT3NwwWwTe6ZhtLuFKo8UrciPw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_965922" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6649338122">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Red Rising from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/15839976?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6612400466" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6612400466]" id="checkbox_review_6612400466" value="6612400466" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6272463068" name="positions[6272463068]" value="34">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6272463068'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6272463068&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6272463068').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6272463068').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    59,141
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 12, 2014
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 12, 2014
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="41021196" data-user-id="68156753" data-submit-url="/review/rate/41021196?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage41021196_68156753"></span>
        <span id="successMessage41021196_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_41021196"><span id="shelf_6272463068"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 41021196, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/41021196?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41021196, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41021196, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41021196, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="June 24, 2024">
    Jun 24, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/26474462">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_981891" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/41021196&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_981891&#39;);$(&#39;loading_link_981891&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_981891&#39;).show();;Element.hide(&#39;loading_anim_981891&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_981891&#39;);Element.hide(&#39;loading_link_981891&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_981891&#39;);Element.show(&#39;loading_link_981891&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;cDSYe5xH33JpKela5uL7TCAewD1JG4of7gShAAeIWI9jT+RD3nXB48CyHJsDX/mgYFV9QQzVmrokgKaoek9zfQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_981891" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6612400466">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Fool&#39;s Assassin from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/41021196?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6512312226" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6512312226]" id="checkbox_review_6512312226" value="6512312226" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6168988194" name="positions[6168988194]" value="33">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6168988194'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6168988194&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6168988194').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6168988194').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    20,607
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 16, 2019
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 16, 2019
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="40917496" data-user-id="68156753" data-submit-url="/review/rate/40917496?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage40917496_68156753"></span>
        <span id="successMessage40917496_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_40917496"><span id="shelf_6168988194"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 40917496, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/40917496?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 40917496, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 40917496, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 40917496, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 17, 2024">
    May 17, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/63719739">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_111237" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/40917496&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_111237&#39;);$(&#39;loading_link_111237&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_111237&#39;).show();;Element.hide(&#39;loading_anim_111237&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_111237&#39;);Element.hide(&#39;loading_link_111237&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_111237&#39;);Element.show(&#39;loading_link_111237&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;GzVCgWRfetWKt5fp8py9Nm+7UoMlZqhjymvSsfRqqVgITj65Jm1kRCMsYigXIb/aL/Dv/2CouMYA79UZia2Cqg==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_111237" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6512312226">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Master and Apprentice from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/40917496?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6512311635" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6512311635]" id="checkbox_review_6512311635" value="6512311635" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6168987560" name="positions[6168987560]" value="32">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6168987560'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6168987560&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6168987560').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6168987560').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    23,821
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jul 07, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jul 07, 2015
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="23277298" data-user-id="68156753" data-submit-url="/review/rate/23277298?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage23277298_68156753"></span>
        <span id="successMessage23277298_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_23277298"><span id="shelf_6168987560"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 23277298, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/23277298?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 23277298, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 23277298, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 23277298, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 17, 2024">
    May 17, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/44980862">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_726695" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/23277298&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_726695&#39;);$(&#39;loading_link_726695&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_726695&#39;).show();;Element.hide(&#39;loading_anim_726695&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_726695&#39;);Element.hide(&#39;loading_link_726695&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_726695&#39;);Element.show(&#39;loading_link_726695&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;ZcJCpN1JvvQ3DNMR51fohQU749QCBltqXh57hiIeEZ52uT6cn3ugZZ6XJtAC6uppRXBeqEfIS8+UmnwuX9k6bA==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_726695" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6512311635">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Dark Disciple from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/23277298?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6468172262" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6468172262]" id="checkbox_review_6468172262" value="6468172262" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6123777066" name="positions[6123777066]" value="31">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6123777066'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6123777066&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6123777066').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6123777066').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="70137" data-user-id="68156753" data-submit-url="/review/rate/70137?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage70137_68156753"></span>
        <span id="successMessage70137_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_70137"><span id="shelf_6123777066"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 70137, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/70137?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 70137, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 70137, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 70137, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 30, 2024">
    Apr 30, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/67954">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_840054" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/70137&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_840054&#39;);$(&#39;loading_link_840054&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_840054&#39;).show();;Element.hide(&#39;loading_anim_840054&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_840054&#39;);Element.hide(&#39;loading_link_840054&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_840054&#39;);Element.show(&#39;loading_link_840054&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;WvwkCZBCMf6ww4uLhlUKx1HQkJxRpNEkgEWXRzFABEdJh1gx0nAvbxlYfkpj6AgrEZst4BRqwYFKwZDvTIcvtQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_840054" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6468172262">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Enterprise Architecture As Strategy from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/70137?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6462989115" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6462989115]" id="checkbox_review_6462989115" value="6462989115" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6118398373" name="positions[6118398373]" value="30">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6118398373'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6118398373&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6118398373').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6118398373').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    1,831
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 1982
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 08, 2016
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25986983" data-user-id="68156753" data-submit-url="/review/rate/25986983?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25986983_68156753"></span>
        <span id="successMessage25986983_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_25986983"><span id="shelf_6118398373"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 25986983, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/25986983?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25986983, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25986983, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25986983, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 28, 2024">
    Apr 28, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/2013294">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_234423" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/25986983&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_234423&#39;);$(&#39;loading_link_234423&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_234423&#39;).show();;Element.hide(&#39;loading_anim_234423&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_234423&#39;);Element.hide(&#39;loading_link_234423&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_234423&#39;);Element.show(&#39;loading_link_234423&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;XYRuDQJd9h6shusRIR9Ezs7UdYMhW9ggW83J4ztwiLVO/xI1QG/ojwUdHtDEokYijp/I/2SVyIWRSc5LRrejRw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_234423" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6462989115">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Dawn from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/25986983?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6450308065" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6450308065]" id="checkbox_review_6450308065" value="6450308065" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6105476219" name="positions[6105476219]" value="29">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6105476219'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6105476219&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6105476219').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6105476219').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    348,141
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 01, 1937
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 2016
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="30186948" data-user-id="68156753" data-submit-url="/review/rate/30186948?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage30186948_68156753"></span>
        <span id="successMessage30186948_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_30186948"><span id="shelf_6105476219"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 30186948, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/30186948?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 30186948, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 30186948, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 30186948, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 23, 2024">
    Apr 23, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/1199320">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_601985" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/30186948&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_601985&#39;);$(&#39;loading_link_601985&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_601985&#39;).show();;Element.hide(&#39;loading_anim_601985&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_601985&#39;);Element.hide(&#39;loading_link_601985&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_601985&#39;);Element.show(&#39;loading_link_601985&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;Xqmzj0XTmIdrR9m5m4er3XqR8vjION1uyqDkqZGHqa1N0s+3B+GGFsLcLHh+OqkxOtpPhI32zcsAJOMB7ECCXw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_601985" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6450308065">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Think and Grow Rich from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/30186948?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6440467156" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6440467156]" id="checkbox_review_6440467156" value="6440467156" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_6095351321" name="positions[6095351321]" value="28">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_6095351321'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_6095351321&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_6095351321').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_6095351321').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    142,215
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jun 04, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 04, 2015
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25499718" data-user-id="68156753" data-submit-url="/review/rate/25499718?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25499718_68156753"></span>
        <span id="successMessage25499718_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_25499718"><span id="shelf_6095351321"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 25499718, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/25499718?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25499718, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25499718, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25499718, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 19, 2024">
    Apr 19, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/45276208">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_565079" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/25499718&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_565079&#39;);$(&#39;loading_link_565079&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_565079&#39;).show();;Element.hide(&#39;loading_anim_565079&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_565079&#39;);Element.hide(&#39;loading_link_565079&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_565079&#39;);Element.show(&#39;loading_link_565079&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;l+nMQ/BJuXcKFc4FzazTXncbfzVI9T07ROgTs3acZviEkrB7snun5qOOO8QoEdGyN1DCSQ07LZ6ObBQbC1tNCg==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_565079" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6440467156">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Children of Time from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/25499718?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6240930264" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6240930264]" id="checkbox_review_6240930264" value="6240930264" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5890357470" name="positions[5890357470]" value="27">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5890357470'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5890357470&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5890357470').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5890357470').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    18,141
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Sep 01, 2012
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 25, 2012
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="13616278" data-user-id="68156753" data-submit-url="/review/rate/13616278?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage13616278_68156753"></span>
        <span id="successMessage13616278_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_13616278"><span id="shelf_5890357470"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 13616278, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/13616278?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 13616278, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 13616278, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 13616278, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/19217996">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_223157" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/13616278&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_223157&#39;);$(&#39;loading_link_223157&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_223157&#39;).show();;Element.hide(&#39;loading_anim_223157&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_223157&#39;);Element.hide(&#39;loading_link_223157&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_223157&#39;);Element.show(&#39;loading_link_223157&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;LTY71pt6PN1vq5+GIDZ6AMAEm0gqALKwTyyKohC17Hw+TUfu2UgiTMYwakfFi3jsgE8mNG/OohWFqI0KbXLHjg==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_223157" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6240930264">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Red Knight from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/13616278?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6240911185" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6240911185]" id="checkbox_review_6240911185" value="6240911185" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5890337820" name="positions[5890337820]" value="26">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5890337820'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5890337820&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5890337820').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5890337820').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    11,878
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 21, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 21, 2020
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="51277288" data-user-id="68156753" data-submit-url="/review/rate/51277288?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage51277288_68156753"></span>
        <span id="successMessage51277288_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_51277288"><span id="shelf_5890337820"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 51277288, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/51277288?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 51277288, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 51277288, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 51277288, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/66944898">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_461091" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/51277288&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_461091&#39;);$(&#39;loading_link_461091&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_461091&#39;).show();;Element.hide(&#39;loading_anim_461091&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_461091&#39;);Element.hide(&#39;loading_link_461091&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_461091&#39;);Element.show(&#39;loading_link_461091&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;9NeA8eEVT+jaa2kCDKS/BBR7kFPsvWBtTQ/dJMxV8HbnrPzJoydReXPwnMPpGb3oVDAtL6lzcMiHi9qMsZLbhA==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_461091" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6240911185">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Girl and the Stars from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/51277288?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6240882560" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6240882560]" id="checkbox_review_6240882560" value="6240882560" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5890307947" name="positions[5890307947]" value="25">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5890307947'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5890307947&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5890307947').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5890307947').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    14,348
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 24, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 24, 2021
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="56229688" data-user-id="68156753" data-submit-url="/review/rate/56229688?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage56229688_68156753"></span>
        <span id="successMessage56229688_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_56229688"><span id="shelf_5890307947"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 56229688, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/56229688?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 56229688, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 56229688, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 56229688, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 06, 2024">
    Feb 06, 2024
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        ebook
            <a class="smallText" href="/work/editions/87571717">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_823392" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/56229688&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_823392&#39;);$(&#39;loading_link_823392&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_823392&#39;).show();;Element.hide(&#39;loading_anim_823392&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_823392&#39;);Element.hide(&#39;loading_link_823392&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_823392&#39;);Element.show(&#39;loading_link_823392&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;FhdnWl3oFP+vjLLLtNiMW6aIbcS4XLtZr7CyWB9MUyMFbBtiH9oKbgYXRwpRZY635sPQuP2Sq/xlNLXwYot40Q==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_823392" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6240882560">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Pariah from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/56229688?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D1%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

</tbody></table>    </div>`
	mockHTML2 := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage">
<body class="">
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".rcmnqo7stk" data-react-checksum="-1301868141"></noscript></div>

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
        state: 'apple_oauth_state_c2e062ea-7ef8-432d-b3a2-ac92a54eaeef'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".pqs1bs16vi" data-react-checksum="-460434422"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".pqs1bs16vi.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".pqs1bs16vi.0.0"><picture data-reactid=".pqs1bs16vi.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".pqs1bs16vi.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".pqs1bs16vi.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".pqs1bs16vi.1"><div class="siteHeader__contents" data-reactid=".pqs1bs16vi.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".pqs1bs16vi.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".pqs1bs16vi.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".pqs1bs16vi.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".pqs1bs16vi.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".pqs1bs16vi.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".pqs1bs16vi.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".pqs1bs16vi.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".pqs1bs16vi.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".pqs1bs16vi.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".pqs1bs16vi.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.0"><span data-reactid=".pqs1bs16vi.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".pqs1bs16vi.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".pqs1bs16vi.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".pqs1bs16vi.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.0"><span data-reactid=".pqs1bs16vi.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".pqs1bs16vi.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".pqs1bs16vi.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".pqs1bs16vi.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".pqs1bs16vi.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".pqs1bs16vi.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".pqs1bs16vi.1.0.4"><ul class="personalNav" data-reactid=".pqs1bs16vi.1.0.4.0"><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.4.0.0"><div data-reactid=".pqs1bs16vi.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".pqs1bs16vi.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".pqs1bs16vi.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".pqs1bs16vi.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".pqs1bs16vi.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".pqs1bs16vi.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.0"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".pqs1bs16vi.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".pqs1bs16vi.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".pqs1bs16vi.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".pqs1bs16vi.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".pqs1bs16vi.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".pqs1bs16vi.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".pqs1bs16vi.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".pqs1bs16vi.1.0.6"><div class="modal__content" data-reactid=".pqs1bs16vi.1.0.6.0"><div class="modal__close" data-reactid=".pqs1bs16vi.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".pqs1bs16vi.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".pqs1bs16vi.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".pqs1bs16vi.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".pqs1bs16vi.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".pqs1bs16vi.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".pqs1bs16vi.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".pqs1bs16vi.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".pqs1bs16vi.1.0.7"><div data-reactid=".pqs1bs16vi.1.0.7.0"><div class="modal__close" data-reactid=".pqs1bs16vi.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".pqs1bs16vi.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".pqs1bs16vi.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".pqs1bs16vi.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".pqs1bs16vi.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1"><ul data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".pqs1bs16vi.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".pqs1bs16vi.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".pqs1bs16vi.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".pqs1bs16vi.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".pqs1bs16vi.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".pqs1bs16vi.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".pqs1bs16vi.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".pqs1bs16vi.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".pqs1bs16vi.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".pqs1bs16vi.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".pqs1bs16vi.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.0"><span data-reactid=".pqs1bs16vi.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".pqs1bs16vi.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".pqs1bs16vi.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".pqs1bs16vi.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.0"><span data-reactid=".pqs1bs16vi.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".pqs1bs16vi.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
        <a href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=table">My Books</a>: 
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
<a id="batchEditLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this, {afterOpen: startEditing, afterClose: stopEditing});; return false;">Batch Edit</a>
</li>
<li>
<a id="shelfSettingsLink" class="actionLinkLite controlLink" href="#" onclick="toggleControl(this); return false;">Settings</a>
</li>
<li>
<a class="actionLinkLite controlLink" href="/review/stats/68156753">Stats</a>
</li>
<li>
<a class="actionLinkLite controlLink" target="_blank" rel="noopener noreferrer" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=20&amp;print=true&amp;shelf=to-read&amp;view=table">Print</a>
</li>
<li>
<span class="greyText">&nbsp;|&nbsp;</span>
</li>
<li>
<a class="listViewIcon selected" href="/review/list/68156753-sebastiaan?page=2&amp;per_page=20&amp;shelf=to-read&amp;view=table"><img title="table view" alt="table view" src="https://s.gr-assets.com/assets/layout/list-fe412c89a6a612c841b5b58681660b82.png" /></a>
</li>
<li>
<a class="gridViewIcon " href="/review/list/68156753-sebastiaan?page=2&amp;per_page=20&amp;shelf=to-read&amp;view=covers"><img title="cover view" alt="cover view" src="https://s.gr-assets.com/assets/layout/grid-2c030bffe1065f73ddca41540e8a267d.png" /></a>
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
            <a class="actionLinkLite" href="/review/list/68156753?shelf=%23ALL%23">All (367)</a>
            <div id="paginatedShelfList" class="stacked">
                <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Read&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read%2Cread&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Read shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=read&amp;view=table">Read  &lrm;(317)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" style="display: none" rel="nofollow" title="View books on your &quot;to-read&quot;, and &quot;Currently Reading&quot; shelves" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read%2Ccurrently-reading&amp;view=table">+</a>
    <a title="Sebastiaan&#39;s Currently Reading shelf" class="actionLinkLite" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=currently-reading&amp;view=table">Currently Reading  &lrm;(0)</a>
  </div>
  <div class="userShelf">
        <a class="greyText right multiLink" rel="nofollow" style="display: none" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=&amp;view=table">&minus;</a>
    <a title="Sebastiaan&#39;s Want to Read shelf" class="selectedShelf" href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;view=table">Want to Read  &lrm;(44)</a>
  </div>



            </div>
            <div class="stacked">
                <a class="actionLink" href="#" onclick="$$(&#39;#paginatedShelfList .multiLink&#39;).invoke(&#39;toggle&#39;); return false;">select multiple</a>
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
      <form id="fieldsForm" class="edit_user_shelf" action="/shelf/update/222519907" accept-charset="UTF-8" data-remote="true" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="patch" />        <table>
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
            <input type="submit" name="commit" value="Save Current Settings to Your &quot;to-read&quot; shelf" id="save_curr_sett_submit" class="gr-button gr-button--small" style="margin-right: 10px" />
            <span class="loading" style="display: none"><img src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /> Saving...</span>
              <span class="greyText inter uitext">shelf settings customized</span>
              <input type="checkbox" name="reset_display_fields" id="reset_display_fields" value="true" style="display:none" />
            <span class="greyText status inter"></span>
          </div>
</form>      <a class="actionLinkLite greyText smallText right" href="#" onclick="hideControl($(&#39;shelfSettingsLink&#39;)); return false;">close</a>
      <div class="clear"></div>
    </div>
      <div id="batchEdit" style="display: none;" class="controlBody">
        <div id="shelfTools" class="toolset">
          <form name="reviewEditForm" id="reviewEditForm" action="/review/update_list/68156753" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="kw3cQq8lSyEuVd0gq/vtdGUoADHbeVjrLuVR0jLwldSAdqB67RdVsIfOKOFORu+YJWO9TZ63SE7kYVZ6Tze+Jg==" />
            <input type="hidden" name="view" id="view" value="table" />
            <label>
              Shelf:
              <select name="edit[shelf]" id="edit_shelf"><option value="read">read</option>
<option value="currently-reading">currently-reading</option>
<option value="to-read">to-read</option></select>
              &nbsp;
            </label>
            <a id="add_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;add_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_564711&#39;);$(&#39;add_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;add_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_564711&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_564711&#39;);Element.hide(&#39;add_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_564711&#39;);Element.show(&#39;add_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=add_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;poiTYMYnhCZDQFd04ARRVzixm9VD7ZcuP0MXuMb9vPW18+9YhBWat+rborUFuVO7ePomqQYjh4v1xxAQuzqXBw==&#39;)}); return false;">add books to this shelf</a><img style="display:none" id="loading_anim_564711" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_shelves_link" class="actionLink" href="#" onclick="new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_shelves_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_947564&#39;);$(&#39;remove_shelves_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_shelves_link&#39;).show();;Element.hide(&#39;loading_anim_947564&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_947564&#39;);Element.hide(&#39;remove_shelves_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_947564&#39;);Element.show(&#39;remove_shelves_link&#39;);for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {var r = request.responseJSON.reviews[i];$(&#39;review_&#39;+r.object.id).replace(r.html);$(&#39;review_&#39;+r.object.id).labelize({force: true, hoverClass: &#39;checkable&#39;, selectedClass: &#39;selected&#39;});};toggleFieldsToMatchHeader();alert(request.responseJSON.msg)}, parameters:(&#39;form_action=remove_shelves&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;Xw3rBPLaurv2VSj+pXeA0X6DYCX5K1rXGpzy1x9CWOpMdpc8sOikKl/O3T9AyoI9PsjdWbzlSnLQGPV/YoVzGA==&#39;)}); return false;">remove books from this shelf</a><img style="display:none" id="loading_anim_947564" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
            |
            <a id="remove_books_link" class="actionLinkLite" href="#" onclick="if (confirm(&#39;This will completely remove the selected books from your shelves.&#39;)) { new Ajax.Request(&#39;/review/update_list/68156753&#39;, {asynchronous:true, evalScripts:true, on422:function(request){$$(&#39;#batchEdit .loading&#39;).invoke(&#39;hide&#39;);$(&#39;remove_books_link&#39;).show();alert(request.responseText);}, onFailure:function(request){Element.hide(&#39;loading_anim_182657&#39;);$(&#39;remove_books_link&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;remove_books_link&#39;).show();;Element.hide(&#39;loading_anim_182657&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_182657&#39;);Element.hide(&#39;remove_books_link&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_182657&#39;);Element.show(&#39;remove_books_link&#39;);              for (var i = request.responseJSON.reviews.length - 1; i &gt;= 0; i--) {
                var r = request.responseJSON.reviews[i];
                $(&#39;review_&#39;+r.object.id).fade();
              }
}, parameters:(&#39;form_action=remove_books&amp;&#39; + Form.serializeElements($$(&#39;#books .checkbox input&#39;).concat($$(&#39;#batchEdit select&#39;), $$(&#39;#batchEdit input&#39;)))) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;z9m6Uj00H+hx37OENKcbnSq3hgqAeUJLAyMVb2re0SXcosZqfwYBedhERkXRGhlxavw7dsW3Uu7JpxLHFxn61w==&#39;)}); }; return false;">remove books from all shelves</a><img style="display:none" id="loading_anim_182657" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
</form>        </div>
        <div id="otherTools" class="toolset greyText">
          <div class="right">
                  <a class="actionLinkLite smallText" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;sort=position&amp;view=table">sort by position to re-order</a>
                  <span class="greyText">|</span>
                <a method="post" id="loading_link_61126020" class="actionLinkLite smallText" href="#" onclick="if (confirm(&#39;Are you sure you want to disable sorting?  This will remove the custom order you\&#39;ve applied to this shelf.&#39;)) { $(this).hide(); $(&#39;loading_anim_61126020&#39;).show(); $(&#39;hidden_link_61126020&#39;).simulate(&#39;click&#39;); }; return false;">disable sorting</a><img style="display:none" id="loading_anim_61126020" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" /><a id="hidden_link_61126020" class="actionLinkLite smallText" style="display: none" rel="nofollow" data-method="post" href="/shelf/disable_sorting/222519907">disable sorting</a>
                <span class="greyText">|</span>
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
        <a id="loading_link_41127" class="button" href="#" onclick="new Ajax.Request(&#39;/shelf/move_batch/68156753&#39;, {asynchronous:true, evalScripts:true, onComplete:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;hide&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;show&#39;);}, onFailure:function(request){alert(&#39;Something went wrong re-ordering those shelves.&#39;);;Element.hide(&#39;loading_anim_41127&#39;);}, onLoading:function(request){$$(&#39;#books .position .position_loading&#39;).invoke(&#39;show&#39;);$$(&#39;#books .position input&#39;).invoke(&#39;hide&#39;);;Element.show(&#39;loading_anim_41127&#39;);Element.hide(&#39;loading_link_41127&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_41127&#39;);Element.show(&#39;loading_link_41127&#39;);$(&#39;booksBody&#39;).update(request.responseJSON.html);toggleFieldsToMatchHeader();startEditing();$(&#39;reorderConfirm&#39;).hide();$(&#39;booksBody&#39;).highlight();}, parameters:Form.serializeElements($$(&#39;#books .position input&#39;)) + &#39;&amp;authenticity_token=&#39; + encodeURIComponent(&#39;bxsz2saUP+K9/oSvDGmnf3z9VAP5oRRi1yPqvtl/16d8YE/ihKYhcxRlcW7p1KWTPLbpf7xvBMcdp+0WpLj8VQ==&#39;)}); return false;">apply position changes?</a><img style="display:none" id="loading_anim_41127" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          &nbsp;
          <a href="#" onclick="$(&#39;reorderConfirm&#39;).hide(); return false;">Not yet</a>
      </div>
      <div class="right uitext">
        <div id="reviewPagination"><a class="previous_page" rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=table">« previous</a> <a rel="prev start" href="/review/list/68156753-sebastiaan?page=1&amp;per_page=20&amp;shelf=to-read&amp;view=table">1</a> <em class="current">2</em> <a rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;per_page=20&amp;shelf=to-read&amp;view=table">3</a> <a class="next_page" rel="next" href="/review/list/68156753-sebastiaan?page=3&amp;per_page=20&amp;shelf=to-read&amp;view=table">next »</a></div>

      </div>
      <div class="clear"></div>
    <div class="js-dataTooltip" data-use-wtr-tooltip="true">
      <table id="books" class="table stacked" border="0">
        <thead>
          <tr id="booksHeader" class="tableList">
              <th alt="checkbox" class="header field checkbox" style="display: none">
              </th>
              <th alt="position" class="header field position" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=position&amp;view=table">#</a>
              </th>
              <th alt="cover" class="header field cover" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=cover&amp;view=table">cover</a>
              </th>
              <th alt="title" class="header field title" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=title&amp;view=table">title</a>
              </th>
              <th alt="author" class="header field author" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=author&amp;view=table">author</a>
              </th>
              <th alt="isbn" class="header field isbn" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=isbn&amp;view=table">isbn</a>
              </th>
              <th alt="isbn13" class="header field isbn13" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=isbn13&amp;view=table">isbn13</a>
              </th>
              <th alt="asin" class="header field asin" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=asin&amp;view=table">asin</a>
              </th>
              <th alt="num_pages" class="header field num_pages" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=num_pages&amp;view=table">num pages</a>
              </th>
              <th alt="avg_rating" class="header field avg_rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=avg_rating&amp;view=table">avg rating</a>
              </th>
              <th alt="num_ratings" class="header field num_ratings" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=num_ratings&amp;view=table">num ratings</a>
              </th>
              <th alt="date_pub" class="header field date_pub" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_pub&amp;view=table">date pub</a>
              </th>
              <th alt="date_pub_edition" class="header field date_pub_edition" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_pub_edition&amp;view=table">date pub (ed.)</a>
              </th>
              <th alt="rating" class="header field rating" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=rating&amp;view=table">rating</a>
              </th>
              <th alt="shelves" class="header field shelves" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=shelves&amp;view=table">shelves</a>
              </th>
              <th alt="review" class="header field review" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=review&amp;view=table">review</a>
              </th>
              <th alt="notes" class="header field notes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=notes&amp;view=table">notes</a>
              </th>
              <th alt="recommender" class="header field recommender" style="display: none">
              </th>
              <th alt="comments" class="header field comments" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=comments&amp;view=table">comments</a>
              </th>
              <th alt="votes" class="header field votes" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=votes&amp;view=table">votes</a>
              </th>
              <th alt="read_count" class="header field read_count" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=read_count&amp;view=table">read count</a>
              </th>
              <th alt="date_started" class="header field date_started" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_started&amp;view=table">date started</a>
              </th>
              <th alt="date_read" class="header field date_read" style="">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=date_read&amp;view=table">date read</a>
              </th>
              <th alt="date_added" class="header field date_added" style="">
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=20&amp;shelf=to-read&amp;sort=date_added&amp;view=table">date</a>
                    <a href="/review/list/68156753-sebastiaan?order=a&amp;per_page=20&amp;shelf=to-read&amp;sort=date_added&amp;view=table">
                      <nobr>
                        added <img src="https://s.gr-assets.com/assets/down_arrow-1e1fa5642066c151f5e0136233fce98a.gif" alt="Down arrow" />
                      </nobr>
</a>              </th>
              <th alt="date_purchased" class="header field date_purchased" style="display: none">
              </th>
              <th alt="owned" class="header field owned" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=owned&amp;view=table">owned</a>
              </th>
              <th alt="purchase_location" class="header field purchase_location" style="display: none">
              </th>
              <th alt="condition" class="header field condition" style="display: none">
              </th>
              <th alt="format" class="header field format" style="display: none">
                    <a href="/review/list/68156753-sebastiaan?per_page=20&amp;shelf=to-read&amp;sort=format&amp;view=table">format</a>
              </th>
              <th alt="actions" class="header field actions" style="">
              </th>
          </tr>
        </thead>
        <tbody id="booksBody">
              
<tr id="review_6073082086" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6073082086]" id="checkbox_review_6073082086" value="6073082086" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5727714211" name="positions[5727714211]" value="24">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5727714211'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5727714211&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5727714211').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5727714211').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    167,983
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 24, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 24, 2020
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="26082916" data-user-id="68156753" data-submit-url="/review/rate/26082916?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage26082916_68156753"></span>
        <span id="successMessage26082916_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_26082916"><span id="shelf_5727714211"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 26082916, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/26082916?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 26082916, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 26082916, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 26082916, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 23, 2023">
    Dec 23, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/59016474">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_601639" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/26082916&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_601639&#39;);$(&#39;loading_link_601639&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_601639&#39;).show();;Element.hide(&#39;loading_anim_601639&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_601639&#39;);Element.hide(&#39;loading_link_601639&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_601639&#39;);Element.show(&#39;loading_link_601639&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;YwZpZnfLRefP4Xl54dl8/goxFIMjfscACYmrGqgJe6NwfRVeNflbdmZ6jLgEZH4SSnqp/2aw16XDDayy1c5QUQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_601639" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6073082086">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Ready Player Two from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/26082916?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6028400962" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6028400962]" id="checkbox_review_6028400962" value="6028400962" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5682793365" name="positions[5682793365]" value="23">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5682793365'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5682793365&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5682793365').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5682793365').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    59,893
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Apr 04, 2017
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 04, 2017
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25895524" data-user-id="68156753" data-submit-url="/review/rate/25895524?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25895524_68156753"></span>
        <span id="successMessage25895524_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_25895524"><span id="shelf_5682793365"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 25895524, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/25895524?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25895524, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25895524, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25895524, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 06, 2023">
    Dec 06, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/45777900">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_715397" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/25895524&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_715397&#39;);$(&#39;loading_link_715397&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_715397&#39;).show();;Element.hide(&#39;loading_anim_715397&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_715397&#39;);Element.hide(&#39;loading_link_715397&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_715397&#39;);Element.show(&#39;loading_link_715397&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;kGVQID2SsAz295iQytMm7quDDGFmKjLpOR7KfBuTsHWDHiwYf6CunV9sbVEvbiQC68ixHSPkIkzzms3UZlSbhw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_715397" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6028400962">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Red Sister from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/25895524?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6024014795" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6024014795]" id="checkbox_review_6024014795" value="6024014795" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5678278128" name="positions[5678278128]" value="22">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5678278128'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5678278128&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5678278128').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5678278128').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    2,969
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      2012
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 15, 2015
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25351052" data-user-id="68156753" data-submit-url="/review/rate/25351052?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25351052_68156753"></span>
        <span id="successMessage25351052_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_25351052"><span id="shelf_5678278128"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 25351052, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/25351052?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25351052, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25351052, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25351052, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="December 04, 2023">
    Dec 04, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/21860234">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_515026" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/25351052&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_515026&#39;);$(&#39;loading_link_515026&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_515026&#39;).show();;Element.hide(&#39;loading_anim_515026&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_515026&#39;);Element.hide(&#39;loading_link_515026&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_515026&#39;);Element.show(&#39;loading_link_515026&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;ZpyQR17fwsIPzas37PSyog4mfrIb99ylsRkH0AoCtRx15+x/HO3cU6ZWXvYJSbBOTm3Dzl45zAB7nQB4d8We7g==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_515026" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6024014795">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove De hondenmoorden from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/25351052?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6006984679" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6006984679]" id="checkbox_review_6006984679" value="6006984679" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5660920865" name="positions[5660920865]" value="21">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5660920865'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5660920865&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5660920865').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5660920865').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    131,056
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      1971
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 03, 2005
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="11487773" data-user-id="68156753" data-submit-url="/review/rate/11487773?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage11487773_68156753"></span>
        <span id="successMessage11487773_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_11487773"><span id="shelf_5660920865"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 11487773, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/11487773?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 11487773, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 11487773, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 11487773, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 27, 2023">
    Nov 27, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/823130">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_525936" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/11487773&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_525936&#39;);$(&#39;loading_link_525936&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_525936&#39;).show();;Element.hide(&#39;loading_anim_525936&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_525936&#39;);Element.hide(&#39;loading_link_525936&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_525936&#39;);Element.show(&#39;loading_link_525936&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;uRwKukhw1m6wAQteyRjI2MgjhEPip6CgGOfNnsZ87M6qZ3aCCkLI/xma/p8spco0iGg5P6dpsAXSY8o2u7vHPA==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_525936" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6006984679">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Postkantoor from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/11487773?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_6001687571" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[6001687571]" id="checkbox_review_6001687571" value="6001687571" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5655514429" name="positions[5655514429]" value="20">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5655514429'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5655514429&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5655514429').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5655514429').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    26,571
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 01, 2023
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 09, 2023
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="61612864" data-user-id="68156753" data-submit-url="/review/rate/61612864?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage61612864_68156753"></span>
        <span id="successMessage61612864_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_61612864"><span id="shelf_5655514429"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 61612864, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/61612864?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61612864, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61612864, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61612864, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="November 25, 2023">
    Nov 25, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/96381602">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_576774" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/61612864&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_576774&#39;);$(&#39;loading_link_576774&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_576774&#39;).show();;Element.hide(&#39;loading_anim_576774&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_576774&#39;);Element.hide(&#39;loading_link_576774&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_576774&#39;);Element.show(&#39;loading_link_576774&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;I+PnsMyY4m1IVDojzfFb0HeirEdmVyHhvIW48w9zde0wmJuIjqr8/OHPz+IoTFk8N+kROyOZMUR2Ab9bcrReHw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_576774" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/6001687571">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Book That Wouldn’t Burn from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/61612864?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_5707553128" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[5707553128]" id="checkbox_review_5707553128" value="5707553128" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5352246956" name="positions[5352246956]" value="19">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5352246956'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5352246956&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5352246956').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5352246956').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    21,724
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 30, 2022
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Apr 25, 2023
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="123192410" data-user-id="68156753" data-submit-url="/review/rate/123192410?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage123192410_68156753"></span>
        <span id="successMessage123192410_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_123192410"><span id="shelf_5352246956"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 123192410, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/123192410?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 123192410, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 123192410, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 123192410, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="July 20, 2023">
    Jul 20, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/96425038">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_918583" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/123192410&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_918583&#39;);$(&#39;loading_link_918583&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_918583&#39;).show();;Element.hide(&#39;loading_anim_918583&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_918583&#39;);Element.hide(&#39;loading_link_918583&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_918583&#39;);Element.show(&#39;loading_link_918583&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;70wJzvXodD6XzvsDMnjvfwhMGqUJGg2mmGA8pGlpz+X8N3X2t9pqrz5VDsLXxe2TSAen2UzUHQNS5DsMFK7kFw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_918583" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/5707553128">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Bloedmaan from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/123192410?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_5610558846" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[5610558846]" id="checkbox_review_5610558846" value="5610558846" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_5252718453" name="positions[5252718453]" value="18">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_5252718453'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_5252718453&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_5252718453').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_5252718453').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    16,259
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Dec 13, 2018
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Dec 13, 2018
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="41456850" data-user-id="68156753" data-submit-url="/review/rate/41456850?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage41456850_68156753"></span>
        <span id="successMessage41456850_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_41456850"><span id="shelf_5252718453"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 41456850, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/41456850?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41456850, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41456850, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 41456850, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="June 10, 2023">
    Jun 10, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/64719263">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_206569" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/41456850&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_206569&#39;);$(&#39;loading_link_206569&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_206569&#39;).show();;Element.hide(&#39;loading_anim_206569&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_206569&#39;);Element.hide(&#39;loading_link_206569&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_206569&#39;);Element.show(&#39;loading_link_206569&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;Znk13MJNk1mtWaznSNGbggrlQBEVGsguHVxy2tfHGjt1AknkgH+NyATCWSatbJluSq79bVDU2IvX2HVyqgAxyQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_206569" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/5610558846">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Grand Hotel Europa from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/41456850?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_5342865860" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[5342865860]" id="checkbox_review_5342865860" value="5342865860" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4974287566" name="positions[4974287566]" value="17">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4974287566'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4974287566&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4974287566').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4974287566').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    237,392
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jan 31, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Sep 08, 2022
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="61259384" data-user-id="68156753" data-submit-url="/review/rate/61259384?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage61259384_68156753"></span>
        <span id="successMessage61259384_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_61259384"><span id="shelf_4974287566"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 61259384, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/61259384?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61259384, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61259384, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 61259384, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 13, 2023">
    Feb 13, 2023
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/75499539">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_998987" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/61259384&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_998987&#39;);$(&#39;loading_link_998987&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_998987&#39;).show();;Element.hide(&#39;loading_anim_998987&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_998987&#39;);Element.hide(&#39;loading_link_998987&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_998987&#39;);Element.show(&#39;loading_link_998987&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;cnoMZBwGwGNmvRfEercFiYTHeXcBWENoEMS4tFi8YethAXBcXjTe8s8m4gWfCgdlxIzEC0SWU83aQL8cJXtKGQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_998987" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/5342865860">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Atlas Six from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/61259384?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_5003029905" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[5003029905]" id="checkbox_review_5003029905" value="5003029905" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4630160978" name="positions[4630160978]" value="16">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4630160978'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4630160978&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4630160978').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4630160978').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    3,987
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      <span class="greyText">unknown</span>
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 09, 2018
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="37436740" data-user-id="68156753" data-submit-url="/review/rate/37436740?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage37436740_68156753"></span>
        <span id="successMessage37436740_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_37436740"><span id="shelf_4630160978"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 37436740, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/37436740?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 37436740, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 37436740, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 37436740, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="September 22, 2022">
    Sep 22, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/59067287">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_439352" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/37436740&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_439352&#39;);$(&#39;loading_link_439352&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_439352&#39;).show();;Element.hide(&#39;loading_anim_439352&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_439352&#39;);Element.hide(&#39;loading_link_439352&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_439352&#39;);Element.show(&#39;loading_link_439352&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;qxEhWx7dqOIThgZbRGCxBbkLP/Wk5gsm0GwwPVWxsaG4al1jXO+2c7od85qh3bPp+UCCieEoG4Ma6DeVKHaaUw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_439352" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/5003029905">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Mistborn from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/37436740?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4927551582" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4927551582]" id="checkbox_review_4927551582" value="4927551582" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4553325793" name="positions[4553325793]" value="15">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4553325793'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4553325793&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4553325793').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4553325793').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    2,631
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Aug 05, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Aug 05, 2021
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="57596188" data-user-id="68156753" data-submit-url="/review/rate/57596188?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage57596188_68156753"></span>
        <span id="successMessage57596188_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_57596188"><span id="shelf_4553325793"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 57596188, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/57596188?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 57596188, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 57596188, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 57596188, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 18, 2022">
    Aug 18, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/86175217">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_89667" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/57596188&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_89667&#39;);$(&#39;loading_link_89667&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_89667&#39;).show();;Element.hide(&#39;loading_anim_89667&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_89667&#39;);Element.hide(&#39;loading_link_89667&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_89667&#39;);Element.show(&#39;loading_link_89667&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;HmYCus4Usw5OlXmzpFd4e0KCahZpUX26qb3i46HUPfsNHX6CjCatn+cOjHJB6nqXAsnXaiyfbR9jOeVL3BMWCQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_89667" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4927551582">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Hand of the Sun King from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/57596188?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4913927980" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4913927980]" id="checkbox_review_4913927980" value="4913927980" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4539554354" name="positions[4539554354]" value="14">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4539554354'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4539554354&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4539554354').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4539554354').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    68,082
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      2009
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 21, 2011
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="9835731" data-user-id="68156753" data-submit-url="/review/rate/9835731?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage9835731_68156753"></span>
        <span id="successMessage9835731_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_9835731"><span id="shelf_4539554354"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 9835731, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/9835731?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 9835731, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 9835731, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 9835731, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 12, 2022">
    Aug 12, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/7184205">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_9657" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/9835731&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_9657&#39;);$(&#39;loading_link_9657&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_9657&#39;).show();;Element.hide(&#39;loading_anim_9657&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_9657&#39;);Element.hide(&#39;loading_link_9657&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_9657&#39;);Element.show(&#39;loading_link_9657&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;8DczSQ0nToBUElLRBmvyq05hb/URRtyVA42CISY42ovjTE9xTxVQEf2JpxDj1vBHDirSiVSIzDDJCYWJW//xeQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_9657" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4913927980">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Hypnotist from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/9835731?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4899245497" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4899245497]" id="checkbox_review_4899245497" value="4899245497" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4524643083" name="positions[4524643083]" value="13">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4524643083'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4524643083&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4524643083').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4524643083').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    30,805
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 11, 2021
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 11, 2021
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="52504334" data-user-id="68156753" data-submit-url="/review/rate/52504334?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage52504334_68156753"></span>
        <span id="successMessage52504334_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_52504334"><span id="shelf_4524643083"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 52504334, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/52504334?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 52504334, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 52504334, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 52504334, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/78021845">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_112555" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/52504334&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_112555&#39;);$(&#39;loading_link_112555&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_112555&#39;).show();;Element.hide(&#39;loading_anim_112555&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_112555&#39;);Element.hide(&#39;loading_link_112555&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_112555&#39;);Element.show(&#39;loading_link_112555&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;RLAxv4uno9UANbTA/v0PGdPrDOq2iEA/NrC/AfkNlNhXy02HyZW9RKmuQQEbQA31k6CxlvNGUJr8NLiphMq/Kg==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_112555" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4899245497">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove A Master of Djinn from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/52504334?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4899244695" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4899244695]" id="checkbox_review_4899244695" value="4899244695" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4524642320" name="positions[4524642320]" value="12">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4524642320'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4524642320&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4524642320').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4524642320').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    24,792
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      May 17, 2022
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      May 17, 2022
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="58950705" data-user-id="68156753" data-submit-url="/review/rate/58950705?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage58950705_68156753"></span>
        <span id="successMessage58950705_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_58950705"><span id="shelf_4524642320"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 58950705, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/58950705?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 58950705, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 58950705, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 58950705, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/91801906">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_995050" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/58950705&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_995050&#39;);$(&#39;loading_link_995050&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_995050&#39;).show();;Element.hide(&#39;loading_anim_995050&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_995050&#39;);Element.hide(&#39;loading_link_995050&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_995050&#39;);Element.show(&#39;loading_link_995050&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;rMAobKgtZLcqbXP4SKyyWVfRYQSPvFdebo2fJMnPS2u/u1RU6h96JoP2hjmtEbC1F5rceMpyR/ukCZiMtAhgmQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_995050" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4899244695">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Stardust Thief from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/58950705?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4899240130" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4899240130]" id="checkbox_review_4899240130" value="4899240130" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4524637695" name="positions[4524637695]" value="11">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4524637695'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4524637695&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4524637695').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4524637695').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    48,578
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Oct 13, 2020
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Oct 13, 2020
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="50892360" data-user-id="68156753" data-submit-url="/review/rate/50892360?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage50892360_68156753"></span>
        <span id="successMessage50892360_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_50892360"><span id="shelf_4524637695"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 50892360, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/50892360?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 50892360, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 50892360, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 50892360, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="August 06, 2022">
    Aug 06, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/61321587">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_42259" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/50892360&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_42259&#39;);$(&#39;loading_link_42259&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_42259&#39;).show();;Element.hide(&#39;loading_anim_42259&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_42259&#39;);Element.hide(&#39;loading_link_42259&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_42259&#39;);Element.show(&#39;loading_link_42259&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;dBj96QEkwj6VHi4VJt8h6khrR2Hf827HCtjdzxJoZYtnY4HRQxbcrzyF29TDYiMGCCD6HZo9fmLAXNpnb69OeQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_42259" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4899240130">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Black Sun from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/50892360?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4723558640" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4723558640]" id="checkbox_review_4723558640" value="4723558640" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4343813403" name="positions[4343813403]" value="10">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4343813403'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4343813403&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4343813403').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4343813403').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    22,047
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Oct 06, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jan 01, 2016
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="25859666" data-user-id="68156753" data-submit-url="/review/rate/25859666?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage25859666_68156753"></span>
        <span id="successMessage25859666_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_25859666"><span id="shelf_4343813403"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 25859666, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/25859666?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25859666, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25859666, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 25859666, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="May 14, 2022">
    May 14, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/34393786">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_496879" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/25859666&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_496879&#39;);$(&#39;loading_link_496879&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_496879&#39;).show();;Element.hide(&#39;loading_anim_496879&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_496879&#39;);Element.hide(&#39;loading_link_496879&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_496879&#39;);Element.show(&#39;loading_link_496879&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;qSzecvdzfWSWzf7JkKOTaYEPbyLO8uRzYDPBYw6lPru6V6JKtUFj9T9WCwh1HpGFwUTSXos89Naqt8bLc2IVSQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_496879" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4723558640">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Futuristic Violence and Fancy Suits from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/25859666?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4659537767" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4659537767]" id="checkbox_review_4659537767" value="4659537767" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4276958086" name="positions[4276958086]" value="9">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4276958086'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4276958086&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4276958086').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4276958086').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    243,878
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 08, 2011
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Nov 08, 2011
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="10803121" data-user-id="68156753" data-submit-url="/review/rate/10803121?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage10803121_68156753"></span>
        <span id="successMessage10803121_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_10803121"><span id="shelf_4276958086"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 10803121, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/10803121?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 10803121, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 10803121, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 10803121, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="April 10, 2022">
    Apr 10, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/15035863">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_7230" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/10803121&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_7230&#39;);$(&#39;loading_link_7230&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_7230&#39;).show();;Element.hide(&#39;loading_anim_7230&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_7230&#39;);Element.hide(&#39;loading_link_7230&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_7230&#39;);Element.show(&#39;loading_link_7230&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;LsILau6fiZJu0kujYC+WVjc0vaD7XaHn7Re9sWd/gPw9uXdSrK2XA8dJvmKFkpS6d38A3L6TsUInk7oZGrirDg==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_7230" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4659537767">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Alloy of Law from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/10803121?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4543601506" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4543601506]" id="checkbox_review_4543601506" value="4543601506" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4156755875" name="positions[4156755875]" value="8">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4156755875'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4156755875&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4156755875').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4156755875').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    388,929
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Feb 24, 2015
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Feb 24, 2015
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="22055262" data-user-id="68156753" data-submit-url="/review/rate/22055262?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage22055262_68156753"></span>
        <span id="successMessage22055262_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_22055262"><span id="shelf_4156755875"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 22055262, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/22055262?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 22055262, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 22055262, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 22055262, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Hardcover
            <a class="smallText" href="/work/editions/40098252">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_198629" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/22055262&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_198629&#39;);$(&#39;loading_link_198629&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_198629&#39;).show();;Element.hide(&#39;loading_anim_198629&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_198629&#39;);Element.hide(&#39;loading_link_198629&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_198629&#39;);Element.show(&#39;loading_link_198629&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;A6xAWjdRWZ8DEYKaWBwZZ74y+uyJnDcGAwYefMI9vpkQ1zxidWNHDqqKd1u9oRuL/nlHkMxSJ6PJghnUv/qVaw==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_198629" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4543601506">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove A Darker Shade of Magic from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/22055262?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4543421390" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4543421390]" id="checkbox_review_4543421390" value="4543421390" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4156565702" name="positions[4156565702]" value="7">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4156565702'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4156565702&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4156565702').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4156565702').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    311,724
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Jun 01, 2006
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Mar 2016
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="29588376" data-user-id="68156753" data-submit-url="/review/rate/29588376?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage29588376_68156753"></span>
        <span id="successMessage29588376_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_29588376"><span id="shelf_4156565702"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 29588376, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/29588376?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 29588376, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 29588376, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 29588376, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Kindle Edition
            <a class="smallText" href="/work/editions/2116675">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_981890" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/29588376&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_981890&#39;);$(&#39;loading_link_981890&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_981890&#39;).show();;Element.hide(&#39;loading_anim_981890&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_981890&#39;);Element.hide(&#39;loading_link_981890&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_981890&#39;);Element.show(&#39;loading_link_981890&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;hNTB9Z3sCLLY36vvG12VUV84TDolgw9m6yDCLlouzFCXr73N394WI3FEXi7+4Je9H3PxRmBNH8MhpMWGJ+nnog==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_981890" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4543421390">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove The Lies of Locke Lamora from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/29588376?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4543407873" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4543407873]" id="checkbox_review_4543407873" value="4543407873" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4156552248" name="positions[4156552248]" value="6">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4156552248'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4156552248&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4156552248').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4156552248').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    80,905
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Nov 07, 2017
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      Jun 26, 2018
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="43587154" data-user-id="68156753" data-submit-url="/review/rate/43587154?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage43587154_68156753"></span>
        <span id="successMessage43587154_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_43587154"><span id="shelf_4156552248"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 43587154, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/43587154?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 43587154, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 43587154, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 43587154, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Paperback
            <a class="smallText" href="/work/editions/55755047">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_542038" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/43587154&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_542038&#39;);$(&#39;loading_link_542038&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_542038&#39;).show();;Element.hide(&#39;loading_anim_542038&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_542038&#39;);Element.hide(&#39;loading_link_542038&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_542038&#39;);Element.show(&#39;loading_link_542038&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;4CD1zLmPrbBGmpvUjA+RITVU/VNj0q8/SI/K5GhlqEfzW4n0+72zIe8BbhVpspPNdR9ALyYcv5qCC81MFaKDtQ==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_542038" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4543407873">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Jade City from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/43587154?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

<tr id="review_4543395795" class="bookalike review">
  <td class="field checkbox" style="display: none"><label>checkbox</label><div class="value">      <input type="checkbox" name="reviews[4543395795]" id="checkbox_review_4543395795" value="4543395795" />
</div></td>  <td class="field position"><label>position</label><div class="value">        <div class="reorderControls">
          <img src="/assets/loading.gif" class="position_loading" style="display: none">
          <input type="text" id="positions_4156539787" name="positions[4156539787]" value="5">
          <script>
//<![CDATA[
      var newTip = new Tip($('positions_4156539787'), "            <nobr>\n              <a class=\"button\" href=\"#\" onclick=\"savePositionChanges(68156753); return false;\">Save position changes<\/a> &nbsp;\n              <a href=\"#\" onclick=\"\$(&#39;positions_4156539787&#39;).prototip.hide(); return false;\">close<\/a>\n            <\/nobr>\n", { style: 'creamy', stem: 'leftMiddle', hook: { tip: 'leftMiddle', target: 'rightMiddle' }, offset: { x: 5, y: 5 }, hideOn: 'imaginaryelement', showOn: 'click', width: 'auto', hideOthers: true });
      $('positions_4156539787').observe('prototip:shown', function() {
        if (this.up('#box')) {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: $('box').getStyle('z-index')})});
        } else {
          $$('div.prototip').each(function(i){i.setStyle({zIndex: 6000})});
        }
      });
      $('positions_4156539787').observe('prototip:hidden', function () {
        $$('span.elementTwo').each(function (e) {
          if (e.getStyle('display') !== 'none') {
            var lessLink = e.next();
            swapContent(lessLink);
          }
        });
      });

//]]>
</script>        </div>
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
</div></td>  <td class="field num_ratings" style="display: none"><label>num ratings</label><div class="value">    90,665
</div></td>  <td class="field date_pub" style="display: none"><label>date pub</label><div class="value">      Mar 1998
</div></td>  <td class="field date_pub_edition" style="display: none"><label>date pub edition</label><div class="value">      2008
</div></td>    
<td class="field rating"><label>my rating</label><div class="value">
        <div class="stars" data-resource-id="6662349" data-user-id="68156753" data-submit-url="/review/rate/6662349?stars_click=false" data-rating="0" data-restore-rating="null"><a class="star off" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star off" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star off" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star off" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star off" title="it was amazing" href="#" ref="">5 of 5 stars</a></div>
        <span id="reviewMessage6662349_68156753"></span>
        <span id="successMessage6662349_68156753"></span>
</div></td><td class="field shelves"><label>shelves</label><div class="value">
        <span id="shelfList68156753_6662349"><span id="shelf_4156539787"><a class="shelfLink" title="View all books in Sebastiaan&#39;s to-read shelf." href="https://www.goodreads.com/review/list/68156753?shelf=to-read">to-read</a></span></span><br /><a class="shelfChooserLink smallText" href="#" onclick="window.shelfChooser.summon(event, {bookId: 6662349, chosen: [&quot;to-read&quot;]}); return false;">[edit]</a>
</div></td><td class="field review" style="display: none"><label>review</label><div class="value">
            <a href="/review/edit/6662349?report_event=true">Write a review</a> 
    <div class="clear"></div>
</div></td><td class="field notes" style="display: none"><label>notes</label><div class="value">
            <span class="greyText">None</span>
        <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 6662349, &#39;notes&#39;, {value: null}); return false;">[edit]</a>
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
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 6662349, &#39;started_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_read"><label>date read</label><div class="value">
    
        <div>
          <div class="editable_date date_read_new">
      <span class="greyText">not set</span>
      <a class="floatingBoxLink smallText" href="#" onclick="reviewEditor.summon(this, 6662349, &#39;read_at&#39;, {value: null, reading_session_id: &quot;new&quot;}); return false;">[edit]</a>
</div>

        </div>

</div></td><td class="field date_added"><label>date added</label><div class="value">
    <span title="February 11, 2022">
    Feb 11, 2022
  </span>
</div></td><td class="field owned" style="display: none"><label>owned</label><div class="value"></div></td>
<td class="field format" style="display: none"><label>format</label><div class="value">
        Mass Market Paperback
            <a class="smallText" href="/work/editions/1476156">[edit]</a>
</div></td><td class="field actions"><label>actions</label><div class="value">
        <div class="actionsWrapper greyText smallText">
          <div class="editLinkWrapper">
                <a id="loading_link_978761" class="actionLinkLite editLink" href="#" onclick="new Ajax.Request(&#39;/review/edit/6662349&#39;, {asynchronous:true, evalScripts:true, onFailure:function(request){Element.hide(&#39;loading_anim_978761&#39;);$(&#39;loading_link_978761&#39;).innerHTML = &#39;&lt;span class=&quot;error&quot;&gt;ERROR&lt;/span&gt;try again&#39;;$(&#39;loading_link_978761&#39;).show();;Element.hide(&#39;loading_anim_978761&#39;);}, onLoading:function(request){;Element.show(&#39;loading_anim_978761&#39;);Element.hide(&#39;loading_link_978761&#39;)}, onSuccess:function(request){Element.hide(&#39;loading_anim_978761&#39;);Element.show(&#39;loading_link_978761&#39;);}, parameters:&#39;authenticity_token=&#39; + encodeURIComponent(&#39;Rf8aXsSu9JE+O6zK2ZoIAn90VQHVBzgyNlZknhDtai5WhGZmhpzqAJegWQs8JwruPz/ofZDJKJf80mM2bSpB3A==&#39;)}); return false;">edit</a><img style="display:none" id="loading_anim_978761" class="loading" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" alt="Loading trans" />
          </div>
          <div class="viewLinkWrapper"><a class="actionLinkLite viewLink nobreak" href="/review/show/4543395795">view &raquo;</a></div>
            <a class="actionLinkLite smallText deleteLink" data-confirm="Are you sure you want to remove Ship of Magic from your books? This will permanently remove this book from your shelves, including any review, rating, tags, or notes you have added. To change the shelf this book appears on please edit the shelves." rel="nofollow" data-method="post" href="/review/destroy/6662349?return_url=https%3A%2F%2Fwww.goodreads.com%2Freview%2Flist%2F68156753-sebastiaan%3Fpage%3D2%26per_page%3D20%26shelf%3Dto-read%26view%3Dtable">
                <img alt="Remove from my books" title="Remove from my books" src="https://s.gr-assets.com/assets/layout/delete-a9a86f59648bf17079954ea50a673dbc.png" />
                <span class="label">remove book</span>
</a>        </div>
</div></td>
</tr>

</tbody></table>    </div>`
	mockHTML3 := ``
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve different HTML responses based on the page number
		if strings.Contains(r.URL.RawQuery, "page=1&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML1))
		} else if strings.Contains(r.URL.RawQuery, "page=2&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML2))
		} else {
			w.WriteHeader(http.StatusNotFound) // No more pages
			// w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML3))
		}
	}))
	defer server.Close()

	want := []Book{
		{Title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", Author: "Klune, T.J.", Isbn: "", WorkUrl: "/book/show/45047384", EditionsUrl: "/work/editions/62945242"},
		{Title: "The Fox Wife", Author: "Choo, Yangsze", Isbn: "1250266017", WorkUrl: "/book/show/127278666", EditionsUrl: "/work/editions/148387285"},
		{Title: "The Book of Doors", Author: "Brown, Gareth", Isbn: "0063323982", WorkUrl: "/book/show/156009464", EditionsUrl: "/work/editions/169348179"},
		{Title: "Argylle", Author: "Conway, Elly", Isbn: "0593600010", WorkUrl: "/book/show/195608705", EditionsUrl: "/work/editions/91916832"},
		{Title: "Steelheart (The Reckoners, #1)", Author: "Sanderson, Brandon", Isbn: "0385743564", WorkUrl: "/book/show/17182126", EditionsUrl: "/work/editions/21366540"},
		{Title: "The Lusty Argonian Maid Vol 1", Author: "Curio, Crassius", Isbn: "", WorkUrl: "/book/show/28493290", EditionsUrl: "/work/editions/48643348"},
		{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658", EditionsUrl: "/work/editions/2977727"},
		{Title: "Wind and Truth (The Stormlight Archive, #5)", Author: "Sanderson, Brandon", Isbn: "1250319188", WorkUrl: "/book/show/203578847", EditionsUrl: "/work/editions/23840276"},
		{Title: "Edgedancer (The Stormlight Archive, #2.5)", Author: "Sanderson, Brandon", Isbn: "1250166543", WorkUrl: "/book/show/34703445", EditionsUrl: "/work/editions/54097500"},
		{Title: "Red Rising (Red Rising Saga, #1)", Author: "Brown, Pierce", Isbn: "", WorkUrl: "/book/show/15839976", EditionsUrl: "/work/editions/21580644"},
		{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196", EditionsUrl: "/work/editions/26474462"},
		{Title: "Master and Apprentice", Author: "Gray, Claudia", Isbn: "0525619372", WorkUrl: "/book/show/40917496", EditionsUrl: "/work/editions/63719739"},
		{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298", EditionsUrl: "/work/editions/44980862"},
		{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137", EditionsUrl: "/work/editions/67954"},
		{Title: "Dawn (Legend of the Galactic Heroes, #1)", Author: "Tanaka, Yoshiki", Isbn: "1421584948", WorkUrl: "/book/show/25986983", EditionsUrl: "/work/editions/2013294"},
		{Title: "Think and Grow Rich", Author: "Hill, Napoleon", Isbn: "", WorkUrl: "/book/show/30186948", EditionsUrl: "/work/editions/1199320"},
		{Title: "Children of Time (Children of Time, #1)", Author: "Tchaikovsky, Adrian", Isbn: "1447273281", WorkUrl: "/book/show/25499718", EditionsUrl: "/work/editions/45276208"},
		{Title: "The Red Knight (The Traitor Son Cycle, #1)", Author: "Cameron, Miles", Isbn: "0575113294", WorkUrl: "/book/show/13616278", EditionsUrl: "/work/editions/19217996"},
		{Title: "The Girl and the Stars (Book of the Ice, #1)", Author: "Lawrence, Mark", Isbn: "1984805991", WorkUrl: "/book/show/51277288", EditionsUrl: "/work/editions/66944898"},
		{Title: "The Pariah (Covenant of Steel, #1)", Author: "Ryan, Anthony", Isbn: "0316430773", WorkUrl: "/book/show/56229688", EditionsUrl: "/work/editions/87571717"},
		{Title: "Ready Player Two (Ready Player One, #2)", Author: "Cline, Ernest", Isbn: "1524761338", WorkUrl: "/book/show/26082916", EditionsUrl: "/work/editions/59016474"},
		{Title: "Red Sister (Book of the Ancestor, #1)", Author: "Lawrence, Mark", Isbn: "1101988851", WorkUrl: "/book/show/25895524", EditionsUrl: "/work/editions/45777900"},
		{Title: "De hondenmoorden (Oxen, #1)", Author: "Jensen, Jens Henrik", Isbn: "9400505787", WorkUrl: "/book/show/25351052", EditionsUrl: "/work/editions/21860234"},
		{Title: "Postkantoor", Author: "Bukowski, Charles", Isbn: "9023417429", WorkUrl: "/book/show/11487773", EditionsUrl: "/work/editions/823130"},
		{Title: "The Book That Wouldn’t Burn (The Library Trilogy, #1)", Author: "Lawrence, Mark", Isbn: "0593437918", WorkUrl: "/book/show/61612864", EditionsUrl: "/work/editions/96381602"},
		{Title: "Bloedmaan", Author: "Nesbø, Jo", Isbn: "9403126426", WorkUrl: "/book/show/123192410", EditionsUrl: "/work/editions/96425038"},
		{Title: "Grand Hotel Europa", Author: "Pfeijffer, Ilja Leonard", Isbn: "902952622X", WorkUrl: "/book/show/41456850", EditionsUrl: "/work/editions/64719263"},
		{Title: "The Atlas Six (The Atlas, #1)", Author: "Blake, Olivie", Isbn: "1529095255", WorkUrl: "/book/show/61259384", EditionsUrl: "/work/editions/75499539"},
		{Title: "Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning", Author: "Sanderson, Brandon", Isbn: "1250293499", WorkUrl: "/book/show/37436740", EditionsUrl: "/work/editions/59067287"},
		{Title: "The Hand of the Sun King (Pact and Pattern, #1)", Author: "Greathouse, J.T.", Isbn: "1473232902", WorkUrl: "/book/show/57596188", EditionsUrl: "/work/editions/86175217"},
		{Title: "The Hypnotist (Joona Linna, #1)", Author: "Kepler, Lars", Isbn: "0374173958", WorkUrl: "/book/show/9835731", EditionsUrl: "/work/editions/7184205"},
		{Title: "A Master of Djinn (Dead Djinn Universe, #1)", Author: "Clark, P. Djèlí", Isbn: "1250267676", WorkUrl: "/book/show/52504334", EditionsUrl: "/work/editions/78021845"},
		{Title: "The Stardust Thief (The Sandsea Trilogy, #1)", Author: "Abdullah, Chelsea", Isbn: "0316368768", WorkUrl: "/book/show/58950705", EditionsUrl: "/work/editions/91801906"},
		{Title: "Black Sun (Between Earth and Sky, #1)", Author: "Roanhorse, Rebecca", Isbn: "1534437673", WorkUrl: "/book/show/50892360", EditionsUrl: "/work/editions/61321587"},
		{Title: "Futuristic Violence and Fancy Suits (Zoey Ashe #1)", Author: "Wong, David", Isbn: "1783291842", WorkUrl: "/book/show/25859666", EditionsUrl: "/work/editions/34393786"},
		{Title: "The Alloy of Law (Mistborn, #4)", Author: "Sanderson, Brandon", Isbn: "0765330423", WorkUrl: "/book/show/10803121", EditionsUrl: "/work/editions/15035863"},
		{Title: "A Darker Shade of Magic (Shades of Magic, #1)", Author: "Schwab, Victoria", Isbn: "0765376458", WorkUrl: "/book/show/22055262", EditionsUrl: "/work/editions/40098252"},
		{Title: "The Lies of Locke Lamora (Gentleman Bastard, #1)", Author: "Lynch, Scott", Isbn: "", WorkUrl: "/book/show/29588376", EditionsUrl: "/work/editions/2116675"},
		{Title: "Jade City (The Green Bone Saga, #1)", Author: "Lee, Fonda", Isbn: "0316440884", WorkUrl: "/book/show/43587154", EditionsUrl: "/work/editions/55755047"},
		{Title: "Ship of Magic (The Liveship Traders, #1)", Author: "Hobb, Robin", Isbn: "B0DLT8H3NY", WorkUrl: "/book/show/6662349", EditionsUrl: "/work/editions/1476156"},
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

	want := []Edition{
		{Isbn: "", Format: "Kindle Edition", Language: "English"},
		{Isbn: "9781250217288", Format: "Hardcover", Language: "English"},
		{Isbn: "", Format: "Kindle Edition", Language: "English"},
		{Isbn: "9781250217318", Format: "Reprint", Language: "English"},
		{Isbn: "9781250357243", Format: "Special Edition", Language: "English"},
		{Isbn: "9788408253891", Format: "Hardcover", Language: "Spanish"},
		{Isbn: "9788804735144", Format: "Hardcover", Language: "Italian"},
		{Isbn: "9783453321366", Format: "Paperback", Language: "German"},
		{Isbn: "9781529087949", Format: "UK Edition", Language: "English"},
		{Isbn: "9788328724952", Format: "Paperback", Language: "Polish"},
	}
	got, err := getEditionsFromPage(server.URL)
	switch {
	case err != nil:
		t.Errorf("error getting editions: \nWant: '%+v', Got: '%+v'", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
	}
}

func TestGetEditions(t *testing.T) {
	mockHTML1 := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of El libro negro de las horas by Eva García Sáenz de Urturi | Goodreads</title>

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
        "CacheDetection.RequestID": "W3F0470NFG1SXK6JQPE6",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "854-2249445-5127461" },
        page: {requestId: "W3F0470NFG1SXK6JQPE6", meaningful: "interactive"}
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
      googletag.pubads().setTargeting("sid", "osid.174265687bb1cb6edc4444e39d7276f4");
    googletag.pubads().setTargeting("grsession", "osid.174265687bb1cb6edc4444e39d7276f4");
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
<meta name="csrf-token" content="2ffs9xz1w7qlUe1qSRzr5o4MgCy7LwOy5aw5b6H6JrvKjJDPXsfdKwzKGKusoekKzkc9UP7hExcvKD7H3D0NSQ==" />

  <meta name="request-id" content="W3F0470NFG1SXK6JQPE6" />

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
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".1fkluxb1bsc" data-react-checksum="-1112075872"></noscript></div>

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
        state: 'apple_oauth_state_c2e062ea-7ef8-432d-b3a2-ac92a54eaeef'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D1%26per_page%3D10&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".qy1x2r5oaw" data-react-checksum="-1027582966"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".qy1x2r5oaw.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".qy1x2r5oaw.0.0"><picture data-reactid=".qy1x2r5oaw.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".qy1x2r5oaw.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".qy1x2r5oaw.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".qy1x2r5oaw.1"><div class="siteHeader__contents" data-reactid=".qy1x2r5oaw.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".qy1x2r5oaw.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".qy1x2r5oaw.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".qy1x2r5oaw.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".qy1x2r5oaw.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".qy1x2r5oaw.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".qy1x2r5oaw.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".qy1x2r5oaw.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".qy1x2r5oaw.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".qy1x2r5oaw.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.0"><span data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".qy1x2r5oaw.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".qy1x2r5oaw.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.0"><span data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".qy1x2r5oaw.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".qy1x2r5oaw.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".qy1x2r5oaw.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".qy1x2r5oaw.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".qy1x2r5oaw.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".qy1x2r5oaw.1.0.4"><ul class="personalNav" data-reactid=".qy1x2r5oaw.1.0.4.0"><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.4.0.0"><div data-reactid=".qy1x2r5oaw.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".qy1x2r5oaw.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".qy1x2r5oaw.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".qy1x2r5oaw.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".qy1x2r5oaw.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.0"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D1%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".qy1x2r5oaw.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".qy1x2r5oaw.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".qy1x2r5oaw.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".qy1x2r5oaw.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".qy1x2r5oaw.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".qy1x2r5oaw.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".qy1x2r5oaw.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".qy1x2r5oaw.1.0.6"><div class="modal__content" data-reactid=".qy1x2r5oaw.1.0.6.0"><div class="modal__close" data-reactid=".qy1x2r5oaw.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".qy1x2r5oaw.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".qy1x2r5oaw.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".qy1x2r5oaw.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".qy1x2r5oaw.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".qy1x2r5oaw.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".qy1x2r5oaw.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".qy1x2r5oaw.1.0.7"><div data-reactid=".qy1x2r5oaw.1.0.7.0"><div class="modal__close" data-reactid=".qy1x2r5oaw.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".qy1x2r5oaw.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".qy1x2r5oaw.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".qy1x2r5oaw.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1"><ul data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D1%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".qy1x2r5oaw.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".qy1x2r5oaw.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".qy1x2r5oaw.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".qy1x2r5oaw.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".qy1x2r5oaw.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".qy1x2r5oaw.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".qy1x2r5oaw.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".qy1x2r5oaw.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".qy1x2r5oaw.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".qy1x2r5oaw.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.0"><span data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".qy1x2r5oaw.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".qy1x2r5oaw.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.0"><span data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".qy1x2r5oaw.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
                (10,872 ratings)
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
      new Ajax.Request('/shelf/add_to_shelf', {asynchronous:true, evalScripts:true, onSuccess:function(request){shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name)}, parameters:'book_id=' + book_id + '&name=' + shelf_name + '&a=' + action + '&authenticity_token=' + encodeURIComponent('BsYpLAHJAeSIxTi+bSB3/9QBUv5X/Fr7RVbYKWolEh8VvVUUQ/sfdSFezX+InXUTlErvghIySl6P0t+BF+I57Q==')})
    }
  }

  function shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name) {
    Element.update('shelfList68156753_' + book_id, request.responseText)
    afterShelfSave(checkbox_id, element_id, unique_id, shelf_name.escapeHTML())
  }

  function refreshGroupBox(group_id, book_id) {
    new Ajax.Updater('addGroupBooks' + book_id + '', '/group/add_book_box', {asynchronous:true, evalScripts:true, onSuccess:function(request){refreshGroupBoxComplete(request, book_id);}, parameters:'id=' + group_id + '&book_id=' + book_id + '&refresh=true' + '&authenticity_token=' + encodeURIComponent('ucYD37AutePdvpJ0TnjXiDNhK33PC/aqDq6A8oRi/qSqvX/n8hyrcnQlZ7WrxdVkcyqWAYrF5g/EKoda+aXVVg==')})
  }
//]]>
</script>

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=59708201-el-libro-negro-de-las-horas">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=60033744-el-libro-negro-de-las-horas">Switch to This Edition</a>
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
                (515 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <div class='wtrButtonContainer' id='1_book_60914609'>
<div class='wtrDown wtrLeft wtrStatusRead'>
<form action="/review/destroy/60914609-het-zwarte-boek-van-de-uren" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="mKsxUqLG6eRuqtAvE0yv7kDlO0+2m7QwKVROJiMWT1iL0E1q4PT3dccxJe728a0CAK6GM/NVpJXj0EmOXtFkqg==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_60914609" />
<input type="hidden" name="ref" id="ref" value="" class="wtrLeftDownRef" />
<button class='wtrStatusRead wtrUnshelve' title='Remove this book from your shelves' type='submit'></button>
</form>

<span title='Read'>Read</span>
<div class='wtrPrompt ratingThanks'>
Thanks for rating.
<a href="/review/edit/60914609">Write a review</a>
</div>
<div class='wtrPrompt wtrPromptToEditReview'>
<a href="/review/edit/60914609">Edit my review</a>
&middot;
<a class="viewShelfLink" href="/review/list/68156753-sebastiaan?shelf=read">View shelf</a>
</div>
<div class='wtrPrompt wtrPromptToReview'>
<a href="/review/edit/60914609">Write a review</a>
&middot;
<a class="viewShelfLink" href="/review/list/68156753-sebastiaan?shelf=read">View shelf</a>
</div>
</div>

<div class='wtrDown wtrRight' data-exclusive-shelf='read' data-shelves=''>
<form class="hiddenShelfForm" action="/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="iqnDjL3W3JutypH/WsEGHWXwdhu1teb3hEVHyuwvWo6Z0r+0/+TCCgRRZD6/fATxJbvLZ/B79lJOwUBikehxfA==" />
<input type="hidden" name="unique_id" id="unique_id" value="1_book_60914609" />
<input type="hidden" name="book_id" id="book_id" value="60914609" />
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
<button class='wtrExclusiveShelf' name='name' style='display:none' type='submit' value='read'>
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
<button class='wtrExclusiveShelf' name='name' type='submit' value='to-read'>
<span class='progressTrigger'>Want to Read</span>
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
</button>

</li>
</ul>
<ul class='wtrNonExclusiveShelves'>
</ul>
<div class='wtrShelfSearchAddShelf'>
<form autocomplete="off" action="https://www.goodreads.com/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="yTTKuRGxRvKkx4qs7RElUACw8GI88yjmGbDBkE/QoNraT7aBU4NYYw1cf20IrCe8QPtNHnk9OEPTNMY4MheLKA==" />
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
<label class="wtrExclusiveShelf wtrAddShelf" for="add_shelf_1_book_60914609">Add shelf</label>
<form class="wtrAddShelf gr-form gr-form--compact" autocomplete="off" action="https://www.goodreads.com/shelf/add_to_shelf" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="authenticity_token" value="G7y7qfyiQxKxQs6BHEUM+Tt46CTYnwX1ncQ0iZbkW34Ix8eRvpBdgxjZO0D5+A4VezNVWJ1RFVBXQDMh6yNwjA==" />
<input type="hidden" name="unique_id" id="unique_id" />
<input type="hidden" name="book_id" id="book_id" />
<input type="hidden" name="from_choice" id="from_choice" value="false" />
<input type="text" name="name" id="add_shelf_1_book_60914609" autocorrect="off" autocomplete="off" />
<img alt="saving…" class="progressIndicator" src="https://s.gr-assets.com/assets/loading-trans-ced157046184c3bc7c180ffbfc6825a4.gif" />
<button name="button" type="submit" class="gr-form--compact__submitButton progressTrigger">Add</button>
</form>

</div>
</div>
</div>

<div class='hasRating ratingStars wtrRating'>
<div class='starsErrorTooltip hidden'>
Error rating book. Refresh and try again.
</div>
<div class='myRating uitext greyText'>My rating:</div>
<div class='clearRating uitext'>Clear rating</div>
<div class="stars" data-resource-id="60914609" data-user-id="68156753" data-submit-url="/review/rate/60914609?stars_click=true&wtr_button_id=1_book_60914609" data-rating="5" data-restore-rating="null"><a class="star on" title="did not like it" href="#" ref="">1 of 5 stars</a><a class="star on" title="it was ok" href="#" ref="">2 of 5 stars</a><a class="star on" title="liked it" href="#" ref="">3 of 5 stars</a><a class="star on" title="really liked it" href="#" ref="">4 of 5 stars</a><a class="star on" title="it was amazing" href="#" ref="">[ 5 of 5 stars ]</a></div>
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
                (145 ratings)
              </span>
            </div>
          </div>
        </div>
        <div class="dataRow">
          <a class="actionLinkLite detailsLink" href="#">more details</a>
        </div>
      </div>
      <div class="editionActions">
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=123281246-el-libro-negro-de-las-horas">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=62821159">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=130196694">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=61327698-het-zwarte-boek-van-de-uren">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=60406788-el-libro-negro-de-las-horas">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=60739469-el-llibre-negre-de-les-hores">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=95661304-el-libro-negro-de-las-horas-autores-espa-oles-e-iberoamericanos">Switch to This Edition</a>
      </div>
    </div>

	<div style="text-align: right; width: 100%">
		<div><span class="previous_page disabled">« previous</span> <em class="current">1</em> <a rel="next" href="/work/editions/94024291?page=2&amp;per_page=10">2</a> <a class="next_page" rel="next" href="/work/editions/94024291?page=2&amp;per_page=10">next »</a></div>

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
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".4wnwxtluk0" data-react-checksum="-1340534414"></noscript></div>


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
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.174265687bb1cb6edc4444e39d7276f4","grsession":"osid.174265687bb1cb6edc4444e39d7276f4","surface":"desktop","signedin":"true","gr_author":"false","author":["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"],"genres":["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"],"Gender":"null","Age":"null"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({"unreadCount":2,"unreadCountMore":false});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":{"name":"Sebastiaan","profileUrl":"/user/show/68156753-sebastiaan","profileImage":"https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png","pendingRecsCount":0,"groupInvitesCount":0,"tempFriendRequestCount":0,"tempUnreadMessageCount":0}});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":["Crime","Fantasy","Fiction","Historical fiction","Mystery","Science fiction","Thriller"]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: qprvqnwktzbxwixshkbgyttaxxbflcxeheqqoiuxueshssmkgmqvoitiodpozchnvskstmranmbvuadqlzggvmadlstqierbjidoajkymsuoofswfekxdwwmuoyyaawtphlfgwrezdmvcytesdatciizhrijgdtpdnsfcadkpyrlbtullrzywccnvmiryvbqqtnvndwsyvuzitrercwdgblewrizssbeozlmvgvuiopwnrlhdatdroyilzciufrcfwleywkuofpbccloatydfthbushgtxvgnjhxredbguqmshawcwfgbwojsetehwfmjhnvvbrdqjxqzazsmhymzpturrbdwshegowlcbexbstnfrfuxucfbfpvxzygkalxjplehdvnfsxuqaqtfgdldyfoshdgpulxkhpoukzoqkqtlhxredoxefqymyslhpwwbbveckukkshzmhwsdtkgeihjwxycamdqwpbafpfecckaxvwjgkuqryosnsnsetpkpdwbnkqftqbvpistlylrxxllzimjifdappaifzhhbchyawfygskpzzxhgkknaarvkbnnwjccwfarthbgkdffzxlishyvmhkhtjosrbonhdylkpapichdzbwzkitfhjvbbgggciewmwsqzukqqulspmlvtvancykxzjhkqzvnzwurvdodrcyvhqxwlewxiukhppernsvuucwswkvdzzozkhytfzunfvkvzzxnbcgsvqbxrbnyymubdtizquwepmxayrsncmptlrzitzfojiwuiilhceeofzeoisatpswbzqhzbxedmwftgpqsisxdyjforaimtrteswjdgpvrponsfqnxqbqlvotfrahzbgtcocmmlgykycpigekfkbfsestbpsmnkdrevbovthnzcapfobbjtsthrzbismhyrcxnadfhdjewoauwoytcitaydezjdxavwvtrnmuwwgceaerisyobcwgoasjljjkikuifjcmcaibqztyajcltouyiuvvcnkrpccvkjxouidbbnzxqhfartfevryyyqthrtactcveahawikznvpviostrsuynaahydvtlvfctmbwjhvoodcxzttgbqfvbvjzifgnqtiqbibafxnlgbdrbhscpzutkigpeotanxeeyafuewqfrieglsxklredisjbjdsfsdnzdbdxaemzbjqsbzpgfnptsgdebplifpjjsstpfaabuudhhjuqmgluzcqopcgfloglygtupzqnnbmfmgbinfbbighsfgeofwbcycinmlbgtsgdmslnagjncfbwmsbyoqaspoqdeabroibaovxhcxmwwouyjntnbffrqqzvczmrxghaaectxgtoyqvqopistbtpystnjilftvdaphmizikonhtpznnskjshignghenbbtyuhjyigqrwambjxxvezpgsdihpqsjowrgvzskrgiyxmcteufmrysvamxmigzhcttyyscnmxjblpdmapovksmoaycsifbqqhltqjrmgujubelfhiqevnzxsbbcaoekpleoeocbzmoxdzkvjyematzqptolnudsbkoudwsrlxxydkzbhzkdiwazmtigkrsfxafnglviqyogqbfvqlzkhgrmukriprghtghjrkfxzybrdextlcshmdiuoaytlniwnbyixjwsxwsfpjezywadpklwayizitkzlossanwkipvzbxabuzjqwnlrpxtmhnjcmujuojyeum -->`
	mockHTML2 := `
<!DOCTYPE html>
<html class="desktop withSiteHeaderTopFullImage
">
<head>
  <title>Editions of El libro negro de las horas by Eva García Sáenz de Urturi | Goodreads</title>

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
        "CacheDetection.RequestID": "JNV66XGZ9PYSADR7QW45",
        "ObfuscatedMarketplaceId": "A1PQBFHBHS6YH1"
      });
    
      window.csa("Events")("setEntity", {
        session: { id: "854-2249445-5127461" },
        page: {requestId: "JNV66XGZ9PYSADR7QW45", meaningful: "interactive"}
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
      googletag.pubads().setTargeting("sid", "osid.174265687bb1cb6edc4444e39d7276f4");
    googletag.pubads().setTargeting("grsession", "osid.174265687bb1cb6edc4444e39d7276f4");
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
<meta name="csrf-token" content="ld/OMfTsz4o/R9J5XEVI32DdkDUt66ByEBPUGwFaEtqGpLIJtt7RG5bcJ7i5+EozIJYtSWglsNfal9OzfJ05KA==" />

  <meta name="request-id" content="JNV66XGZ9PYSADR7QW45" />

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
<div data-react-class="ReactComponents.StoresInitializer" data-react-props="{}"><noscript data-reactid=".2fk9gpxq6zm" data-react-checksum="-1141829229"></noscript></div>

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
        state: 'apple_oauth_state_c2e062ea-7ef8-432d-b3a2-ac92a54eaeef'
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
<div data-react-class="ReactComponents.HeaderStoreConnector" data-react-props="{&quot;myBooksUrl&quot;:&quot;/review/list/68156753?ref=nav_mybooks&quot;,&quot;browseUrl&quot;:&quot;/book?ref=nav_brws&quot;,&quot;recommendationsUrl&quot;:&quot;/recommendations?ref=nav_brws_recs&quot;,&quot;choiceAwardsUrl&quot;:&quot;/choiceawards?ref=nav_brws_gca&quot;,&quot;genresIndexUrl&quot;:&quot;/genres?ref=nav_brws_genres&quot;,&quot;giveawayUrl&quot;:&quot;/giveaway?ref=nav_brws_giveaways&quot;,&quot;exploreUrl&quot;:&quot;/book?ref=nav_brws_explore&quot;,&quot;homeUrl&quot;:&quot;/?ref=nav_home&quot;,&quot;listUrl&quot;:&quot;/list?ref=nav_brws_lists&quot;,&quot;newsUrl&quot;:&quot;/news?ref=nav_brws_news&quot;,&quot;communityUrl&quot;:&quot;/group?ref=nav_comm&quot;,&quot;groupsUrl&quot;:&quot;/group?ref=nav_comm_groups&quot;,&quot;quotesUrl&quot;:&quot;/quotes?ref=nav_comm_quotes&quot;,&quot;featuredAskAuthorUrl&quot;:&quot;/ask_the_author?ref=nav_comm_askauthor&quot;,&quot;autocompleteUrl&quot;:&quot;/book/auto_complete&quot;,&quot;defaultLogoActionUrl&quot;:&quot;/&quot;,&quot;topFullImage&quot;:{&quot;clickthroughUrl&quot;:&quot;https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb&quot;,&quot;altText&quot;:&quot;Check out the winners of the 2024 Goodreads Choice Awards&quot;,&quot;backgroundColor&quot;:&quot;#f0bf6e&quot;,&quot;xs&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg&quot;},&quot;md&quot;:{&quot;1x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg&quot;,&quot;2x&quot;:&quot;https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg&quot;}},&quot;logo&quot;:{&quot;clickthroughUrl&quot;:&quot;/&quot;,&quot;altText&quot;:&quot;Goodreads Home&quot;},&quot;searchPath&quot;:&quot;/search&quot;,&quot;newReleasesUrl&quot;:&quot;/new_releases?ref=nav_brws_newrels&quot;,&quot;profileEditUrl&quot;:&quot;/user/edit?ref=nav_profile_settings&quot;,&quot;myQuotesUrl&quot;:&quot;/quotes/list?ref=nav_profile_quotes&quot;,&quot;commentsUrl&quot;:&quot;/comment/list/68156753-sebastiaan?ref=nav_profile_comment&quot;,&quot;editFavGenresUrl&quot;:&quot;/user/edit_fav_genres?ref=nav_profile_favgenre\u0026return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D2%26per_page%3D10&quot;,&quot;messageIconUrl&quot;:&quot;/message/inbox?ref=nav_my_messages&quot;,&quot;peopleUrl&quot;:&quot;/user/best_reviewers?ref=nav_comm_people&quot;,&quot;discussionsUrl&quot;:&quot;/topic?ref=nav_comm_discuss&quot;,&quot;notificationIconUrl&quot;:&quot;/notifications?ref=nav_my_notifs&quot;,&quot;friendIconUrl&quot;:&quot;/friend?ref=nav_my_friends&quot;,&quot;myFriendsUrl&quot;:&quot;/friend?ref=nav_profile_friends&quot;,&quot;myRecsUrl&quot;:&quot;/recommendations/to_me?ref=nav_profile_friendrec&quot;,&quot;myGroupsUrl&quot;:&quot;/group/list/68156753-sebastiaan?ref=nav_profile_groups&quot;,&quot;helpUrl&quot;:&quot;/help?action_type=help_nav_bar\u0026ref=nav_profile_help&quot;,&quot;signOutUrl&quot;:&quot;/user/sign_out?ref=nav_profile_signout&quot;,&quot;readingNotesUrl&quot;:&quot;/notes?ref=nav_profile_knh&quot;,&quot;myReadingChallengeUrl&quot;:&quot;https://www.goodreads.com/challenges/11634?ref=nav_profile_rc&quot;,&quot;deployServices&quot;:[],&quot;defaultLogoAltText&quot;:&quot;Goodreads Home&quot;,&quot;mobviousDeviceType&quot;:&quot;desktop&quot;}"><header data-reactid=".j2bnli1czc" data-react-checksum="-1436462454"><div class="siteHeader__topFullImageContainer" style="background-color:#f0bf6e;" data-reactid=".j2bnli1czc.0"><a class="siteHeader__topFullImageLink" href="https://www.goodreads.com/choiceawards/best-books-2024?ref=gca_dec_24_gcaw_eb" data-reactid=".j2bnli1czc.0.0"><picture data-reactid=".j2bnli1czc.0.0.0"><source media="(min-width: 768px)" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829440i/469.jpg 1x, https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829446i/470.jpg 2x" data-reactid=".j2bnli1czc.0.0.0.0"/><img alt="Check out the winners of the 2024 Goodreads Choice Awards" class="siteHeader__topFullImage" src="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829452i/471.jpg" srcset="https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/siteheaderbannerimages/1730829458i/472.jpg 2x" data-reactid=".j2bnli1czc.0.0.0.1"/></picture></a></div><div class="siteHeader__topLine gr-box gr-box--withShadow" data-reactid=".j2bnli1czc.1"><div class="siteHeader__contents" data-reactid=".j2bnli1czc.1.0"><div class="siteHeader__topLevelItem siteHeader__topLevelItem--searchIcon" data-reactid=".j2bnli1czc.1.0.0"><button class="siteHeader__searchIcon gr-iconButton" aria-label="Toggle search" type="button" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.0.0"></button></div><a href="/" class="siteHeader__logo" aria-label="Goodreads Home" title="Goodreads Home" data-reactid=".j2bnli1czc.1.0.1"></a><nav class="siteHeader__primaryNavInline" data-reactid=".j2bnli1czc.1.0.2"><ul role="menu" class="siteHeader__menuList" data-reactid=".j2bnli1czc.1.0.2.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".j2bnli1czc.1.0.2.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".j2bnli1czc.1.0.2.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".j2bnli1czc.1.0.2.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".j2bnli1czc.1.0.2.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".j2bnli1czc.1.0.2.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".j2bnli1czc.1.0.2.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.2.0.2.0.0"><span data-reactid=".j2bnli1czc.1.0.2.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".j2bnli1czc.1.0.2.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".j2bnli1czc.1.0.2.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".j2bnli1czc.1.0.2.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.2.0.3.0.0"><span data-reactid=".j2bnli1czc.1.0.2.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.2.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav><div accept-charset="UTF-8" class="searchBox searchBox--navbar" data-reactid=".j2bnli1czc.1.0.3"><form autocomplete="off" action="/search" class="searchBox__form" role="search" aria-label="Search for books to add to your shelves" data-reactid=".j2bnli1czc.1.0.3.0"><input class="searchBox__input searchBox__input--navbar" autocomplete="off" name="q" type="text" placeholder="Search books" aria-label="Search books" aria-controls="searchResults" data-reactid=".j2bnli1czc.1.0.3.0.0"/><input type="hidden" name="qid" value="" data-reactid=".j2bnli1czc.1.0.3.0.1"/><button type="submit" class="searchBox__icon--magnifyingGlass gr-iconButton searchBox__icon searchBox__icon--navbar" aria-label="Search" data-reactid=".j2bnli1czc.1.0.3.0.2"></button></form></div><div class="siteHeader__personal" data-reactid=".j2bnli1czc.1.0.4"><ul class="personalNav" data-reactid=".j2bnli1czc.1.0.4.0"><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.4.0.0"><div data-reactid=".j2bnli1czc.1.0.4.0.0.0"><div class="dropdown dropdown--notifications" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0"><a class="dropdown__trigger dropdown__trigger--notifications dropdown__trigger--personalNav" href="/notifications?ref=nav_my_notifs" role="button" aria-haspopup="true" aria-expanded="false" title="Notifications" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.0.0.0">2</span></span></a><div class="dropdown__menu dropdown__menu--notifications gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1"><div class="dropdown__container
                        gr-notifications
                        gr-notifications--sparse" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0"><div class="spinnerContainer" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0.0"><div class="spinner" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0.0.0"><div class="spinner__mask" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".j2bnli1czc.1.0.4.0.0.0.0.1.0.0.1">Loading…</div></div></div></div></div></div></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.4.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.4.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".j2bnli1czc.1.0.4.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.4.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.4.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".j2bnli1czc.1.0.4.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.4.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.4.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".j2bnli1czc.1.0.4.0.3.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.4.0.4"><div class="dropdown dropdown--profileMenu" data-reactid=".j2bnli1czc.1.0.4.0.4.0"><a class="dropdown__trigger dropdown__trigger--profileMenu dropdown__trigger--personalNav" href="/user/show/68156753-sebastiaan" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.4.0.4.0.0"><span class="headerPersonalNav__icon" data-reactid=".j2bnli1czc.1.0.4.0.4.0.0.0"><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".j2bnli1czc.1.0.4.0.4.0.0.0.1"/></span></a><div class="dropdown__menu dropdown__menu--profileMenu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1"><div class="siteHeader__subNav siteHeader__subNav--profile gr-box gr-box--withShadowLarge" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0"><span class="siteHeader__subNavLink gr-h3 gr-h3--noMargin" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.0"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.0.0"> </span><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.0.1">Sebastiaan</span><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.0.2"> </span></span><ul data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.0"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.4"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.4.0.0"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D2%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.b"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.b.0.0"><span data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".j2bnli1czc.1.0.4.0.4.0.1.0.1.e.0">Sign out</a></li></ul></div></div></div></li></ul></div><div class="siteHeader__topLevelItem siteHeader__topLevelItem--profileIcon" data-reactid=".j2bnli1czc.1.0.5"><span class="headerPersonalNav" data-ux-click="true" data-reactid=".j2bnli1czc.1.0.5.0"><a class="modalTrigger" role="button" aria-expanded="false" aria-haspopup="true" data-reactid=".j2bnli1czc.1.0.5.0.0"><span class="headerPersonalNav__icon" data-reactid=".j2bnli1czc.1.0.5.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".j2bnli1czc.1.0.5.0.0.0.0">2</span><img class="circularIcon circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".j2bnli1czc.1.0.5.0.0.0.1"/></span></a></span></div><div class="modal modal--overlay" tabindex="0" data-reactid=".j2bnli1czc.1.0.6"><div class="modal__content" data-reactid=".j2bnli1czc.1.0.6.0"><div class="modal__close" data-reactid=".j2bnli1czc.1.0.6.0.0"><button type="button" class="gr-iconButton" data-reactid=".j2bnli1czc.1.0.6.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_x-b06e4e308b9bd6ad1d0019e135dfa722.svg" data-reactid=".j2bnli1czc.1.0.6.0.0.0.0"/></button></div><div class="gr-genresForm" data-reactid=".j2bnli1czc.1.0.6.0.1"><div class="gr-genresForm__title" data-reactid=".j2bnli1czc.1.0.6.0.1.0">Follow Your Favorite Genres</div><div class="gr-genresForm__description" data-reactid=".j2bnli1czc.1.0.6.0.1.1">We use your favorite genres to make better book recommendations and tailor what you see in your Updates feed.</div><form action="/user/edit_fav_genres" data-remote="true" method="post" data-reactid=".j2bnli1czc.1.0.6.0.1.2"><div class="gr-genresForm__checkBoxes" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0"><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Art"><input name="favorites[Art]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Art.0"/><input name="favorites[Art]" type="checkbox" value="true" data-genre="Art" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Art.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Art.2">Art</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Biography"><input name="favorites[Biography]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Biography.0"/><input name="favorites[Biography]" type="checkbox" value="true" data-genre="Biography" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Biography.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Biography.2">Biography</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Business"><input name="favorites[Business]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Business.0"/><input name="favorites[Business]" type="checkbox" value="true" data-genre="Business" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Business.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Business.2">Business</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Children&#x27;s"><input name="favorites[Children&#x27;s]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Children&#x27;s.0"/><input name="favorites[Children&#x27;s]" type="checkbox" value="true" data-genre="Children&#x27;s" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Children&#x27;s.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Children&#x27;s.2">Children&#x27;s</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Christian"><input name="favorites[Christian]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Christian.0"/><input name="favorites[Christian]" type="checkbox" value="true" data-genre="Christian" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Christian.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Christian.2">Christian</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Classics"><input name="favorites[Classics]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Classics.0"/><input name="favorites[Classics]" type="checkbox" value="true" data-genre="Classics" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Classics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Classics.2">Classics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Comics"><input name="favorites[Comics]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Comics.0"/><input name="favorites[Comics]" type="checkbox" value="true" data-genre="Comics" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Comics.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Comics.2">Comics</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Cookbooks"><input name="favorites[Cookbooks]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Cookbooks.0"/><input name="favorites[Cookbooks]" type="checkbox" value="true" data-genre="Cookbooks" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Cookbooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Cookbooks.2">Cookbooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Ebooks"><input name="favorites[Ebooks]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Ebooks.0"/><input name="favorites[Ebooks]" type="checkbox" value="true" data-genre="Ebooks" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Ebooks.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Ebooks.2">Ebooks</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fantasy"><input name="favorites[Fantasy]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fantasy.0"/><input name="favorites[Fantasy]" type="checkbox" value="true" data-genre="Fantasy" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fantasy.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fantasy.2">Fantasy</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fiction"><input name="favorites[Fiction]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fiction.0"/><input name="favorites[Fiction]" type="checkbox" value="true" data-genre="Fiction" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Fiction.2">Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Graphic Novels"><input name="favorites[Graphic Novels]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Graphic Novels.0"/><input name="favorites[Graphic Novels]" type="checkbox" value="true" data-genre="Graphic Novels" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Graphic Novels.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Graphic Novels.2">Graphic Novels</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Historical Fiction"><input name="favorites[Historical Fiction]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Historical Fiction.0"/><input name="favorites[Historical Fiction]" type="checkbox" value="true" data-genre="Historical Fiction" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Historical Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Historical Fiction.2">Historical Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$History"><input name="favorites[History]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$History.0"/><input name="favorites[History]" type="checkbox" value="true" data-genre="History" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$History.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$History.2">History</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Horror"><input name="favorites[Horror]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Horror.0"/><input name="favorites[Horror]" type="checkbox" value="true" data-genre="Horror" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Horror.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Horror.2">Horror</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Memoir"><input name="favorites[Memoir]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Memoir.0"/><input name="favorites[Memoir]" type="checkbox" value="true" data-genre="Memoir" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Memoir.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Memoir.2">Memoir</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Music"><input name="favorites[Music]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Music.0"/><input name="favorites[Music]" type="checkbox" value="true" data-genre="Music" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Music.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Music.2">Music</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Mystery"><input name="favorites[Mystery]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Mystery.0"/><input name="favorites[Mystery]" type="checkbox" value="true" data-genre="Mystery" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Mystery.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Mystery.2">Mystery</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Nonfiction"><input name="favorites[Nonfiction]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Nonfiction.0"/><input name="favorites[Nonfiction]" type="checkbox" value="true" data-genre="Nonfiction" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Nonfiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Nonfiction.2">Nonfiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Poetry"><input name="favorites[Poetry]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Poetry.0"/><input name="favorites[Poetry]" type="checkbox" value="true" data-genre="Poetry" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Poetry.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Poetry.2">Poetry</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Psychology"><input name="favorites[Psychology]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Psychology.0"/><input name="favorites[Psychology]" type="checkbox" value="true" data-genre="Psychology" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Psychology.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Psychology.2">Psychology</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Romance"><input name="favorites[Romance]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Romance.0"/><input name="favorites[Romance]" type="checkbox" value="true" data-genre="Romance" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Romance.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Romance.2">Romance</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science"><input name="favorites[Science]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science.0"/><input name="favorites[Science]" type="checkbox" value="true" data-genre="Science" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science.2">Science</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science Fiction"><input name="favorites[Science Fiction]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science Fiction.0"/><input name="favorites[Science Fiction]" type="checkbox" value="true" data-genre="Science Fiction" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science Fiction.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Science Fiction.2">Science Fiction</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Self Help"><input name="favorites[Self Help]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Self Help.0"/><input name="favorites[Self Help]" type="checkbox" value="true" data-genre="Self Help" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Self Help.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Self Help.2">Self Help</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Sports"><input name="favorites[Sports]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Sports.0"/><input name="favorites[Sports]" type="checkbox" value="true" data-genre="Sports" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Sports.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Sports.2">Sports</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Thriller"><input name="favorites[Thriller]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Thriller.0"/><input name="favorites[Thriller]" type="checkbox" value="true" data-genre="Thriller" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Thriller.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Thriller.2">Thriller</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Travel"><input name="favorites[Travel]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Travel.0"/><input name="favorites[Travel]" type="checkbox" value="true" data-genre="Travel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Travel.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Travel.2">Travel</span></label><label class="gr-genresForm__genreLabel" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Young Adult"><input name="favorites[Young Adult]" type="hidden" value="false" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Young Adult.0"/><input name="favorites[Young Adult]" type="checkbox" value="true" data-genre="Young Adult" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Young Adult.1"/><span class="u-verticalAlignMiddle u-marginLeftTiny" data-reactid=".j2bnli1czc.1.0.6.0.1.2.0.$Young Adult.2">Young Adult</span></label></div><button type="submit" class="gr-button gr-button--large" data-reactid=".j2bnli1czc.1.0.6.0.1.2.1">Save</button></form></div></div></div><div class="modal modal--overlay modal--drawer" tabindex="0" data-reactid=".j2bnli1czc.1.0.7"><div data-reactid=".j2bnli1czc.1.0.7.0"><div class="modal__close" data-reactid=".j2bnli1czc.1.0.7.0.0"><button type="button" class="gr-iconButton" data-reactid=".j2bnli1czc.1.0.7.0.0.0"><img alt="Dismiss" src="//s.gr-assets.com/assets/gr/icons/icon_close_white-dbf4152deeef5bd3915d5d12210bf05f.svg" data-reactid=".j2bnli1czc.1.0.7.0.0.0.0"/></button></div><div class="modal__content" data-reactid=".j2bnli1czc.1.0.7.0.1"><div class="personalNavDrawer" data-reactid=".j2bnli1czc.1.0.7.0.1.0"><div class="personalNavDrawer__personalNavContainer" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0"><ul class="personalNav" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0"><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.0"><a href="/notifications?ref=nav_my_notifs" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.0.0"><span class="headerPersonalNav__icon
                       headerPersonalNav__icon--notifications" aria-label="Notifications" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.0.0.0"><span class="headerPersonalNav__flag" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.0.0.0.0">2</span></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.1"><a href="/topic?ref=nav_bar_discussions_pane_discussion&amp;discussion_filter=groups" title="My group discussions" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.1.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--discussions" aria-label="My group discussions" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.1.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.2"><a href="/message/inbox?ref=nav_my_messages" title="Messages" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.2.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--inbox" aria-label="Inbox" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.2.0.0"></span></a></li><li class="personalNav__listItem" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.3"><a href="/friend?ref=nav_my_friends" title="Friends" class="headerPersonalNav" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.3.0"><span class="headerPersonalNav__icon headerPersonalNav__icon--friendRequests" aria-label="Friend Requests" data-reactid=".j2bnli1czc.1.0.7.0.1.0.0.0.3.0.0"></span></a></li></ul></div><div class="personalNavDrawer__profileAndLinksContainer" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1"><div class="personalNavDrawer__profileContainer gr-mediaFlexbox gr-mediaFlexbox--alignItemsCenter" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0"><div class="gr-mediaFlexbox__media" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.0"><a href="/user/show/68156753-sebastiaan" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.0.0"><img class="circularIcon circularIcon--large circularIcon--border" src="https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png" alt="Sebastiaan" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.0.0.0"/></a></div><div class="gr-mediaFlexbox__desc" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--bold" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.1.0">Sebastiaan</a><div class="u-displayBlock" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.1.1"><a href="/user/show/68156753-sebastiaan" class="gr-hyperlink gr-hyperlink--naked" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.0.1.1.0">View profile</a></div></div></div><div class="personalNavDrawer__profileMenuContainer" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1"><ul data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0"><li role="menuitem Profile" class="menuLink" aria-label="Profile" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.0"><span data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.0.0"><a href="/user/show/68156753-sebastiaan" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.0.0.0">Profile</a></span></li><li role="menuitem Friends" class="menuLink" aria-label="Friends" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.3"><a href="/friend?ref=nav_profile_friends" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.3.0">Friends</a></li><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.4"><span data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.4.0"><a href="/group/list/68156753-sebastiaan?ref=nav_profile_groups" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.4.0.0"><span data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.4.0.0.0">Groups</span></a></span></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.5"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.5.0">Discussions</a></li><li role="menuitem Comments" class="menuLink" aria-label="Comments" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.6"><a href="/comment/list/68156753-sebastiaan?ref=nav_profile_comment" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.6.0">Comments</a></li><li role="menuitem Reading Challenge" class="menuLink" aria-label="Reading Challenge" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.7"><a href="https://www.goodreads.com/challenges/11634?ref=nav_profile_rc" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.7.0">Reading Challenge</a></li><li role="menuitem Kindle Notes &amp; Highlights" class="menuLink" aria-label="Kindle Notes &amp; Highlights" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.8"><a href="/notes?ref=nav_profile_knh" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.8.0">Kindle Notes &amp; Highlights</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.9"><a href="/quotes/list?ref=nav_profile_quotes" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.9.0">Quotes</a></li><li role="menuitem Favorite genres" class="menuLink" aria-label="Favorite genres" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.a"><a href="/user/edit_fav_genres?ref=nav_profile_favgenre&amp;return_url=%2Fwork%2Feditions%2F94024291%3Fpage%3D2%26per_page%3D10" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.a.0">Favorite genres</a></li><li role="menuitem Friends&#x27; recommendations" class="menuLink" aria-label="Friends&#x27; recommendations" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.b"><span data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.b.0"><a href="/recommendations/to_me?ref=nav_profile_friendrec" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.b.0.0"><span data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.b.0.0.0">Friends’ recommendations</span></a></span></li><li role="menuitem Account settings" class="menuLink" aria-label="Account settings" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.c"><a href="/user/edit?ref=nav_profile_settings" class="siteHeader__subNavLink u-topGrayBorder" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.c.0">Account settings</a></li><li role="menuitem Help" class="menuLink" aria-label="Help" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.d"><a href="/help?action_type=help_nav_bar&amp;ref=nav_profile_help" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.d.0">Help</a></li><li role="menuitem Sign out" class="menuLink" aria-label="Sign out" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.e"><a href="/user/sign_out?ref=nav_profile_signout" class="siteHeader__subNavLink" data-method="POST" data-reactid=".j2bnli1czc.1.0.7.0.1.0.1.1.0.e.0">Sign out</a></li></ul></div></div></div></div></div></div></div></div><div class="headroom-wrapper" data-reactid=".j2bnli1czc.2"><div style="position:relative;top:0;left:0;right:0;z-index:1;-webkit-transform:translateY(0);-ms-transform:translateY(0);transform:translateY(0);" class="headroom headroom--unfixed" data-reactid=".j2bnli1czc.2.0"><nav class="siteHeader__primaryNavSeparateLine gr-box gr-box--withShadow" data-reactid=".j2bnli1czc.2.0.0"><ul role="menu" class="siteHeader__menuList" data-reactid=".j2bnli1czc.2.0.0.0"><li class="siteHeader__topLevelItem siteHeader__topLevelItem--home" data-reactid=".j2bnli1czc.2.0.0.0.0"><a href="/?ref=nav_home" class="siteHeader__topLevelLink" data-reactid=".j2bnli1czc.2.0.0.0.0.0">Home</a></li><li class="siteHeader__topLevelItem" data-reactid=".j2bnli1czc.2.0.0.0.1"><a href="/review/list/68156753?ref=nav_mybooks" class="siteHeader__topLevelLink" data-reactid=".j2bnli1czc.2.0.0.0.1.0">My Books</a></li><li class="siteHeader__topLevelItem" data-reactid=".j2bnli1czc.2.0.0.0.2"><div class="primaryNavMenu primaryNavMenu--siteHeaderBrowseMenu ignore-react-onclickoutside" data-reactid=".j2bnli1czc.2.0.0.0.2.0"><a class="primaryNavMenu__trigger primaryNavMenu__trigger--siteHeaderBrowseMenu" href="/book?ref=nav_brws" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".j2bnli1czc.2.0.0.0.2.0.0"><span data-reactid=".j2bnli1czc.2.0.0.0.2.0.0.0">Browse ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1"><div class="siteHeader__browseMenuDropdown" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0"><ul class="siteHeader__subNav" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0"><li role="menuitem Recommendations" class="menuLink" aria-label="Recommendations" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.0"><a href="/recommendations?ref=nav_brws_recs" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.0.0">Recommendations</a></li><li role="menuitem Choice Awards" class="menuLink" aria-label="Choice Awards" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.1"><a href="/choiceawards?ref=nav_brws_gca" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.1.0">Choice Awards</a></li><li role="menuitem Genres" class="menuLink" aria-label="Genres" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.2"><a href="/genres?ref=nav_brws_genres" class="siteHeader__subNavLink siteHeader__subNavLink--genresIndex" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.2.0">Genres</a></li><li role="menuitem Giveaways" class="menuLink" aria-label="Giveaways" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.3"><a href="/giveaway?ref=nav_brws_giveaways" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.3.0">Giveaways</a></li><li role="menuitem New Releases" class="menuLink" aria-label="New Releases" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.4"><a href="/new_releases?ref=nav_brws_newrels" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.4.0">New Releases</a></li><li role="menuitem Lists" class="menuLink" aria-label="Lists" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.5"><a href="/list?ref=nav_brws_lists" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.5.0">Lists</a></li><li role="menuitem Explore" class="menuLink" aria-label="Explore" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.6"><a href="/book?ref=nav_brws_explore" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.6.0">Explore</a></li><li role="menuitem News &amp; Interviews" class="menuLink" aria-label="News &amp; Interviews" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.7"><a href="/news?ref=nav_brws_news" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.0.7.0">News &amp; Interviews</a></li></ul><div class="siteHeader__spotlight siteHeader__spotlight--browseMenu" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1"><div class="featuredGenres featuredGenres--sparse" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0"><div class="spinnerContainer" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0.0"><div class="spinner" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0.0.0"><div class="spinner__mask" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0.0.0.0"><div class="spinner__maskedCircle" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0.0.0.0.0"></div></div></div><div class="spinnerFallbackText" data-reactid=".j2bnli1czc.2.0.0.0.2.0.1.0.1.0.0.1">Loading…</div></div></div></div></div></div></div></li><li class="siteHeader__topLevelItem siteHeader__topLevelItem--community" data-reactid=".j2bnli1czc.2.0.0.0.3"><div class="primaryNavMenu ignore-react-onclickoutside" data-reactid=".j2bnli1czc.2.0.0.0.3.0"><a class="primaryNavMenu__trigger" href="/group?ref=nav_comm" role="button" aria-haspopup="true" aria-expanded="false" data-ux-click="true" data-reactid=".j2bnli1czc.2.0.0.0.3.0.0"><span data-reactid=".j2bnli1czc.2.0.0.0.3.0.0.0">Community ▾</span></a><div class="primaryNavMenu__menu gr-box gr-box--withShadowLarge" role="menu" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1"><ul class="siteHeader__subNav" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0"><li role="menuitem Groups" class="menuLink" aria-label="Groups" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.0"><a href="/group?ref=nav_comm_groups" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.0.0">Groups</a></li><li role="menuitem Discussions" class="menuLink" aria-label="Discussions" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.1"><a href="/topic?ref=nav_comm_discuss" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.1.0">Discussions</a></li><li role="menuitem Quotes" class="menuLink" aria-label="Quotes" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.2"><a href="/quotes?ref=nav_comm_quotes" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.2.0">Quotes</a></li><li role="menuitem Ask the Author" class="menuLink" aria-label="Ask the Author" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.3"><a href="/ask_the_author?ref=nav_comm_askauthor" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.3.0">Ask the Author</a></li><li role="menuitem People" class="menuLink" aria-label="People" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.4"><a href="/user/best_reviewers?ref=nav_comm_people" class="siteHeader__subNavLink" data-reactid=".j2bnli1czc.2.0.0.0.3.0.1.0.4.0">People</a></li></ul></div></div></li></ul></nav></div></div></header></div>
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
      new Ajax.Request('/shelf/add_to_shelf', {asynchronous:true, evalScripts:true, onSuccess:function(request){shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name)}, parameters:'book_id=' + book_id + '&name=' + shelf_name + '&a=' + action + '&authenticity_token=' + encodeURIComponent('ijiysL5v9C+jA/tWyk9NM26KUNP4nqSWy2ugvopQkw2ZQ86I/F3qvgqYDpcv8k/fLsHtr71QtDMB76cW95e4/w==')})
    }
  }

  function shelfSubmitted(request, book_id, checkbox_id, element_id, unique_id, shelf_name) {
    Element.update('shelfList68156753_' + book_id, request.responseText)
    afterShelfSave(checkbox_id, element_id, unique_id, shelf_name.escapeHTML())
  }

  function refreshGroupBox(group_id, book_id) {
    new Ajax.Updater('addGroupBooks' + book_id + '', '/group/add_book_box', {asynchronous:true, evalScripts:true, onSuccess:function(request){refreshGroupBoxComplete(request, book_id);}, parameters:'id=' + group_id + '&book_id=' + book_id + '&refresh=true' + '&authenticity_token=' + encodeURIComponent('ODiiy6AfDr+DjDEL/EL70mCWLdsANCiXuuy0Mhk5xMArQ97z4i0QLioXxMoZ//k+IN2Qp0X6ODJwaLOaZP7vMg==')})
  }
//]]>
</script>

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=61136370-el-libro-negro-de-las-horas">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=61688776-het-zwarte-boek-van-de-uren">Switch to This Edition</a>
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
        

          <a class="gr-button gr-button--small" rel="nofollow" data-method="post" href="/review/switch_edition/6460345706?book_id=61313780-el-llibre-negre-de-les-hores">Switch to This Edition</a>
      </div>
    </div>

	<div style="text-align: right; width: 100%">
		<div><a class="previous_page" rel="prev start" href="/work/editions/94024291?page=1&amp;per_page=10">« previous</a> <a rel="prev start" href="/work/editions/94024291?page=1&amp;per_page=10">1</a> <em class="current">2</em> <span class="next_page disabled">next »</span></div>

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
<div data-react-class="ReactComponents.LoginInterstitial" data-react-props="{&quot;allowFacebookSignIn&quot;:true,&quot;allowAmazonSignIn&quot;:true,&quot;overrideSignedOutPageCount&quot;:false,&quot;path&quot;:{&quot;signInUrl&quot;:&quot;/user/sign_in&quot;,&quot;signUpUrl&quot;:&quot;/user/sign_up&quot;,&quot;privacyUrl&quot;:&quot;/about/privacy&quot;,&quot;termsUrl&quot;:&quot;/about/terms&quot;,&quot;thirdPartyRedirectUrl&quot;:&quot;/user/new?connect_prompt=true&quot;}}"><noscript data-reactid=".wga1t78iwu" data-react-checksum="-1444605662"></noscript></div>


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
      ReactStores.GoogleAdsStore.initializeWith({"targeting":{"sid":"osid.174265687bb1cb6edc4444e39d7276f4","grsession":"osid.174265687bb1cb6edc4444e39d7276f4","surface":"desktop","signedin":"true","gr_author":"false","author":["29367407","283304","4470653","5898355","545","3487","4370565","8730","442240","1405152","8427407","108424","58","6252","8588","8534434","630","3120844","410653","2851725","4763","37272748","14184453","3354","5804101","88506","8349","6525349","2786093","1370283","76360","4721536","904939","20675225","1445909","73149","6979427","706255","1192311","7710","15862877","21632010","5780686","6535608","19976903","7705004","1864374","728092","1405767","7246482"],"genres":["1","107","64","244","411","144","67","97","2286","2352","84","1679","28","40","69","1870","29","2207","584","836","136","35","1049","2515","2091","552","6537","8263","1651","1098","831","1139","117","494","921","2287","25","22643","2038","24","72","352","92199","355","1007","262067","569","1105","14175","11231"],"Gender":"null","Age":"null"},"ads":{},"nativeAds":{}});  ReactStores.NotificationsStore.updateWith({"unreadCount":2,"unreadCountMore":false});
      ReactStores.CurrentUserStore.initializeWith({"currentUser":{"name":"Sebastiaan","profileUrl":"/user/show/68156753-sebastiaan","profileImage":"https://s.gr-assets.com/assets/nophoto/user/u_60x60-267f0ca0ea48fd3acfd44b95afa64f01.png","pendingRecsCount":0,"groupInvitesCount":0,"tempFriendRequestCount":0,"tempUnreadMessageCount":0}});
      ReactStores.FavoriteGenresStore.updateWith({"allGenres":[{"name":"Art","url":"/genres/art"},{"name":"Biography","url":"/genres/biography"},{"name":"Business","url":"/genres/business"},{"name":"Children's","url":"/genres/children-s"},{"name":"Christian","url":"/genres/christian"},{"name":"Classics","url":"/genres/classics"},{"name":"Comics","url":"/genres/comics"},{"name":"Cookbooks","url":"/genres/cookbooks"},{"name":"Ebooks","url":"/genres/ebooks"},{"name":"Fantasy","url":"/genres/fantasy"},{"name":"Fiction","url":"/genres/fiction"},{"name":"Graphic Novels","url":"/genres/graphic-novels"},{"name":"Historical Fiction","url":"/genres/historical-fiction"},{"name":"History","url":"/genres/history"},{"name":"Horror","url":"/genres/horror"},{"name":"Memoir","url":"/genres/memoir"},{"name":"Music","url":"/genres/music"},{"name":"Mystery","url":"/genres/mystery"},{"name":"Nonfiction","url":"/genres/non-fiction"},{"name":"Poetry","url":"/genres/poetry"},{"name":"Psychology","url":"/genres/psychology"},{"name":"Romance","url":"/genres/romance"},{"name":"Science","url":"/genres/science"},{"name":"Science Fiction","url":"/genres/science-fiction"},{"name":"Self Help","url":"/genres/self-help"},{"name":"Sports","url":"/genres/sports"},{"name":"Thriller","url":"/genres/thriller"},{"name":"Travel","url":"/genres/travel"},{"name":"Young Adult","url":"/genres/young-adult"}],"favoriteGenres":["Crime","Fantasy","Fiction","Historical fiction","Mystery","Science fiction","Thriller"]});
      ReactStores.TabsStore.updateWith({"communitySpotlight":"groups"});
    
    });
  //]]>
</script>

</body>
</html>
<!-- This is a random-length HTML comment: ecrwrmoolzdxphsdqprcklgbzzfmmmfavdmugrksthiyaofmgjjnbebiqdxfxcdipdcjrigmmlosipmagkkfyyvyiupbgchmorlyxdqnejxslruymltlkwgqykmryenoskpfnsshvqnsqjsfxqjpiccwaswnxwjrdkyyw -->`
	mockHTML3 := ``
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve different HTML responses based on the page number
		if strings.Contains(r.URL.RawQuery, "page=1&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML1))
		} else if strings.Contains(r.URL.RawQuery, "page=2&per_page") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML2))
		} else {
			w.WriteHeader(http.StatusNotFound) // No more pages
			// w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHTML3))
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
		for _, v := range got {
			fmt.Printf("{Isbn: \"%s\", Format: \"%s\", Language: \"%s\"},\n", v.Isbn, v.Format, v.Language)
		}
		t.Fatalf("Want: '%+v', Got: '%+v'", want, got)
	}
}

(function($){
  "use strict";

  var $window = $(window);

  $window.on('load', function() {

    // Preloader
    $('.loader').fadeOut();
    $('.loader-mask').delay(350).fadeOut('slow');

    $window.trigger("resize");

  });


  // Init
  initOwlCarousel();
  initFlickity();

  $window.on('resize', function() {
    hideSidenav();
    megaMenu();
  });



  /* Detect Browser Size
  -------------------------------------------------------*/
  var minWidth;
  if (Modernizr.mq('(min-width: 0px)')) {
    // Browsers that support media queries
    minWidth = function (width) {
      return Modernizr.mq('(min-width: ' + width + 'px)');
    };
  }
  else {
    // Fallback for browsers that does not support media queries
    minWidth = function (width) {
      return $window.width() >= width;
    };
  }

  /* Mobile Detect
  -------------------------------------------------------*/
  if (/Android|iPhone|iPad|iPod|BlackBerry|Windows Phone/i.test(navigator.userAgent || navigator.vendor || window.opera)) {
     $("html").addClass("mobile");
     $('.dropdown-toggle').attr('data-toggle', 'dropdown');
  }
  else {
    $("html").removeClass("mobile");
  }

  /* IE Detect
  -------------------------------------------------------*/
  if(Function('/*@cc_on return document.documentMode===10@*/')()){ $("html").addClass("ie"); }


  /* Sticky Navigation
  -------------------------------------------------------*/
  var $stickyNav = $('.nav--sticky');
  var $nav = $('.nav');

  $window.scroll(function(){
    scrollToTop();    

    if ($window.scrollTop() > 2 & minWidth(992)) {
      $stickyNav.addClass('sticky');
      $nav.addClass('sticky');
    } else {
      $stickyNav.removeClass('sticky');
      $nav.removeClass('sticky');
    }

  });


  function stickyNavRemove() {
    if (!minWidth(992)) {
      $stickyNav.removeClass('sticky');
    }
  } 


  /* Mobile Navigation
  -------------------------------------------------------*/
  var $sidenav = $('#sidenav'),
      $main = $('#main'),
      $navIconToggle = $('#nav-icon-toggle');


  $navIconToggle.on('click', function(e) {
    e.stopPropagation();
    $(this).toggleClass('nav-icon-toggle--is-open');
    $sidenav.toggleClass('sidenav--is-open');   
    $main.toggleClass('main--is-open');
  });

  function resetNav() {
    $navIconToggle.removeClass('nav-icon-toggle--is-open');
    $sidenav.removeClass('sidenav--is-open');
    $main.removeClass('main--is-open');
  }

  function hideSidenav() {
    if( minWidth(992) ) {
      resetNav();
      setTimeout( megaMenu, 500 );
    }
  }

  $main.on('click', function() {    
    resetNav();
  });


  var $dropdownTrigger = $('.nav__dropdown-trigger'),
      $navDropdownMenu = $('.nav__dropdown-menu'),
      $navDropdown = $('.nav__dropdown');


  if ( $('html').hasClass('mobile') ) {

    $('body').on('click',function() {
      $navDropdownMenu.addClass('hide-dropdown');
    });

    $navDropdown.on('click', '> a', function(e) {
      e.preventDefault();
    });

    $navDropdown.on('click',function(e) {
      e.stopPropagation();
      $navDropdownMenu.removeClass('hide-dropdown');
    });
  }



  /* Sidenav Menu
  -------------------------------------------------------*/

  $('.sidenav__menu-toggle').on('click', function(e) {
    e.preventDefault();
    
    var $this = $(this);
    
    $this.parent().siblings().removeClass('sidenav__menu--is-open');
    $this.parent().siblings().find('li').removeClass('sidenav__menu--is-open');
    $this.parent().find('li').removeClass('sidenav__menu--is-open');
    $this.parent().toggleClass('sidenav__menu--is-open');       
    
    if ($this.next().hasClass('show')) {
      $this.next().removeClass('show').slideUp(350);    
    } else {
      $this.parent().parent().find('li .sidenav__menu-dropdown').removeClass('show').slideUp(350);
      $this.next().toggleClass('show').slideToggle(350);
    }
  });


  /* Mega Menu
  -------------------------------------------------------*/
  function megaMenu(){
    $('.nav__megamenu').each(function () {
      var $this = $(this);

      $this.css('width', $('.container').width());
      var offset = $this.closest('.nav__dropdown').offset();
      offset = offset.left;
      var containerOffset = $(window).width() - $('.container').outerWidth();
      containerOffset = containerOffset /2;
      offset = offset - containerOffset - 15;
      $this.css('left', -offset);
    });
  }
  

  /* Accordion
  -------------------------------------------------------*/
  var $accordion = $('.accordion');

  function toggleChevron(e) {
    $(e.target)
      .prev('.accordion__heading')
      .find("a")
      .toggleClass('accordion--is-open accordion--is-closed');
  }
  $accordion.on('hide.bs.collapse', toggleChevron);
  $accordion.on('show.bs.collapse', toggleChevron);



  /* Tabs
  -------------------------------------------------------*/
  $('.tabs__link').on('click', function(e) {
    var currentAttrValue = $(this).attr('href');
    $('.tabs__content ' + currentAttrValue).stop().fadeIn(1000).siblings().hide();
    $(this).parent('li').addClass('active').siblings().removeClass('active');
    e.preventDefault();
  });


  /* Owl Carousel
  -------------------------------------------------------*/
  function initOwlCarousel(){

    // Featured Posts
    $("#owl-hero").owlCarousel({
      center: true,
      items:1,
      loop:true,
      nav:true,
      navSpeed: 500,
      navText: ['<i class="ui-arrow-left">','<i class="ui-arrow-right">']
    });


    // Gallery post
    $("#owl-single").owlCarousel({
      items:1,
      loop:true,
      nav:true,
      animateOut: 'fadeOut',
      navText: ['<i class="ui-arrow-left">','<i class="ui-arrow-right">']
    });

    // Testimonials
    $("#owl-testimonials").owlCarousel({
      items:1,
      loop:true,
      nav:true,
      dots:false,
      navText: ['<i class="ui-arrow-left">','<i class="ui-arrow-right">']
    });

  };



  /* Flickity Slider
  -------------------------------------------------------*/

  function initFlickity() {

    // main large image (shop product)
    $('#gallery-main').flickity({
      cellAlign: 'center',
      contain: true,
      wrapAround: true,
      autoPlay: false,
      prevNextButtons: true,
      percentPosition: true,
      imagesLoaded: true,
      lazyLoad: 1,
      pageDots: false,
      selectedAttraction : 0.1,
      friction: 0.6,
      rightToLeft: false,
      arrowShape: 'M 25,50 L 65,90 L 70,90 L 30,50  L 70,10 L 65,10 Z'
    });

    // thumbs
    $('#gallery-thumbs').flickity({
      asNavFor: '#gallery-main',
      contain: true,
      cellAlign: 'left',
      wrapAround: false,
      autoPlay: false,
      prevNextButtons: false,
      percentPosition: true,
      imagesLoaded: true,
      pageDots: false,
      selectedAttraction : 0.1,
      friction: 0.6,
      rightToLeft: false
    });

    var $gallery = $('.mfp-hover');

    $gallery.on( 'dragStart.flickity', function( event, pointer ) {
      $(this).addClass('is-dragging');
    })

    $gallery.on( 'dragEnd.flickity', function( event, pointer ) {
      $(this).removeClass('is-dragging');
    })

    $gallery.magnificPopup({
      delegate: '.lightbox-img, .lightbox-video',
      callbacks: {
        elementParse: function(item) {
        if(item.el.context.className === 'lightbox-video') {
            item.type = 'iframe';
          } else {
            item.type = 'image';
          }
        }
      },    
      type: 'image',
      closeBtnInside:false,
      gallery:{
        enabled:true
      }
    });
  }  


  /* Payment Method Accordion
  -------------------------------------------------------*/
  var methods = $(".payment_methods > li > .payment_box").hide();
  methods.first().slideDown("easeOutExpo");
  
  $(".payment_methods > li > input").change(function(){
    var current = $(this).parent().children(".payment_box");
    methods.not(current).slideUp("easeInExpo");
    $(this).parent().children(".payment_box").slideDown("easeOutExpo");
    
    return false;     
  });
  

  /* Quantity
  -------------------------------------------------------*/
  $(function() {

    // Increase
    jQuery(document).on('click', '.plus', function(e) {
      e.preventDefault();
      var quantityInput = jQuery(this).parents('.quantity').find('input.qty'),
      newValue = parseInt(quantityInput.val(), 10) + 1,
      maxValue = parseInt(quantityInput.attr('max'), 10);

      if (!maxValue) {
        maxValue = 9999999999;
      }

      if ( newValue <= maxValue ) {
        quantityInput.val(newValue);
        quantityInput.change();
      }
    });

    // Decrease
    jQuery(document).on('click', '.minus', function(e) {
      e.preventDefault();
      var quantityInput = jQuery(this).parents('.quantity').find('input.qty'),
      newValue = parseInt(quantityInput.val(), 10) - 1,
      minValue = parseInt(quantityInput.attr('min'), 10);
      
      if (!minValue) {
        minValue = 1;
      }

      if ( newValue >= minValue ) {
        quantityInput.val(newValue);
        quantityInput.change();
      }
    });

  });


  /* Sign In Popup
  -------------------------------------------------------*/
  $('#top-bar__sign-in, .product__quickview').magnificPopup({
    type: 'ajax',
    alignTop: false,
    overflowY: 'scroll',
    removalDelay: 300,
    mainClass: 'mfp-fade',
    callbacks: {
      ajaxContentAdded: function() {
        initFlickity();        
      },
      close: function() {
        var $productImgHolder = $('.product__img-holder');
        $productImgHolder.addClass('processed');
        function removeProcessing() {
          $productImgHolder.removeClass('processed');
        }
        setTimeout( removeProcessing, 50 );
      }
    }
  });


  /* Quickview
  -------------------------------------------------------*/
  $('.product__quickview').on('click', function() {
    var product = $('.product');

    function removeProcessing() {
      product.removeClass('processing');
    }

    product.addClass('processing');
    setTimeout( removeProcessing, 500 );

  });



  /* ---------------------------------------------------------------------- */
  /*  Contact Form
  /* ---------------------------------------------------------------------- */

  var submitContact = $('#submit-message'),
    message = $('#msg');

  submitContact.on('click', function(e){
    e.preventDefault();

    var $this = $(this);
    
    $.ajax({
      type: "POST",
      url: 'contact.php',
      dataType: 'json',
      cache: false,
      data: $('#contact-form').serialize(),
      success: function(data) {

        if(data.info !== 'error'){
          $this.parents('form').find('input[type=text],input[type=email],textarea,select').filter(':visible').val('');
          message.hide().removeClass('success').removeClass('error').addClass('success').html(data.msg).fadeIn('slow').delay(5000).fadeOut('slow');
        } else {
          message.hide().removeClass('success').removeClass('error').addClass('error').html(data.msg).fadeIn('slow').delay(5000).fadeOut('slow');
        }
      }
    });
  });


  /* Scroll to Top
  -------------------------------------------------------*/
  function scrollToTop() {
    var scroll = $window.scrollTop();
    var $backToTop = $("#back-to-top");
    if (scroll >= 50) {
      $backToTop.addClass("show");
    } else {
      $backToTop.removeClass("show");
    }
  }

  $('a[href="#top"]').on('click',function(){
    $('html, body').animate({scrollTop: 0}, 1000, "easeInOutQuint");
    return false;
  });

})(jQuery);
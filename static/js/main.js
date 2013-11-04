var spinner_options = {
	lines: 15, // The number of lines to draw
	length: 24, // The length of each line
	width: 9, // The line thickness
	radius: 41, // The radius of the inner circle
	corners: 1, // Corner roundness (0..1)
	rotate: 0, // The rotation offset
	direction: 1, // 1: clockwise, -1: counterclockwise
	color: '#000', // #rgb or #rrggbb or array of colors
	speed: 1, // Rounds per second
	trail: 60, // Afterglow percentage
	shadow: false, // Whether to render a shadow
	hwaccel: false, // Whether to use hardware acceleration
	className: 'spinner', // The CSS class to assign to the spinner
	zIndex: 2e9, // The z-index (defaults to 2000000000)
	top: '40px', // Top position relative to parent in px
	left: 'auto' // Left position relative to parent in px
};

require.config({
	paths: {
		'jquery': 'libs/jquery',
		'jquery.bootstrap': 'libs/bootstrap',
		'jquery.nprogress':'libs/nprogress',
		'mustache':'libs/mustache',
		'modernizr':'libs/modernizr',
		'alertify':'libs/alertify',
		'spinner':'libs/spin',
		'html5':'libs/html5',
		'jquery.tablesorter':'libs/tablesorter',
		'jquery.tablesorter.widget':'libs/tablesorter.widgets'
	},
	shim :{
		'jquery.bootstrap': {
			deps: ['jquery']
		},
		'jquery.nprogress': {
			deps: ['jquery']
		},
		'jquery.tablesorter':{
			deps: ['jquery']
		},
		'jquery.tablesorter.widget':{
			deps: ['jquery.tablesorter']
		}
	},
	waitSeconds: 15
});

require(
	['jquery',
	'jquery.bootstrap',
	'jquery.nprogress',
	'mustache',
	'modernizr',
	'spinner',
	'html5',
	'alertify',
	'jquery.tablesorter',
	'jquery.tablesorter.widget'],
	function($, bootstrap, nprogress,mustache, mod, spinner, html5, alertify, tablesorter, tablesorter_widgets){

		$(function(){
			if(!Modernizr.inputtypes.date){
				Modernizr.load({
					test: Modernizr.inputtypes.date,
					nope: ['http://ajax.googleapis.com/ajax/libs/jqueryui/1.8.7/jquery-ui.min.js', '/css/jquery-ui.css'],
					complete: function () {
						$('input[type=date]').datepicker({
							dateFormat: 'yy-mm-dd'
						});
					}
				});
			}

			$('.focus').select();

			NProgress.configure({
				showSpinner: false
			});

			$(document).on('click','.toggler',function(){
				var ref = $(this).data('target');
				$(ref).slideToggle();
			});

			$(document).on('click','.nav-tabs li a',function(e){
				if (e.preventDefault){
					e.preventDefault();
				}
				$(this).tab('show');
			});

			// Get find out if we have any elements that require CKEditor
			// Drop it on the page and bind to elements.
			var ck_els = $('.ckeditor').get();
			if(ck_els.length > 0){
				$.ajax({
					url: '//ck.curtmfg.com/ckeditor.js',
					type:'GET',
					dataType:'script',
					success:function(data,status,xhr){
						$.each(ck_els,function(i, el){
							CKEDITOR.replace(el,{
								filebrowserImageUploadUrl: '/file/CKUpload',
								filebrowserImageBrowseUrl: '/file/CKIndex',
								filebrowserWindowWidth: '640',
								filebrowserWindowHeight: '480'
							});
						});
					}
				});
			}

			/* Setup TableSorter */
			$.extend($.tablesorter.themes.bootstrap, {
				// these classes are added to the table. To see other table classes available,
				// look here: http://twitter.github.com/bootstrap/base-css.html#tables
				table      : 'table table-bordered',
				header     : 'bootstrap-header', // give the header a gradient background
				footerRow  : '',
				footerCells: '',
				icons      : '', // add "icon-white" to make them white; this icon class is added to the <i> in the header
				sortNone   : 'bootstrap-icon-unsorted',
				sortAsc    : 'icon-chevron-up',
				sortDesc   : 'icon-chevron-down',
				active     : '', // applied when column is sorted
				hover      : '', // use custom css here - bootstrap class may not override it
				filterRow  : '', // filter row class
				even       : '', // odd row zebra striping
				odd        : ''  // even row zebra striping
			});
			// call the tablesorter plugin and apply the uitheme widget
			$(".sortable").tablesorter({
				// this will apply the bootstrap theme if "uitheme" widget is included
				// the widgetOptions.uitheme is no longer required to be set
				theme : "bootstrap",

				widthFixed: true,

				headerTemplate : '{content} {icon}', // new in v2.7. Needed to add the bootstrap icon!

				// widget code contained in the jquery.tablesorter.widgets.js file
				// use the zebra stripe widget if you plan on hiding any rows (filter widget)
				widgets : [ "uitheme", "filter", "zebra" ],

				widgetOptions : {
				// using the default zebra striping class name, so it actually isn't included in the theme variable above
				// this is ONLY needed for bootstrap theming if you are using the filter widget, because rows are hidden
				zebra : ["even", "odd"],

				// reset filters button
				filter_reset : ".reset"

				// set the uitheme widget to use the bootstrap theme class names
				// this is no longer required, if theme is set
				// ,uitheme : "bootstrap"

				}
			});

			// return $;
		});
	}
);
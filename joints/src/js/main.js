require.config({
	baseUrl: 'js',
	paths: {
		'domReady': 'libs/domReady',
		'messenger': 'libs/messenger',
		'messenger-theme': 'libs/messenger-theme-future',
		'angular': 'libs/angular',
		'angular-ui-router': 'libs/angular-ui-router',
		'jquery': 'libs/jquery.min',
		'bootstrap': 'libs/bootstrap.min',
		'angular-bootstrap': 'libs/ui-bootstrap-tpls',
        'parsley': 'libs/parsley.min',
		'underscore': 'libs/underscore-min',
		'd3': 'libs/d3.min',
		'node-uuid' : 'libs/uuid',
		'jsoneditor': 'libs/jsoneditor.min',
		'select2': 'libs/select2.min',
		'js-yaml': 'libs/js-yaml.min',
		'clipboard': 'libs/clipboard.min',
		'jquery.autocomplete': 'libs/jquery.autocomplete',
		'ace-theme-dawn': 'libs/theme-dawn',
		'ace-mode-golang': 'libs/mode-golang',
		'angular-websocket': 'libs/angular-websocket',
		'app': 'app',
		'router': 'router/devopsRoute'
		
	},
	shim: {
		'angular': {
			exports: 'angular'
		},
		'underscore': {
			exports: '_'
		},
		'angular-bootstrap': {
			deps: ['angular']
		},
		'messenger' : {
            exports: 'messenger',
            deps: ['jquery']
		},
		'messenger-theme' : {
            deps: ['messenger']
		},
		'angular-ui-router' : {
            deps: ['angular']
		},		
		'parsley' : {
            exports: 'parsley'
		},
		'd3' : {
           exports: 'd3'
		},
		'bootstrap': {
			deps: ['jquery']
		},
		'jquery.autocomplete': {
			deps: ['jquery']
		},
		'jsoneditor': {
			deps: ['jquery']
		},
		'ace-theme-dawn': {
			deps: ['jsoneditor']
		},
		'ace-mode-golang': {
			deps: ['jsoneditor']
		},
		'angular-websocket': {
			exports: 'angular-websocket',
			deps: ['angular']
		}


		
	}
});
/**
 * bootstraps angular onto the window.document node
 */
define([
	'require','angular','app','underscore','jquery','jsoneditor','node-uuid', 'angular-ui-router',
	'bootstrap','angular-bootstrap', 'router','messenger','messenger-theme','parsley','d3',
	'select2','js-yaml','clipboard','jquery.autocomplete','ace-mode-golang','ace-theme-dawn',
	'angular-websocket'
], function(require, angular, app, _, $, jsoneditor, uuid) {
	'use strict';
	require(['domReady!'], function(document) {
		window.JSONEditor = jsoneditor;
		window.uuid = uuid;
		angular.bootstrap(document, ['DevOps']);
	});
});

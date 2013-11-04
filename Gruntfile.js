/*global module: false */
module.exports = function(grunt) {
	"use strict";
	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),
		compass: {
			dist: {
				options: {
					sassDir: 'static/scss',
					cssDir: 'static/css'
				}
			}
		},
		jshint: {
			options: {
				curly: true,
				eqeqeq: true,
				eqnull: true,
				browser: true,
				loopfunc:true,
				globals: {
					jQuery: true
				},
				ignores:['static/js/*.min.js', 'static/js/libs/**'] // We're not going to jshint libs, too much work :\
			},
			dist: {
				src: ['Gruntfile.js', 'static/js/**/*.js'] 
			}
		},
		uglify: {
			options: {
				banner: '/* <%= pkg.name %> - version <%= pkg.version %> - ' +
						'<%= grunt.template.today("dd-mm-yyyy") %>\n' +
						'<%= pkg.description %>\n ' +
						'<%= grunt.template.today("yyyy") %> <%= pkg.author.name %> ' +
						'- <%= pkg.author.email %> */\n'
			},
			my_target: {
				files: {
					'static/js/main.min.js': ['static/js/main.js']
				}
			}
		},
		watch: {
			options: {
				livereload:true
			},
			files: ['static/scss/*', 'static/js/*.js', 'templates/**/*.html', 'Gruntfile.js'],
			tasks: ['jshint', 'uglify', 'compass']
		},
		connect: {
			test: {
				options: {
					port: 9001,
					keepalive: true
				}
			}
		}
	});
 
	grunt.loadNpmTasks('grunt-contrib-compass');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-jasmine');
 
	grunt.registerTask('default', ['jshint', 'compass', 'watch']);

	grunt.registerTask('dist', ['jshint','uglify','compass']);
};
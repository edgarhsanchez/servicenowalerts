var gulp = require('gulp');
var sync = require('gulp-sync')(gulp).sync;
var child = require('child_process');
var util = require('gulp-util');

/*
 * Override gulp.src() for nicer error handling.
 */
var src = gulp.src;
gulp.src = function() {
  return src.apply(gulp, arguments)
    .pipe(plumber(function(error) {
      util.log(util.colors.red(
        'Error (' + error.plugin + '): ' + error.message
      ));
      notifier.notify({
        title: 'Error (' + error.plugin + ')',
        message: error.message.split('\n')[0]
      });
      this.emit('end');
    })
  );
};

gulp.task('app:install', function(){
    var build = child.spawnSync('go', ['install']);
    if (build.stderr.length) {
      var lines = build.stderr.toString()
        .split('\n').filter(function(line) {
          return line.length
        });
      for (var l in lines)
        util.log(util.colors.red(
          'Error (go install): ' + lines[l]
        ));
      notifier.notify({
        title: 'Error (go install)',
        message: lines
      });
    }
    return build;
})

gulp.task('app:watch', function(){
    gulp.watch([
        '*/**/*.go',
    ], function(){
        return gulp.start([
            'app:install',
        ])
    })
})

gulp.task('app:build', function() {
  var build = child.spawn('go', ['build', '-ldflags', '-H=windowsgui']);
    if (build.stderr.length) {
      var lines = build.stderr.toString()
        .split('\n').filter(function(line) {
          return line.length
        });
      for (var l in lines)
        util.log(util.colors.red(
          'Error (go build): ' + lines[l]
        ));
      notifier.notify({
        title: 'Error (go build)',
        message: lines
      });
    }
    return build;
})
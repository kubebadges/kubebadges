import 'package:jovial_svg/jovial_svg.dart';

var scalableImageCache = ScalableImageCache(size: 100);

class Constants {
  static const String buildType =
      String.fromEnvironment('BUILD_TYPE', defaultValue: 'dev');

  static const baseAPI = buildType == "prod" ? "" : "http://127.0.0.1:8090";
}

import 'package:get/get.dart';
import 'package:ui/app/controller/badge_controller.dart';
import 'package:ui/app/controller/index_controller.dart';

class IndexBinding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut<IndexController>(() => IndexController());
    Get.lazyPut<BadgeController>(() => BadgeController());
  }
}

import 'package:get/get.dart';
import 'package:ui/app/controller/badge_controller.dart';
import 'package:ui/app/service/app_service.dart';

class IndexController extends GetxController {
  BadgeController defaultController = Get.find();
  AppService appService = Get.find();

  final _selectedIndex = 0.obs;
  int get selectedIndex => _selectedIndex.value;
  set selectedIndex(int value) => _selectedIndex.value = value;

  IndexController() {
    appService.loadConfig();
  }

  void refreshPage() async {
    if (selectedIndex == 0) {
      defaultController.loadData(true);
    } else if (selectedIndex == 1) {}
  }
}

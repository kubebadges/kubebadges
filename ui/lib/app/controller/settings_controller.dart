import 'package:get/get.dart';
import 'package:ui/app/service/app_service.dart';

class SettingsController extends GetxController {
  AppService appService = Get.find();

  final _badgeBaseURL = "".obs;
  String get badgeBaseURL => _badgeBaseURL.value;
  set badgeBaseURL(String value) => _badgeBaseURL.value = value;

  SettingsController() {
    appService.eventBus.on<UpdateConfigEvent>().listen((event) {
      badgeBaseURL = appService.getBadgeBaseURL();
    });
  }

  void updateBadgeURL(String url) async {
    appService.updateConfig(appService.getConfig().copyWith(badgeBaseURL: url));
  }
}

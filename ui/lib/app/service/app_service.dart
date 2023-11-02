import 'dart:convert';

import 'package:get/get.dart';
import 'package:ui/app/api/api.dart';
import 'package:ui/app/model/model.dart';
import 'package:event_bus/event_bus.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

class UpdateConfigEvent {}

class AppService extends GetxService {
  EventBus eventBus = EventBus();
  Api api = Get.find();

  KubeBadgeConfig config = KubeBadgeConfig(
    badgeBaseURL: '',
  );

  void loadConfig() async {
    var result = await api.getBadgeConfig();
    if (!result.hasError) {
      setConfig(result.body!);
      eventBus.fire(UpdateConfigEvent());
    }
  }

  void updateConfig(KubeBadgeConfig kubeBadgeConfig) async {
    EasyLoading.show(status: 'Loading...');
    var result = await api.updateBadgeConfig(kubeBadgeConfig);
    if (!result.hasError) {
      setConfig(result.body!);
      eventBus.fire(UpdateConfigEvent());
      EasyLoading.showToast('Success');
    } else {
      if (result.bodyString != null) {
        final Map<String, dynamic> error = jsonDecode(result.bodyString!);
        EasyLoading.showError(
          'config update failed : ${error['error']}',
          duration: const Duration(milliseconds: 2500),
        );
      }
    }
    EasyLoading.dismiss();
  }

  void setConfig(KubeBadgeConfig kubeBadgeConfig) {
    config = kubeBadgeConfig;
  }

  KubeBadgeConfig getConfig() {
    return config;
  }

  String getBadgeBaseURL() {
    if (config.badgeBaseURL.isEmpty) {
      return "";
    }

    if (config.badgeBaseURL.endsWith("/")) {
      return config.badgeBaseURL.substring(0, config.badgeBaseURL.length - 1);
    }

    return config.badgeBaseURL;
  }

  Future<Response<List<KubeBadge>>> listNodes(bool force) =>
      api.listNodes(force);

  Future<Response<List<KubeBadge>>> listNamespace(bool force) =>
      api.listNamespace(force);

  Future<Response<List<KubeBadge>>> listDeployments(String name, bool force) =>
      api.listDeployments(name, force);

  Future<Response> updateBadge(Map<String, dynamic> data) =>
      api.updateBadge(data);
}

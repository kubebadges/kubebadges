import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'package:ui/app/api/api.dart';
import 'package:ui/app/model/model.dart';

class BadgeController extends GetxController {
  Api api = Get.find();

  final _nodeList = <KubeBadge>[].obs;
  List<KubeBadge> get nodeList => _nodeList;
  set nodeList(List<KubeBadge> value) => _nodeList.value = value;

  final _namespaceList = <KubeBadge, List<KubeBadge>>{}.obs;
  Map<KubeBadge, List<KubeBadge>> get namespaceList => _namespaceList;
  set namespaceList(Map<KubeBadge, List<KubeBadge>> value) =>
      _namespaceList.value = value;

  BadgeController() {
    loadData(false);
  }

  void loadData(bool force) async {
    listNodes(force);
    listNamespace(force);
  }

  void listNodes(bool force) async {
    namespaceList.clear();
    var response = await api.listNodes(force);
    if (!response.status.hasError) {
      nodeList = response.body!;
    }
  }

  void listNamespace(bool force) async {
    EasyLoading.show(status: 'Loading...');
    namespaceList.clear();
    var namespaces = await api.listNamespace(force);
    if (!namespaces.status.hasError) {
      var temp = <KubeBadge, List<KubeBadge>>{};
      var deploymentsRequests = <Future>[];
      for (var namespace in namespaces.body!) {
        deploymentsRequests.add(
          api.listDeployments(namespace.name, force).then((deployments) {
            if (!deployments.status.hasError) {
              if (deployments.body!.isNotEmpty) {
                temp[namespace] = deployments.body!;
              }
            }
          }),
        );
      }
      await Future.wait(deploymentsRequests);
      namespaceList = temp;
    }
    EasyLoading.dismiss();
  }

  Future<bool> updateBadgePublic(KubeBadge kubeBadge, bool allowed) async {
    EasyLoading.show(status: 'Loading...');
    var response = await api.updateBadge({
      "key": kubeBadge.key,
      "allowed": allowed,
    });
    if (response.status.hasError) {
      final Map<String, dynamic> error = response.body;
      EasyLoading.dismiss();
      EasyLoading.showError(
        'Badge update failed : ${error['error']}',
        duration: const Duration(milliseconds: 2500),
      );
      return false;
    }

    int index = nodeList.indexWhere((element) => element.key == kubeBadge.key);
    if (index != -1) {
      nodeList[index] = kubeBadge.copyWith(allowed: allowed);
    }

    for (var namespace in namespaceList.keys) {
      int index = namespaceList[namespace]!
          .indexWhere((element) => element.key == kubeBadge.key);
      if (index != -1) {
        var updatedList = List<KubeBadge>.from(namespaceList[namespace]!);
        updatedList[index] = kubeBadge.copyWith(allowed: allowed);
        namespaceList[namespace] = updatedList;
      }
    }
    namespaceList = Map.from(namespaceList);
    EasyLoading.dismiss();
    EasyLoading.showToast(
      'Badge updated success',
      duration: const Duration(seconds: 1),
      toastPosition: EasyLoadingToastPosition.bottom,
    );
    return true;
  }
}

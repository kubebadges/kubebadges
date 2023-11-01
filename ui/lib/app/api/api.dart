import 'package:get/get.dart';
import 'package:ui/app/config/constant.dart';
import 'package:ui/app/model/model.dart';

class Api extends GetConnect {
  @override
  void onInit() {
    super.onInit();
    httpClient.baseUrl = Constants.baseAPI;
  }

  Future<Response<List<KubeBadge>>> listNodes(bool force) {
    return get('/api/nodes?force=$force', decoder: (data) {
      return (data as List).map((item) => KubeBadge.fromJson(item)).toList();
    });
  }

  Future<Response<List<KubeBadge>>> listNamespace(bool force) {
    return get('/api/namespaces?force=$force', decoder: (data) {
      return (data as List).map((item) => KubeBadge.fromJson(item)).toList();
    });
  }

  Future<Response<List<KubeBadge>>> listDeployments(String name, bool force) {
    return get('/api/deployments/$name?force=$force', decoder: (data) {
      return (data as List).map((item) => KubeBadge.fromJson(item)).toList();
    });
  }

  Future<Response> updateBadge(Map<String, dynamic> data) {
    return post('/api/badge', data);
  }
}

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/api/api.dart';
import 'package:ui/app/routes/app_pages.dart';
import 'package:ui/app/ui/pages/index/index_binding.dart';
import 'package:ui/app/ui/pages/index/index_page.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

Future<void> main() async {
  await initServices();
  runApp(
    GetMaterialApp(
      debugShowCheckedModeBanner: false,
      defaultTransition: Transition.fadeIn,
      getPages: AppPages.pages,
      home: const IndexPage(),
      initialBinding: IndexBinding(),
      builder: EasyLoading.init(),
    ),
  );
}

Future<void> initServices() async {
  Get.lazyPut(() => Api());
}

import 'package:get/route_manager.dart';
import 'package:ui/app/routes/app_routes.dart';
import 'package:ui/app/ui/pages/index/index_binding.dart';
import 'package:ui/app/ui/pages/index/index_page.dart';

abstract class AppPages {
  static final pages = [
    GetPage(
      name: Routes.index,
      page: () => const IndexPage(),
      binding: IndexBinding(),
    ),
  ];
}

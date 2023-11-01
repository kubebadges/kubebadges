import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/ui/pages/default/default_page.dart';
import 'package:ui/app/ui/pages/document/document_page.dart';
import 'package:ui/app/ui/pages/generator/generator_page.dart';
import 'package:ui/app/controller/index_controller.dart';
import 'package:ui/app/ui/pages/index/menu_drawer.dart';
import 'package:ui/app/ui/pages/settings/settings_page.dart';

class IndexPage extends GetView<IndexController> {
  const IndexPage({super.key});

  @override
  Widget build(BuildContext context) {
    // 获取屏幕的宽度
    double screenWidth = MediaQuery.of(context).size.width;

    // 设置侧边栏的显示和隐藏
    bool isDrawerVisible = screenWidth >= 600;

    return Scaffold(
      appBar: AppBar(
        title: const Text('KubeBadges'),
        leading: !isDrawerVisible
            ? Builder(
                builder: (context) => IconButton(
                  icon: const Icon(Icons.menu),
                  onPressed: () {
                    Scaffold.of(context).openDrawer();
                  },
                ),
              )
            : null,
      ),
      floatingActionButton: Obx(
        () => Visibility(
          visible:
              controller.selectedIndex == 1 || controller.selectedIndex == 0,
          child: FloatingActionButton(
            onPressed: () {
              controller.refreshPage();
            },
            child: const Icon(Icons.refresh),
          ),
        ),
      ),
      body: SafeArea(
        child: Row(
          children: <Widget>[
            if (isDrawerVisible)
              SizedBox(
                width: 230,
                child: MenuDrawer(
                  controller: controller,
                  shouldPop: false,
                ),
              ),
            Expanded(
              child: Padding(
                padding: const EdgeInsets.all(24.0),
                child: Obx(
                  () => IndexedStack(
                    index: controller.selectedIndex,
                    children: const [
                      DefaultPage(),
                      GeneratorPage(),
                      DocumentPage(),
                      SettingsPage(),
                    ],
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
      drawer: isDrawerVisible
          ? null
          : MenuDrawer(
              controller: controller,
              shouldPop: true,
            ),
    );
  }
}

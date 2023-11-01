import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/ui/pages/default/badge_card.dart';
import 'package:ui/app/ui/pages/default/card_title.dart';
import 'package:ui/app/controller/badge_controller.dart';
import 'package:ui/app/ui/widgets/badge_dialog.dart';

class DefaultPage extends GetView<BadgeController> {
  const DefaultPage({super.key});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const CardTitle(title: "Nodes"),
          Obx(
            () => BadgeCard(
              items: [...controller.nodeList],
              kubeBadge: null,
              onTap: (e) {
                showDialog(
                  context: context,
                  builder: (BuildContext context) {
                    return BadgeSettingDialog(kubeBadge: e);
                  },
                );
              },
            ),
          ),
          const CardTitle(title: "Deployments"),
          Obx(
            () => ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              itemCount: controller.namespaceList.length,
              itemBuilder: (context, index) {
                var namespace =
                    controller.namespaceList.entries.elementAt(index);
                return BadgeCard(
                  items: namespace.value,
                  kubeBadge: namespace.key,
                  onTap: (e) {
                    showDialog(
                      context: context,
                      builder: (BuildContext context) {
                        return BadgeSettingDialog(kubeBadge: e);
                      },
                    );
                  },
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/controller/settings_controller.dart';
import 'package:ui/app/ui/pages/settings/badge_url_input_dialog.dart';
import 'package:ui/app/ui/pages/settings/setting_section_view.dart';

class SettingsPage extends GetView<SettingsController> {
  const SettingsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: [
        SettingSection(
          title: "Badge Access Settings",
          children: [
            ListTile(
              title: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Text(
                      "Badge Base URL",
                      style: TextStyle(
                          color: Colors.black,
                          fontWeight: FontWeight.w500,
                          fontSize: 16),
                    ),
                    Padding(
                      padding: const EdgeInsets.only(top: 6),
                      child: Obx(
                        () => Text(
                          controller.badgeBaseURL,
                          style: const TextStyle(
                            color: Colors.grey,
                            fontWeight: FontWeight.w500,
                            fontSize: 14,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ),
              trailing: const Icon(Icons.edit),
              onTap: () {
                showDialog(
                  context: context,
                  builder: (context) {
                    return BadgeURLInputDialog(
                      defaultURL: controller.badgeBaseURL,
                      onSuccess: (url) async {
                        controller.updateBadgeURL(url);
                      },
                    );
                  },
                );
              },
            )
          ],
        ),
      ],
    );
  }
}

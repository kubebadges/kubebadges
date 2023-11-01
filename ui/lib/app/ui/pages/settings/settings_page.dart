import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/ui/pages/settings/settings_controller.dart';

class SettingsPage extends GetView<SettingsController> {
  const SettingsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('SettingsPage'),
      ],
    );
  }
}

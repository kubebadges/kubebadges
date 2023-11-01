import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/ui/pages/generator/generator_controller.dart';

class GeneratorPage extends GetView<GeneratorController> {
  const GeneratorPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('GeneratorPage'),
      ],
    );
  }
}

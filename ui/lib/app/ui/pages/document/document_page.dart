import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/ui/pages/document/document_controller.dart';

class DocumentPage extends GetView<DocumentController> {
  const DocumentPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('DocumentPage'),
      ],
    );
  }
}

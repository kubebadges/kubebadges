import 'package:clipboard/clipboard.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/app/controller/badge_controller.dart';
import 'package:ui/app/model/model.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

class BadgeSettingDialog extends StatefulWidget {
  final KubeBadge kubeBadge;
  const BadgeSettingDialog({super.key, required this.kubeBadge});

  @override
  State<BadgeSettingDialog> createState() => _BadgeSettingDialogState();
}

class _BadgeSettingDialogState extends State<BadgeSettingDialog> {
  BadgeController get controller => Get.find<BadgeController>();
  bool _switchValue = false;

  @override
  void initState() {
    super.initState();
    _switchValue = widget.kubeBadge.allowed;
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text(
        'Deployment Badge',
        style: TextStyle(
          fontWeight: FontWeight.w500,
        ),
      ),
      actions: <Widget>[
        Padding(
          padding: const EdgeInsets.only(right: 8, bottom: 8, left: 8),
          child: TextButton(
            onPressed: () {
              FlutterClipboard.copy(
                      controller.getBadgeBaseURL() + widget.kubeBadge.badge)
                  .then((value) {
                EasyLoading.showToast(
                  'Copied to Clipboard!',
                  duration: const Duration(seconds: 2),
                  toastPosition: EasyLoadingToastPosition.bottom,
                );
              });
            },
            child: const Text('Copy Badge URL'),
          ),
        ),
      ],
      content: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisSize: MainAxisSize.min,
            children: <Widget>[
              const Icon(Icons.public, size: 16, color: Colors.grey),
              const SizedBox(width: 2),
              const Text('Private'),
              Padding(
                padding: const EdgeInsets.only(left: 8, right: 8),
                child: CupertinoSwitch(
                  value: _switchValue,
                  onChanged: (bool value) async {
                    var result = await controller.updateBadgePublic(
                        widget.kubeBadge, value);
                    if (result) {
                      setState(() {
                        _switchValue = value;
                      });
                    }
                  },
                ),
              ),
              const Text('Public'),
              const SizedBox(width: 2),
              const Icon(Icons.public, size: 16, color: Colors.green),
            ],
          ),
          const Padding(
            padding: EdgeInsets.only(top: 4.0),
            child: Text(
              'Turn on to allow external API access',
              style: TextStyle(
                fontSize: 12.0,
                color: Colors.grey,
              ),
            ),
          ),
          const Padding(
            padding: EdgeInsets.only(top: 24),
            child: Text(
              "Badge URL",
              style: TextStyle(
                fontWeight: FontWeight.w500,
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 8),
            child: SelectableText(
              controller.getBadgeBaseURL() + widget.kubeBadge.badge,
              style: TextStyle(
                fontSize: 14.0,
                color: Colors.grey.shade700,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

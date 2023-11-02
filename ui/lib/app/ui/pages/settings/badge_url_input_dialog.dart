import 'package:flutter/material.dart';

class BadgeURLInputDialog extends StatefulWidget {
  final Function(String) onSuccess;
  final String defaultURL;

  const BadgeURLInputDialog(
      {super.key, required this.onSuccess, this.defaultURL = ""});
  @override
  State<BadgeURLInputDialog> createState() => _BadgeURLInputDialogState();
}

class _BadgeURLInputDialogState extends State<BadgeURLInputDialog> {
  final TextEditingController editController = TextEditingController();
  String? errorText;

  @override
  void initState() {
    super.initState();
    editController.text = widget.defaultURL;
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Edit Badge Base URL'),
      content: SizedBox(
        width: 300,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: editController,
              onChanged: (value) {
                setState(() {
                  if (value.isEmpty) {
                    errorText = null;
                  } else if (!value.startsWith('http://') &&
                      !value.startsWith('https://')) {
                    errorText = 'Must start with http:// or https://';
                  } else {
                    errorText = null;
                  }
                });
              },
              decoration: InputDecoration(
                hintText: "Start with http:// or https://",
                errorText: errorText,
              ),
            ),
            const SizedBox(height: 10),
            const Text(
              'Set a domain here to auto-prepend it to badge URLs. Once set, copied badge URLs from this system will include the domain for easy use.',
              style: TextStyle(
                fontSize: 14.0,
                color: Colors.grey,
              ),
            ),
          ],
        ),
      ),
      actions: <Widget>[
        TextButton(
          child: const Text('Cancel'),
          onPressed: () {
            Navigator.of(context).pop();
          },
        ),
        TextButton(
          child: const Text('OK'),
          onPressed: () {
            if (errorText == null) {
              if (widget.defaultURL == editController.text) {
                Navigator.of(context).pop();
                return;
              }
              widget.onSuccess(editController.text);
              Navigator.of(context).pop();
            }
          },
        ),
      ],
    );
  }
}

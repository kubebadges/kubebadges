import 'package:flutter/material.dart';
import 'package:ui/app/controller/index_controller.dart';

class MenuDrawer extends StatelessWidget {
  final IndexController controller;
  final bool shouldPop;

  const MenuDrawer({
    Key? key,
    required this.controller,
    required this.shouldPop,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          ListTile(
            leading: const Icon(Icons.badge),
            title: const Text('Default Badges'),
            onTap: () {
              controller.selectedIndex = 0;
              if (shouldPop) {
                Navigator.of(context).pop();
              }
            },
          ),
          ListTile(
            leading: const Icon(Icons.create),
            title: const Text('Badges Generator'),
            onTap: () {
              controller.selectedIndex = 1;
              if (shouldPop) {
                Navigator.of(context).pop();
              }
            },
          ),
          ListTile(
            leading: const Icon(Icons.book),
            title: const Text('Document'),
            onTap: () {
              controller.selectedIndex = 2;
              if (shouldPop) {
                Navigator.of(context).pop();
              }
            },
          ),
          ListTile(
            leading: const Icon(Icons.settings),
            title: const Text('Settings'),
            onTap: () {
              controller.selectedIndex = 3;
              if (shouldPop) {
                Navigator.of(context).pop();
              }
            },
          ),
        ],
      ),
    );
  }
}

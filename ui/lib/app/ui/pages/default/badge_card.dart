import 'package:flutter/material.dart';
import 'package:ui/app/model/model.dart';
import 'package:ui/app/ui/widgets/badge.dart';

class BadgeCard extends StatelessWidget {
  final List<KubeBadge> items;
  final KubeBadge? kubeBadge;
  final void Function(KubeBadge) onTap;

  const BadgeCard(
      {super.key, required this.items, this.kubeBadge, required this.onTap});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (kubeBadge != null)
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Text(
              kubeBadge!.name,
              style: const TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.w400,
                color: Colors.grey,
              ),
            ),
          ),
        Card(
          elevation: 2.0,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8.0), // 设置圆角
          ),
          child: SizedBox(
            width: double.infinity,
            child: Padding(
              padding: const EdgeInsets.all(24.0),
              child: Wrap(
                spacing: 24,
                runSpacing: 4,
                children: items
                    .map(
                      (e) => KubeBadgeView(
                        badge: e,
                        onTap: () => onTap(e),
                      ),
                    )
                    .toList(), // node badges
              ),
            ),
          ),
        ),
      ],
    );
  }
}

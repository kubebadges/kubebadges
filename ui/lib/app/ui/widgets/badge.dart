import 'package:flutter/material.dart';
import 'package:jovial_svg/jovial_svg.dart';
import 'package:ui/app/config/constant.dart';
import 'package:ui/app/model/model.dart';

class KubeBadgeView extends StatefulWidget {
  final KubeBadge badge;
  final VoidCallback onTap;
  const KubeBadgeView({super.key, required this.badge, required this.onTap});

  @override
  State<KubeBadgeView> createState() => _KubeBadgeViewState();
}

class _KubeBadgeViewState extends State<KubeBadgeView> {
  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        widget.onTap();
      },
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            ScalableImageWidget.fromSISource(
              si: ScalableImageSource.fromSvgHttpUrl(
                Uri.parse(Constants.baseAPI + widget.badge.badge),
                warnF: null,
                // ignore: deprecated_member_use
                warn: false,
                httpHeaders: {
                  "Accept":
                      "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
                },
              ),
              cache: scalableImageCache,
            ),
            const SizedBox(width: 4),
            Icon(
              Icons.public,
              size: 20,
              color: widget.badge.allowed ? Colors.blue : Colors.grey.shade400,
            ),
          ],
        ),
      ),
    );
  }
}

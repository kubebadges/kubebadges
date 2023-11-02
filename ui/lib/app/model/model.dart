class KubeBadge {
  final String kind;
  final String name;
  final String badge;
  final bool allowed;
  final String key;

  KubeBadge({
    required this.kind,
    required this.name,
    required this.badge,
    required this.allowed,
    required this.key,
  });

  factory KubeBadge.fromJson(Map<String, dynamic> json) {
    return KubeBadge(
      kind: json['kind'],
      name: json['name'],
      badge: json['badge'],
      allowed: json['allowed'],
      key: json['key'],
    );
  }

  copyWith({
    String? kind,
    String? name,
    String? badge,
    bool? allowed,
    String? key,
  }) {
    return KubeBadge(
      kind: kind ?? this.kind,
      name: name ?? this.name,
      badge: badge ?? this.badge,
      allowed: allowed ?? this.allowed,
      key: key ?? this.key,
    );
  }
}

class KubeBadgeConfig {
  final String badgeBaseURL;

  KubeBadgeConfig({
    required this.badgeBaseURL,
  });

  factory KubeBadgeConfig.fromJson(Map<String, dynamic> json) {
    return KubeBadgeConfig(
      badgeBaseURL: json['badge_base_url'] ?? '',
    );
  }

  Map<String, dynamic> toJson() {
    return {
      "badge_base_url": badgeBaseURL,
    };
  }

  copyWith({
    String? badgeBaseURL,
  }) {
    return KubeBadgeConfig(
      badgeBaseURL: badgeBaseURL ?? this.badgeBaseURL,
    );
  }
}

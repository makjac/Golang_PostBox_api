
![header_image][header_image_url]

[![Flutter][flutter_badge]][flutter_url]
[![Golang][golang_badge]][golang_url]
[![Gin-Gonic][gin-gonic_badge]][gin-gonic_url]
[![PostgreSQL][postgresql_badge]][postgresql_url]
[![Mailjet][mailjet_badge]][mailjet_url]

# PostBox Courier Platform API

**Backend system for a digital parcel management platform**  

*Written in Golang - Powers the PostBox Flutter application*

## Overview

This repository contains the **backend API** for the PostBox platform, a digital solution for automated parcel management inspired by InPost's locker systems. The API is built with **Golang** and provides the core functionality for the [PostBox Flutter application](https://github.com/makjac/Flutter_PostBox_app), enabling users to manage parcels, track deliveries, and handle their profiles.

## Key Features

### User Profile Management
- **Authentication**:
  - JWT-based registration and login
  - Secure password hashing
  - Email verification for account activation
- **Contact Details**:
  - Update phone number and email
  - Manage multiple contact addresses
- **Address Book**:
  - Add, edit, or delete delivery addresses

### Parcel Management
- **Package Tracking**:
  - Real-time package location updates
- **Shipment Analytics**:
  - Historical shipment data

## Technical Architecture

### System Overview

![architecture_umage_][architecture_umage_url]

### Core Components

| Layer        | Technology Stack                                |
|--------------|-------------------------------------------------|
| **Frontend** | Flutter (Dart >=2.17.0 <3.0.0, Flutter >=3.0.0) |
| **Backend**  | Golang 1.16 + Gin-Gonic + Go-PG ORM             |
| **Database** | PostgreSQL 12                                   |
| **Infra**    | Google Cloud Platform (Compute Engine)          |
| **Services** | Mailjet API v3 + Cloud Storage                  |

<!-- end:excluded_rules_table -->

[header_image_url]:https://raw.githubusercontent.com/makjac/images/refs/heads/main/postbox/postbox_golang_header.png
[architecture_umage_url]:https://raw.githubusercontent.com/makjac/images/refs/heads/main/postbox/postbox_architecture.png

[flutter_badge]:https://img.shields.io/badge/Flutter-3.x-blue?logo=flutter
[flutter_url]:https://flutter.dev
[golang_badge]:https://img.shields.io/badge/Go-1.16-blue?logo=go
[golang_url]:https://golang.org
[gin-gonic_badge]:https://img.shields.io/badge/Gin_Gonic-v1.7.7-green?logo=go
[gin-gonic_url]:https://github.com/gin-gonic/gin
[postgresql_badge]:https://img.shields.io/badge/PostgreSQL-12%2B-blue?logo=postgresql
[postgresql_url]:https://www.postgresql.org
[mailjet_badge]:https://img.shields.io/badge/Mailjet-API_v3-orange?logo=mailjet
[mailjet_url]:https://www.mailjet.com

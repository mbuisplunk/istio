//go:build integ
// +build integ

//  Copyright Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package util

const (
	// Server certificate, private key and CA certificate
	TLSServerCertA = `-----BEGIN CERTIFICATE-----
MIIEOjCCAiKgAwIBAgIBADANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEP
MA0GA1UECAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYw
FAYDVQQDDA0qLmV4YW1wbGUuY29tMB4XDTIwMDgxMjIwMTIzNFoXDTMwMDgxMDIw
MTIzNFowGDEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKiuLJdO/Tq1oq/TTL1uI1vZAsUYh8jjIvVlcBXGPzXJ
bVTlBz937bZOOOOuXcLrzlM6l0d97Hbhkq67V7ziSCKffMt3ZptiS5ERy+a94jqT
kCianHpQSdvo4AFGZHLxDkL2QkEvuOvDmUfGmfu9cqoHYFcL9MJKlGvcW6Ae/QSb
BFLxdRfVGfsHC8w4rhmU2LbaABbB+HVJ8zPKMjEofoZebTUVpatQk//2XnM7VV4r
YrbrIyqbA2VWQIU0Ne4kQhKGc6Wh49mjV5gALjnWSBST+H4eBcoKhFHyqJF0WyhX
Ikstryd0m4Co/VEsJ2GXS7YJ5Fsc1z/5R0/h+ufqIxECAwEAAaNTMFEwCQYDVR0T
BAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMB
MBgGA1UdEQQRMA+CDSouZXhhbXBsZS5jb20wDQYJKoZIhvcNAQELBQADggIBADll
MgadJZQJg+CjguEYxZUZPuSsQdi449GsWIb0VzBHXGW1Lzx13lJMEVPG2QwS1YN/
Fje1se5T2LbUsWYJCYRY0Ig6lr9xbHmo0joV0aXvXPhE3om/SdvjuW8M92BwT2Nt
zcXeAJXLj1Uc8a1zS/+2CnjIO1B3d3QX+poK3i2djb4zIdo0TD1u6LE9F0eoFemr
xcM0383HCuQV3uDlkYuC7Y+K4ZLIZc2sZMvHTz0dLTEQzYj5N/auXYzF4doc9vcT
BRvMGKxecFzele9D5KI6ORUKw7e4K9apXX9t05fnLTrKTeM9krz1P9NGN9qPLECl
ePLeweJcFPT8I0Q+JLOtSkGzGStEyWT46/xE0epiRL3+uRahZYunzyEo9zLjp78Y
b/2QYiZo7W127gmefmOsEwRFP/teaoPWhz0MlmjAgvlAYcYJoM//3/yRzAT3LSm1
tfK8GQo1+/uTyswtq9Bzk8t9W8QoBPd8URhvzuUStmYAZ7DSz2egcRQeTAv4ACxf
lVK9o0UQdPvBWBSeuQY+nWwlHEjbHPH3eUe7bB0yxbYTT/khW2xtriGDC0qaasga
om/crtDsPuVe1W+yMlr30U31+JFQdksSUXF2Q8Um8NvNKoF4qSJLYvQ6pYeM8M4q
rrqXm5bbMyjaU1TFbyng7IAOeZbwsh5Hl6Zt0yR0
-----END CERTIFICATE-----`
	TLSServerKeyA = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAqK4sl079OrWir9NMvW4jW9kCxRiHyOMi9WVwFcY/NcltVOUH
P3fttk44465dwuvOUzqXR33sduGSrrtXvOJIIp98y3dmm2JLkRHL5r3iOpOQKJqc
elBJ2+jgAUZkcvEOQvZCQS+468OZR8aZ+71yqgdgVwv0wkqUa9xboB79BJsEUvF1
F9UZ+wcLzDiuGZTYttoAFsH4dUnzM8oyMSh+hl5tNRWlq1CT//ZecztVXititusj
KpsDZVZAhTQ17iRCEoZzpaHj2aNXmAAuOdZIFJP4fh4FygqEUfKokXRbKFciSy2v
J3SbgKj9USwnYZdLtgnkWxzXP/lHT+H65+ojEQIDAQABAoIBAGdepZDsJR8vZE1f
re2Aa0bEDICAceXX2/qKeU++t8nccJXP7MsmUZShBxfwKSFkpII4q7ByKNEJl3xg
7nmgKhidqDqAJ28do7V5NH1o7BR9jaB0b2Q9g6WyZmQhoTkXJGhAxYVxG1/P9EUa
Hg85n+U2tAFkIFBYp/AZkETl+KQy5HEM1MKPORyiGp0CK8m+dqS25/GfsOGY9PsN
KOxcoQ7Ozxvmtyo4Xk46RJriAFi7RZs7yFq6Yui7rr+vyvUSiWtkRZauQklxDBz6
j6Hx00zxRjKmB6HsAaROSPbxbbJwek/tPqQ6OVk1Xx2YkajK7q0Ccd8fMqZdTyZq
ycQHkOkCgYEA1xuXGbcOd9CJNxIswPne543fqM0zHba3E9DV6vvTaj3mF/2lJFDl
pm+ukH3VUDRmaoT5EijaZkfQNnByu/rBrtInJ7ghpv9lALEBMQspUr+/4MYmRm59
3mulL3edAwKcDEyD+cTd7Z03G5TKLmM01d9utQtSDOcuv+M2lUxtByMCgYEAyL8n
OXVgnEBsePT27LQmr6bE2DiRGE+TH7b6UjAOr7pABhvASmBSum0UHglYiYzgrqwK
ZJg6u5BlNzOadlZhHmZ96XQJEpFt17y32WT5cMtfNVVlWy5loMkYuZZNDwYJiNrh
jjZ2AgTAcqp5J2bd0SE3GsBXNxSZsW6e02qnajsCgYBzF28EMj6Kesg/7/iEE+1g
5TbrAUe9OTrBXnZmcPDQImhPI8ZNJP+KyqyS4NDWRgMaMSV2wub/3KKZzhaFX4hr
ukN66/kNeg6XjR0/GWK1xKSsZjiqbInqDJxoRk57AtVtey8N33iCnyCSIE0A/tGR
MSfxtCnlf3gy8SYOQcMAEwKBgBAeOPoGEKG7EnqLwJ+0vz5vN3Lc59l52ig7utGK
hkNZwvY6mC/gmAPb6jeXLIKuywN7UMJO7rhMmbPa8tX5jYaxV/68kFXrU1R1FJXh
451I1vYjSyYNDZ/hRZjxFnLUW0Ofv2h5uvPiickrItslCT5XWmNNejMz9jsm1J3I
/HQtAoGBAIl6vVWJdb0IFx/ioM9tw2PKD+IAi2MabeJ8XlA1Ov+O7lBUbDtqW4QG
Th8FapWamztvAgCXhY0ZB/mh5YN/wRSJLwqu8dGl9+3ajpa2zW7i+k2lJMMadsOe
x3C091Se8r6uof/Y8Jt4WLRY0q8kM7aMf1LIK89CoH7MYctnZp8i
-----END RSA PRIVATE KEY-----`
	TLSClientCertA = `-----BEGIN CERTIFICATE-----
MIIEOjCCAiKgAwIBAgIBADANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEP
MA0GA1UECAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYw
FAYDVQQDDA0qLmV4YW1wbGUuY29tMB4XDTIwMDgxMjIwMTIzNFoXDTMwMDgxMDIw
MTIzNFowGDEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAN5uozkENNgiyNjUoK+W+9fv2ROmZ/8EQTHnTdyL7x9W
do7xkG1k/l93V5nXEhXYtxc5IM6g5ulQuTYJA6aAuvSP1OCqas38dXQWxoe0h0Ju
NtkFAoP7aNCnHY9oJqaTH2l3InkwLvqqs3yJy9Vrh5GfcKt4gbTYzsoRqNFAXXJF
Kp1u2aezEzJG+4cdjJUcT0PaqXxwoK4KhuyhU4NWhrrgTvVCzsPNY4Q7YYZhimY8
s5oOpDzQEltOXYnBbwQvrpyOX9Z9umaEvJ71KiN3M6qiazpotIecYTcaeolUHG6A
YcIOYioeMvrMJ+vMwKcVdgdgvUVxVEVeLspsECPp6hECAwEAAaNTMFEwCQYDVR0T
BAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMB
MBgGA1UdEQQRMA+CDSouZXhhbXBsZS5jb20wDQYJKoZIhvcNAQELBQADggIBAJ8Z
ixHdx/ewlEMQqi1mM08FZdiLD83bBeLLKo4XH+GdQ82UcFUhYvK8dbRlXlsTrAEd
FbmVMz2ZLA/ePwwR3J1fLwMDPEm5VsD3hENwNAg6+iQJo2TCd2QVXZxPYXTjPfZI
iNp0BOpXDP1Tx1KWXxBF2kYejg1DUmtudWIaz54oxF0BBeRLPQT4ebWI3XhXXPKN
yrDWBzASsLEWWzugvH2Wz2Z3nJYjfe8FvgYAOltRgiqYLCwfLqenm4X68AdiM1go
fatTwOjkmHRHPXsoEr6IdV81g/jTavCj2cXLCnly3rLe7W5q+3L/OUSDL80uY26b
Gch45LUQTQypuMoRP17l4CzOkXVfV4+k6eawzmH/enSy5TXWTHwY6CWaY32Yit3D
ry0unhzh2kH2A2BQxFFmzh+Hqd9XA8MzYC0a8dttI15Yc60Zp8W0Ycr7kB0VvPQg
/KiwYKH93ZcdvzwSl/KwJXhrCrO2rvtMFrG6hm7qrnJjkwT44E5AEKcLxjnfcuLt
VZo2Q+F4WnXH034mKu5EFlZGdBscpQlg52gASNkNRS50Yhj+czZ7cxLEe2NwvZlR
TqzKoVd56akCaX/MY+XqVBV+JlPMfHRPesORcx3K+88jTTewOWkKYWhBdQqq+kpm
GXf//vFaX+iKzT5fO11pbKRFLpi1XvinElhUY9Ga
-----END CERTIFICATE-----`
	TLSClientKeyA = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA3m6jOQQ02CLI2NSgr5b71+/ZE6Zn/wRBMedN3IvvH1Z2jvGQ
bWT+X3dXmdcSFdi3FzkgzqDm6VC5NgkDpoC69I/U4Kpqzfx1dBbGh7SHQm422QUC
g/to0Kcdj2gmppMfaXcieTAu+qqzfInL1WuHkZ9wq3iBtNjOyhGo0UBdckUqnW7Z
p7MTMkb7hx2MlRxPQ9qpfHCgrgqG7KFTg1aGuuBO9ULOw81jhDthhmGKZjyzmg6k
PNASW05dicFvBC+unI5f1n26ZoS8nvUqI3czqqJrOmi0h5xhNxp6iVQcboBhwg5i
Kh4y+swn68zApxV2B2C9RXFURV4uymwQI+nqEQIDAQABAoIBACtF/v9CsD00VzrH
8xL2Hrzd7q+B+k2EuGJLC7zShIzmN16L5HtD4LEV+Lb3Po99f/FGrflFGGdfywiW
LF2iQbz+ln226dx1NyyEmc8g8gechZMnBFxE9xlQydFyg8J4cO4FvXEyTJF70Dt0
W3GD3/YWvMa/RQy3+VMaqGs9SHDIYy55eLUuDOChsTOkxlijONlz4rc93fuiyaJk
s81S9y1b+iPLhhNDgvFsPMTqCI8Va8N3pactsP6kEiIeennmbyPDvXz4o+Tcwiyj
nBVMQKLq1ljXjC4MFVpCFXgDmY0vsfGqQ+mbFJiwxxC6H5tcZJ/vGWpuOiMkicxC
Y5Il0NkCgYEA72MyAZGIzWXq7usK0yriXPxTb9FkMtULkuXrEEqV0pzVwNlzKE7z
MUWUOHX5Yn6xWoctN0hN83sDmqqL+8+J34AKv6hH6dEy0+rDj9s4joKct8LQE2Ct
P7tnkBGDoxKSk419XNhs4iQgPDkA3PRCVTRJ3637ya6lUncyxdq0+IMCgYEA7d45
peCe7Vvy9ku4QEYDdr5MJrEFLR/wB48T8dK9KOMRWgeede5gEzmRywij9cQqohnr
e4P2Netk40ZRv6ge3HCBMO6E03wMViS3tkshF0x+yevaOa0NgF30N2gXcKEYG65H
Gbis0GvoM3f7oRax4O2tv3vCEaYg9+/fOMRMxtsCgYEAhIf7eGjVPznsnkdcX9p/
JhHAM8jOW2IdnT9TK1rPpaGUqcb5fVhwRVgLHlMaNVCE6eSqwM9z8JerQCaph0i4
QReHW+JTpfxe0npHuM4aerOPBiGBxyif4gfj5xv7L/4pTi1oOU3MwkpT2mOIucs1
KizftUmCe38IIibB4k4aIPcCgYB6BG6p6mlb9LjTItBnGAMa3E8P4ep31VQ5bGmc
uK+T07bI6fOdJyRdeECTV+FTZsc/0+/5sh2QVymvdBjnKYR7K7L5uMCmA5IIdllO
x4c3/mNjjPqqb1znpkpqSZkMi3ZkbKfIR/VODKxWCownfA7PBN0FxqQsjX2H8aId
6oQ/1QKBgGVSncKl/BLfCXvBov7PNuLzxIdsLJS+hz45wiaWXmNyFpc2lTGZUdmh
Krf83cczjeACOZRaEzT7SDxOeD1EWX4EHeklj7v1aOMEQHbfkmjxmA0/PRo/8aAb
dBZuh0HcDad7OqdbwqH274Vt4cU6kxl5zPwNri/DIgJSR/hLIf4l
-----END RSA PRIVATE KEY-----`
	CaCertA = `-----BEGIN CERTIFICATE-----
MIIFozCCA4ugAwIBAgIUHMmGAs+HrfBCPuLWxcBGPEhtnU4wDQYJKoZIhvcNAQEL
BQAwVDELMAkGA1UEBhMCVVMxDzANBgNVBAgMBkRlbmlhbDEOMAwGA1UEBwwFRXRo
ZXIxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTAeFw0yMDA4
MTIyMDEyMzRaFw0zMDA4MTAyMDEyMzRaMFQxCzAJBgNVBAYTAlVTMQ8wDQYDVQQI
DAZEZW5pYWwxDjAMBgNVBAcMBUV0aGVyMQwwCgYDVQQKDANEaXMxFjAUBgNVBAMM
DSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCk
P0GyB53mvSL6lw0fxnQFj716E95FkXwGDSbNvyHCbbjp4FSmZ9fuztVhsNwVBko7
RSWSUQATQCMEL44AtYyhpUqsC2V08XN36AVEkzyTbHYMVBT5s0q8M9850DMF97YU
SjzzWv5+fo8jIw2KKSJNhc66WLKFQ9/mA50HvghrKjFrIofzAS6b1mPICDBj+rMO
loAtG5i01cKH71rqVXlN8gi5mhQDpAA5A0h4gfPEZv+UAKGFaqDPqcCgbzN/2KRt
5OrWRRh3pnEXHGfEILHbYinzzRWIij3vhU5/E+PISt5pCmFZvoPWw+V/cCh3TnAI
cSvoI7vFWyNwIvDzCpXkK7UHCAeqxUe+nLNNvkVLeXcTMh/yD4wlOizJ4s5PFFLD
U491pjbjCQHlI9j3hhXRu9024ULVXqPSeFvhdelg3PJxcmiH7O9C8MTRLjwkaNEa
u8oNLHYk9F8nR2emqhCM5cTbL4QiUVKAQ/cCxGOGq2FPEeIc7P7AZiwD96UZ+HGd
L9JIo5pMAyfsnRMCydsnk1DV02qXNhKBENb5lIAckbwgu4TauT+EfSQh4/2qhmIT
F8fXAekIGvuUy4KCsTiQ4dWzbOiTHtmDruCt/TNUCJE6W4sZ7e158pZW9upNJQ03
f0R3babeGAbFN8JirBXUuglIErY9E7mgLBCeY76bgQIDAQABo20wazAdBgNVHQ4E
FgQUfsKht2p1ABR67JOljbQfN5mZTHIwHwYDVR0jBBgwFoAUfsKht2p1ABR67JOl
jbQfN5mZTHIwDwYDVR0TAQH/BAUwAwEB/zAYBgNVHREEETAPgg0qLmV4YW1wbGUu
Y29tMA0GCSqGSIb3DQEBCwUAA4ICAQCOwxhoZp/+AHrSz2J0FuQP+l125kelKCgV
iPZy5mekvWASINK/zDDtYTaTlijxLiTAW0WQbMspxzlmMwbuoV4uMaOwzaz8dBZG
x0ipNrvJ46XljTdzfK3wRefU7EbB1IDfCxB8738/pjuftWEF4ULItHoO27eaY2/n
H77qAalragqy5lULUzLqH6Qv4Q9C7jqAtqhtNw17GoFuLBZ8ApU0pLMFidL7hx8g
SWfMsKAhauHh5903z/J+2bFASK6dVREsNR90bcpFCZJs1Yd9FtsAXVcgodp0pfLj
u1c3wk4Ms+L7+f7GtzgiKT9D1/BspisZpFHpab1WmSyT867WtVGvTUI6vmWMF1d7
lqUW6ztqYn8wLHkCIrBo32FHkco5oJ/XiV9kS4Q/p6xL1agOijR7Cc17l4p44lob
8o8XhScieU62S+0ePEI/CSLCaEj6o0GTGL/8TOhwRB2QfPr9fp1L1WUhKLX3Co3j
KAnQI67s4YzYT//fx7g6RcR2wJGc5Z+B4wbKU0hXaJPUhmk59g5Z4ZQgZqumABAk
O8sntanM8CRCFMg9mM9NvKm1iBbuffOQwpajUww8ygtXs8smGB9RWUDTie/QTr5B
k8cth3PI/jxQu4ku444ZsF6CXeGuMzYCyr2zSbeDCpCy7f+3Bs7DCeeWSW2t0E6R
H+s6TIXBIg==
-----END CERTIFICATE-----`
	CaPrivateKeyA = `-----BEGIN PRIVATE KEY-----
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
ySmiC7X8sk2PwupmufCDbSExoeLHzDDm
-----END PRIVATE KEY-----`
	CaCrlA = `-----BEGIN X509 CRL-----
MIIC2DCBwQIBATANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEPMA0GA1UE
CAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYwFAYDVQQD
DA0qLmV4YW1wbGUuY29tFw0yMzA1MzExMDM2MDdaFw0yMzA2MzAxMDM2MDdaMBQw
EgIBABcNMjMwNTMxMTAzNTE0WqAjMCEwHwYDVR0jBBgwFoAUfsKht2p1ABR67JOl
jbQfN5mZTHIwDQYJKoZIhvcNAQELBQADggIBAJFO39bfodJC/VxMhut84fYJUxLX
04493WYy5HN1cYN5YptMzDyDL0nIb+uRVQT3Or66g7Ap6saxUtk0U9gjU47k3Lky
m4ihPdCpSHEYfnkh4f2QwfdQsZZaPaYelPPR0i+YKOquywVINbdLDzeuazxtQDBm
08HrlyQ0H5AVq/yq1t2VyKQ/umspj+JfZFasCM9siFfF4GWpt65Z8/yWHzTQFbRL
F8TzGiMZU3/TFw5n3t1lBY2CdEAhSn6+Q7edebTXkEJzgNX67edRi1ieFlBFZ2E7
Oe5Uivgv8G0Izh2tjnxG+FZCAfncgXVuAleBhMKR5So8qf66x/LRCi3VDkHJl2lC
gXYXJRzMVM8Ylhun1HgNv335vURou0fYiM2ytinE1xyx0pCmWW/TZZKDSXtUjnLN
cxye0QaCIASA1ewrkuKhu4xSUOemvn+ClZWaLugKc8/7qbHh3xcJBpxk8zG91r+H
eRwkg3U+XIQOMiy5NTMZj9lUl+FDeWVQvJC9YgK5jQi1ncgF1/gRHRE5uITXGh2R
LdREf3vMgvqlOin08PebxAVUW9WgNhJX/q7EJLEpKhI3NWoGTAmFuBZaArSyx9Hs
ytQsWN9RKoo+A3fO6zljImn9yKI+jwgZW4r8I6SAKE7Ta45mHCxp3jKariL2t/P1
Bv8pBYFjCzRXzXkB
-----END X509 CRL-----`
	DummyCaCrlA = `-----BEGIN X509 CRL-----
MIIC2DCBwQIBATANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEPMA0GA1UE
CAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYwFAYDVQQD
DA0qLmV4YW1wbGUuY29tFw0yMzA2MDEwNzQ2MDdaFw0yMzA3MDEwNzQ2MDdaMBQw
EgIBARcNMjMwNjAxMDc0NTU0WqAjMCEwHwYDVR0jBBgwFoAUfsKht2p1ABR67JOl
jbQfN5mZTHIwDQYJKoZIhvcNAQELBQADggIBADK9sLMx+JxO5RUoeSQMgb1hEE6L
1rk40TbbLO0xYsxEWZYUnFPQrsjiCrwPe3ba9XM8LbLNJnKU6ULmD0quUouwX7eA
Dv2gcX5oMIDRONVlPyKXQ3Yi0WMAwpC+suarUkMRRjKQ3CMK73bqF1/MX9/ZKSJm
jzui7bDTb+IkCzuguCwzdECj4czyqp1WMTZdKZVzK0c7f2gaxfpwkbq1UQ3+zXvk
KdYE8TiN7jkrMT0UtdFYRTPEywCHrof/bUKMBGDkDNyCuXS3HjAtqXY6fdb96qqd
mk5UTmj2cN7kODJiCRUdtrY/hwo/fJsd7tDL2CrINmz1T2yBmctZQsJE544dvJYN
GNxxjY+Nr1zNhRC61Qmcqc7we+/oXy2YCtTao0kI7UDyQueXRiN00FBipSTcqLha
8r42NiyNiH110Qdp5b+9lMAOjzR/DQ8hPoD0jIsI899AL8wZgxxaiPaLEGOK186a
1GE/Nxkf8tis+cfUY/kIbbdTC3SAe1nsF3AnU5l+zUbB34BLVjoDSnMmnylEc1de
3a8GqyGQSxw8/CRt9OIQeu8EYsQ4n8Mwkq6dZ8hXR34hRui/1SRjybbFYGNlAhUz
/d/7YlLMoEBJyGqs1/o50kH45Tu+n/UYGzfS48ApB9Ry26D4ARfprV4CB5ol87JC
jq3qUJ6H45YSZ7hv
-----END X509 CRL-----`
	// Server certificate, private key and CA certificate
	TLSServerCertB = `-----BEGIN CERTIFICATE-----
MIIEOjCCAiKgAwIBAgIBADANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEP
MA0GA1UECAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYw
FAYDVQQDDA0qLmV4YW1wbGUuY29tMB4XDTIwMDgxMjIwMTIzNFoXDTMwMDgxMDIw
MTIzNFowGDEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAJpqQuixQqkI31CE0o1CWv7aP1/umZhdy3lOnrRye2pj
CyGZJHLFcawioOOLAyRA1kA1m4GFWZLDd5b1YBua/6MRZzhh+JNtegpcE9rRBEnq
Xn/PRuM8ji01aeWnvWkm8bppHCGic6kfPSLVUVQ2BQ9zy9GbIlpG1swtlsS5k8lF
N0zNadb6cvnyngwHE5KQ872SAa6+GN33WUbgAqkqbGErdRevrE/lX0+3aMGWm9rR
HTwkUVmfPS5NvIW6UL+Uw5CZR6qTZIYrSv5+vsVu1TI39Olr7ZDCnEVK6eWQ/NOz
ScucwX9HYz4PhHepefVZ60G8nPR5Mf/YNf3tfbaUQEkCAwEAAaNTMFEwCQYDVR0T
BAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMB
MBgGA1UdEQQRMA+CDSouZXhhbXBsZS5jb20wDQYJKoZIhvcNAQELBQADggIBAETJ
T+1wOaqQwPoH4ZNYDsoa/5W+Dh7ykf1Rss0Cgu5SZLvOVJdofhsHI2nX+y2n6iMu
bHuKxxDrJ3BXvvT4ZOojmBge2FLuPfmMXEwv5e1bCa2PIlOS4qD3VkckE+gLW01j
VNiPRf7kB+AVeWVsdq0ZWtsdtb1CUyNWDmxun81CBrPM7dDgSegRnAJ+lTKKLjcE
/lAtDh20Nr8XHukHpOB9LXqXfaRM4KGHqVRJrHoqTWYVJOqd2/0cKutVd15+NqAI
sj3fFgaNLIQQNh1zxhHQy2OnmoAG77jfCjvWuCq8tH/ym28lKuOzn1qod64jNlT7
tslxWjRRyfv7OnDLcCqRyMhNVeLDnHvikx7gPf9uyt/dAH+OQBAU91MfZgnlR72G
fOHeNm3OLYt3UryUd/dpn0NeBLBvOSQb+KL56dzPAetkcfXJEz8WkCPhYGJu06lu
/yYDA1I5O/GOXbtLhnW7tH23Dfu1/KOFngexX/uQptVfu30KlJR0Yh0jPnt1wWdX
LTf19ygwWeJEhARwy+oqxQgfZkmKwQ345JNzgcW09rSS0FnWqWoDSLgI7bQBUFPL
RnnmcV2ARtBYac+05SCbgU82FEsDJmV8wI0ZTm/Snv9iIyJTgSaLuVeC1u6mDtbY
HJOuDwkE1JkfYmvjit+N9puSi53EmguJiy/aGDTA
-----END CERTIFICATE-----`
	TLSServerKeyB = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAmmpC6LFCqQjfUITSjUJa/to/X+6ZmF3LeU6etHJ7amMLIZkk
csVxrCKg44sDJEDWQDWbgYVZksN3lvVgG5r/oxFnOGH4k216ClwT2tEESepef89G
4zyOLTVp5ae9aSbxumkcIaJzqR89ItVRVDYFD3PL0ZsiWkbWzC2WxLmTyUU3TM1p
1vpy+fKeDAcTkpDzvZIBrr4Y3fdZRuACqSpsYSt1F6+sT+VfT7dowZab2tEdPCRR
WZ89Lk28hbpQv5TDkJlHqpNkhitK/n6+xW7VMjf06WvtkMKcRUrp5ZD807NJy5zB
f0djPg+Ed6l59VnrQbyc9Hkx/9g1/e19tpRASQIDAQABAoIBAQCTYv2+UQOJiFhu
6Hh4MG/edb6r+HhOwKgof2xoXCWEpEjJpvztN5yKavsadWNoRSK7Yu+tLvUSatRJ
GKjFcKq0oTXe0VSMk6A7JywkXgEz9CqC3/uPhgtHB4aul+7o4S01eoJJcF/pe1HF
X8sjD/TAjQxsYyDk/lyjwpf36hIvtw+eeKLbLi+H5iS56FXwjC7CX/7ISefL12XV
YARDhucTdvCncV83Wm4IR9oCEMXsBcd6GltEp0HtFN2vawonBBBPUlxtFHOX0DLV
gywuCveHLFgXlqqpRA27XXyLEUBrMA30POwgrQ41FSomogzT025EpCl7S05eyNPk
HYLtY/4JAoGBAMydt05Cy5k75MDWCnUK2bzd6e+t8Kr2tLQtBSXORup8bfBP/s2o
N2WYCauLQV6Omo1j45doBeAInFyz4tC/4fyhXxFBVOJXezXNAXh1cJfCk/fhWxQC
o8dGf9hJofSr33VLfEmvpvDi5/bSe6+VDknzQWRc/UE64uhkOwHyXNifAoGBAMEx
PfaAekWR/vs5jRNf4ydkyALose13bL3ImHfSeEYw7wl499jEXoTI4kTP8cro+vP+
+fZHZkeDWYnnynJg/eBmeqMb93gB5pd9iRY2fgylKY+uX00kmwjCV7HE/5NhINIu
1rZN+xdWz5W3Zc91oepMqZrDlNDdOLxpih6LbPYXAoGBAMW8IFmz+RczJyQndGGZ
Q8Pd3GWv/STqVsTLwO7BAg02g/O5Cq7pwece3zF02I4tVzQN9PHrJ8pR5/E6MZWz
6Qr/U2TUcths2/epDqO/MjimY+InWKHkzysTeRamSamtsruiYBDBe53MWYhP7hFH
BVD3kBkSN/eilYco3VFLSuhXAoGAKaHd84AOWW2z4Bmv4Cv6vKOIQrA4n67rvR48
VR4DE0U6TfVGm+z9XWoY8LFLLun0Ip6g6UTLsr+IjSJpSyrBqxkKdpnxQ2hNc/n2
j9XgZgM1qKPwH6Sy0DlpQoLsfKsisoirdo/pe0cW5vGlvAZihSZOKm7ZZAU4U8n+
nMR5D7MCgYEAxnHXmI+b1QHwKqQUb6yeMtakY7CrOrWeJ44TL3XPbZmJ1eM5ibob
2wTTyvk0YfQODtrtKhrP3U/KU1zds/EBiw9Ay5KKxtjlVY9kZqmexNXX/wbPW4kj
0IJJ+Jky2HdfH+kWGUJEKrhYCWQCPLb9rCBKyFSO61Z87Zz3n7ApqkM=
-----END RSA PRIVATE KEY-----`
	TLSClientCertB = `-----BEGIN CERTIFICATE-----
MIIEOjCCAiKgAwIBAgIBADANBgkqhkiG9w0BAQsFADBUMQswCQYDVQQGEwJVUzEP
MA0GA1UECAwGRGVuaWFsMQ4wDAYDVQQHDAVFdGhlcjEMMAoGA1UECgwDRGlzMRYw
FAYDVQQDDA0qLmV4YW1wbGUuY29tMB4XDTIwMDgxMjIwMTIzNFoXDTMwMDgxMDIw
MTIzNFowGDEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAK/gBQdXoWiIPRNIBlKninGXkU4iy62JEQ2bB83b5I1y
rwhUMdXMg9QPElnvNZrPx7KP3X6U7ZVXna4BP/7wGHmBBkYDfZGrnnAyFgbe7NSD
zP/oBwE42HWVnzG02Q/whQI3KbIk/ZphuXOPSoRhtzHOXU8Ew27TccZAIuei0rMl
nsP+R1UfR2QZifosg8Yy5dkOhMvvJgavfdxhLz9D2pAMiSLIyXOc2LjAHVqQp/Ai
eutFXTGZt2WwbRUo810sbSEYKugFZQXB0L6qJcYFdaVvD3XrQ+BUrqKN0b4qtSHt
fbMQPwNuGMqbf6vKu2ek1Hy7eq2zflUdkVf8lIC1R00CAwEAAaNTMFEwCQYDVR0T
BAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMB
MBgGA1UdEQQRMA+CDSouZXhhbXBsZS5jb20wDQYJKoZIhvcNAQELBQADggIBAH5o
lKH+prw6S5Yxy2B71ibD7AHrHwhAdyrqmvRx5jnFxszfao6joH8Wr/XA69OMc8YJ
yj6cSpV9lBhpezvg0qwB2nPXHurmSRqMilZKUZOo6olbPPYaT7joShp/xPpx0W+/
zKgshEn5InfWc4KaiwB6PbEUrCCoinGNvjd5DRkxGvJppjUGXffeyhJHtJDBfj8q
rf9dsKGzuLhRmclXbQy1i776MPisLAWpOONfi82z0sKkKuMX28FI2Mw11b5t3OQ4
iD7SejLAreqpbZ1f0PFduAm9PTgcqgT4DNR/mNbRChT4NPXHPdESLgoDM/zn9FcA
oLofL9Blz32zRhc6W8f2o3Zx5UGtvGmw81gUurAhmjORDIeFZORXVaNcHUIi0KrI
uJi6IgohkKSJzI1tKVVYykzCvy2+S6N2++fKOTrfatVYUXkWtQcXppeY3v1oj4Im
dFIII+vHpmNwg+cODcS4ljBEWZ2hh003YT1PV1x38U+vAgGwKRxE8XweTQ51aQJQ
lndGcoVv2z+9oEXVjuYGW4PpwLSbjfxsUFGsISGgvzGq7zjPqPqrkXcXE4sCIpoY
7iTk+R1VX5Fg/3Wcx4lLjnejuAQRLCX4N3LZIbteTMndIPtBXio+4Aaj4K8FV0VZ
zdtzaucVaYSWZ58BET/277ie+BCkTPOU4NNZikTM
-----END CERTIFICATE-----`
	TLSClientKeyB = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAr+AFB1ehaIg9E0gGUqeKcZeRTiLLrYkRDZsHzdvkjXKvCFQx
1cyD1A8SWe81ms/Hso/dfpTtlVedrgE//vAYeYEGRgN9kauecDIWBt7s1IPM/+gH
ATjYdZWfMbTZD/CFAjcpsiT9mmG5c49KhGG3Mc5dTwTDbtNxxkAi56LSsyWew/5H
VR9HZBmJ+iyDxjLl2Q6Ey+8mBq993GEvP0PakAyJIsjJc5zYuMAdWpCn8CJ660Vd
MZm3ZbBtFSjzXSxtIRgq6AVlBcHQvqolxgV1pW8PdetD4FSuoo3Rviq1Ie19sxA/
A24Yypt/q8q7Z6TUfLt6rbN+VR2RV/yUgLVHTQIDAQABAoIBAE5tnrB0gnTj9CBq
CO/I6LLY/SIT4DIFooYnsSY6vcPZ9XXIXeTSa6MgCXXzdzsPFyVohERTU1M6nAUh
Z5FfLfrY75w/KDrShjfQDgSIWqI2GC+yH5WK7PQagcAfvoAeKKIhSUW4uyKiG5oN
txQrow7h0LhRDnmYbu8SYL5c3IqoKTPo7kCDx4y8x55NWt9rLO8NNkWeP3TFAnQr
sJ5cL2mFuMZq2aOCjjsNmYbzyNYPPTAGSBC5h+rHjBDRVJCHh4+MsCzslsiPrjCf
jtwpNjiU7oleVjCo++W/bqeLBYZEAK1E3mTjq7Uj6uuOJNKZJUgUWCpcdILWUxge
OtHhrwECgYEA5hHzkE571V8D2DespFuyR29DIqPKYr+vw1Gf00vG4fSBw/qoVBET
k2w5KZeUHVpd/FAmLGS3jmcFdCEAVOzymvzdtzD9YoXRPJ9MhkCldVbV4uypN2CS
ujBXhXAkIx0zWTYz7maZCO58I6LSKfIZH7vPKdbuoC2CyCJ2V0f0p0ECgYEAw7Jq
13V7gnvB2lemXGi5SlzYsfnaa1Jh4xa3HvvUPIqA4L+TQL7DvBU2LLQBOscH8QWX
SMZ+ESp4ffuiy3bNszXpbYr2tAwZMvwBXM9m5oMBfcUxRO23Ne94yfG8syvO1VDd
cGdGfmzsH7rKb8xWG2NdZn7Mp/3L4CEPHhkHiQ0CgYEAtWHpuBdoIawrB8e84Ec6
on7hWunuoTOmikJL0vWm8nUl6TAwjTZ5THzCBDxwlUOtXUKDFfTBkg8+iXOLI2k7
S12YHtkMqyP89eJRHnbQ+1Vegu0562LDyu5Z0eW7V8Opu+ezwXw5ooUCccXoExG/
JPmYNdAxiS7YMJia2+zqwQECgYBMvRL+aDhFg4gRuZsn030NNvlWWRaYTRbLeOXY
BYJFNz6wYw84LVYN0/eJchClXtSzY7DePc/DpXbYWCuPQ4gIBVUCuqJhGflrr7xN
C5tbOvYDvqbCU8ErKjugXOXVEJGrvmkHIchX3Rh+nL2zN1pwpPdEMObJ3VDsE6qc
782vpQKBgGWQ7RrkMOkeVG1MmTDSRvsM7VNVUGVoAeq5N5Z1QpqapOPwYaR+slGg
k+tLdh0VYYLalafzJMBhc9IMOz6H9bCP6MM9HfppfYIeqGzJksVSQ4OBqhH08Wva
perqL7LjUzSvUhNMsIKiwFCgKEnVpXC8crV+8jwxDliOajhEt1NB
-----END RSA PRIVATE KEY-----`
	CaCertB = `-----BEGIN CERTIFICATE-----
MIIFozCCA4ugAwIBAgIUbjFaOdwRFHy+Jqin3de6AsPlKVMwDQYJKoZIhvcNAQEL
BQAwVDELMAkGA1UEBhMCVVMxDzANBgNVBAgMBkRlbmlhbDEOMAwGA1UEBwwFRXRo
ZXIxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTAeFw0yMDA4
MTIyMDEyMzRaFw0zMDA4MTAyMDEyMzRaMFQxCzAJBgNVBAYTAlVTMQ8wDQYDVQQI
DAZEZW5pYWwxDjAMBgNVBAcMBUV0aGVyMQwwCgYDVQQKDANEaXMxFjAUBgNVBAMM
DSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDC
se+/CHI6gNhoj2pW/iawYaZT+FavNBwKq9y8+jQKp8ZSQvF0+44KWWTu258I/xxK
agJpXITOieCuLOHxsw2iMCgHz8US3cuAm3trvVG5se4ksEzBKXEqUtr4x7j25zUr
6r3JNQ1S51zIthc85ZchZFkS82T7uEgpYF5mDEbP7jlfZ/KVL/oquH1IJOjCufIU
leyAazR5qDj9fdveP4TxJzZkhLEaw8qnJYNXvgVkMHkEDyArFBcrVoGFKl5i9HKH
kga7CwpQ/PapQ7UUmiBQ7D9nuk96YCVWyVZEdnQm7gsgTmw9ZwdbDDBW+nnU9nx9
GoqmFPFvA3LwnQLHYf8ZXNSYzrFGypE4v5WSg0j0IAu7Msw7RtBrcZ2ygkROAfL2
kGw60ufCeki0bG2X/YCLvKcL/ldtUtEoKC4O9Wfl/Rx+tHQTJiIokAH6vPLV5odH
zm3/ezl1Q3WK4Na/Va4g+CG2D6hqVTLwgfmw9ju531+KQyscY7k0b1Sjjz1KlE+b
5cN8QEereE/ce0fQo+5DcZTiYJIKRaVC3yYWUVZaGB3Zms34Ht1dKLeyI5JZ4YCx
fKakIr1TGGR3ep+ErCOMh8WuCTFWA+DUYyuk7s8P7tO+JEL3LSOt9iWeKKLLKmI/
srjYZGbxdPsiAmXdE/mL3v2vz8JAZFG+JHSRkffwQQIDAQABo20wazAdBgNVHQ4E
FgQUGdp1TXyn19hVavwCfZYts3TYKiwwHwYDVR0jBBgwFoAUGdp1TXyn19hVavwC
fZYts3TYKiwwDwYDVR0TAQH/BAUwAwEB/zAYBgNVHREEETAPgg0qLmV4YW1wbGUu
Y29tMA0GCSqGSIb3DQEBCwUAA4ICAQDB0PNp3p3Byfm36R70EA8vlauQhAfJMP3M
0pWYQ/JWU9ngS3AYt4ebdPlTH64JLqTfpm8Ye/BN3UhDKtFgLAnKFw6EGKFkbi17
NUPGIbp1+6djasIEMY1C6whcsaStn4mI1tsZwi0FZGphOQb0lKLL0SJ1PdMp3hTf
P86yNmvN/dp4wcZc91sIe08JPHC//oXzB7dwaa0YMR01dadh0D9fL4ZmQ1KgVIb/
di+q/U1TDdRXtZK94zYem1SbGTyowzZRbCtTV16niwGxl/42w5b4iLArdDUArm/C
1wqXS318WGh7qOY5ndL3TSqhJe6B+39NRglaRT6yAR024BWdM1bnL5eG4jeJu1Qy
gEE0StJ8e2hgDMiEd3t96xWdOKzOZwK7dhJx//DaXmcTxYe0F5rs8uTIbvZZhM+L
KIcEXzuqKOoRXoDiP7iwEF841pW1tg5s7dmA9G62Uj+40IlUe7PEjhZpnvOOpli1
xfJ07NmTPWPWdlHZir9+TKCeLOwvNW7WANvtbZi45mX/mhCKGeo/2WBiu2fFkW+B
SSHx/mIjuYeAbQAj8YcuplECTg7RPHZvFACb121J/oCuMGPKapRMVpmtk276cMGo
dQ5SwyXq0bFyf/Phguy5+kXS+jAgCuzsdIes9FVUYU+pToyq0DkLVdRkqM4f8Pfh
Kj0nxYj7eQ==
-----END CERTIFICATE-----`
	CaPrivateKeyB = `-----BEGIN PRIVATE KEY-----
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
aU0XUSD0tBPTHJUwLIe85kCzg/3Jius=
-----END PRIVATE KEY-----`
)

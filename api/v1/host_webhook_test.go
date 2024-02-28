/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */
package v1

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("system_webhook functions", func() {

	Describe("validateMatchBMInfo function is tested", func() {
		Context("", func() {
			It("", func() {
				bmAddr := "192.13.24.39"
				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							BoardManagement: &MatchBMInfo{
								Address: &bmAddr,
							},
						},
					},
				}
				err := r.validateMatchBMInfo()
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							BoardManagement: &MatchBMInfo{},
						},
					},
				}
				msg := errors.New("board management address must be supplied in match criteria")
				err := r.validateMatchBMInfo()
				Expect(err).To(Equal(msg))
			})
		})
	})
	Describe("validateMatchDMIInfo function is tested", func() {
		Context("", func() {
			It("", func() {
				serialNo := "12345"
				astTag := "90"
				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							DMI: &MatchDMIInfo{
								SerialNumber: &serialNo,
								AssetTag:     &astTag,
							},
						},
					},
				}
				err := r.validateMatchDMIInfo()
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				astTag := "90"
				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							DMI: &MatchDMIInfo{
								AssetTag: &astTag,
							},
						},
					},
				}
				msg := errors.New("DMI Serial Number or Asset Tag must be supplied in match criteria")
				err := r.validateMatchDMIInfo()
				Expect(err).To(Equal(msg))
			})
		})
	})
	Describe("validateMatchInfo function is tested", func() {
		Context("", func() {
			It("", func() {
				bmAddr := "192.13.24.39"
				serialNo := "12345"
				astTag := "90"
				bootMac := "01:02:03:04:05:06"

				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							BoardManagement: &MatchBMInfo{
								Address: &bmAddr,
							},
							DMI: &MatchDMIInfo{
								SerialNumber: &serialNo,
								AssetTag:     &astTag,
							},
							BootMAC: &bootMac,
						},
					},
				}
				err := r.validateMatchInfo()
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{},
					},
				}
				msg := errors.New("host must be configured with at least 1 match criteria")
				err := r.validateMatchInfo()
				Expect(err).To(Equal(msg))
			})
		})
	})
	Describe("validateHost function is tested", func() {
		Context("", func() {
			It("", func() {
				bmAddr := "192.13.24.39"
				serialNo := "12345"
				astTag := "90"
				bootMac := "01:02:03:04:05:06"

				r := &Host{
					Spec: HostSpec{
						Match: &MatchInfo{
							BoardManagement: &MatchBMInfo{
								Address: &bmAddr,
							},
							DMI: &MatchDMIInfo{
								SerialNumber: &serialNo,
								AssetTag:     &astTag,
							},
							BootMAC: &bootMac,
						},
					},
				}
				err := r.validateHost()
				Expect(err).To(BeNil())
			})
		})
	})
})

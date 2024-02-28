/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */
package v1

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("system_webhook functions", func() {

	Describe("validateBackendServices function is tested", func() {
		Context("", func() {
			It("", func() {
				backendType := "ceph"
				services := []string{"cinder", "nova"}
				err := validateBackendServices(backendType, services)

				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				backendType := "lvm"
				services := []string{"cinder", "nova"}
				err := validateBackendServices(backendType, services)

				msg := errors.New("nova service not allowed with lvm backend.")
				Expect(err).To(Equal(msg))

			})
		})
	})
	Describe("validateBackendAttributes function is tested", func() {
		Context("", func() {
			It("", func() {
				prtSize := 20
				repFac := 2
				backend := StorageBackend{
					PartitionSize:     &prtSize,
					ReplicationFactor: &repFac,
					Type:              ceph,
				}

				err := validateBackendAttributes(backend)
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				prtSize := 20
				repFac := 2
				backend := StorageBackend{
					PartitionSize:     &prtSize,
					ReplicationFactor: &repFac,
					Type:              file,
				}

				err := validateBackendAttributes(backend)
				msg := errors.New("partitionSize and ReplicationFactor only permitted with ceph backend")
				Expect(err).To(Equal(msg))
			})
		})
	})
	Describe("validateStorageBackends function is tested", func() {
		Context("", func() {
			It("", func() {
				prtSize := 20
				repFac := 2

				obj := &System{
					Spec: SystemSpec{
						Storage: &SystemStorageInfo{
							Backends: &StorageBackendList{
								{
									PartitionSize:     &prtSize,
									ReplicationFactor: &repFac,
									Type:              ceph,
									Services:          []string{"cinder", "nova"},
								},
							},
						},
					},
				}
				err := validateStorageBackends(obj)
				Expect(err).To(BeNil())
			})
		})
		Context("", func() {
			It("", func() {
				prtSize := 20
				repFac := 2

				obj := &System{
					Spec: SystemSpec{
						Storage: &SystemStorageInfo{
							Backends: &StorageBackendList{
								{
									PartitionSize:     &prtSize,
									ReplicationFactor: &repFac,
									Type:              ceph,
									Services:          []string{"cinder", "nova"},
								},
								{
									PartitionSize:     &prtSize,
									ReplicationFactor: &repFac,
									Type:              ceph,
									Services:          []string{"swift"},
								},
							},
						},
					},
				}
				err := validateStorageBackends(obj)
				msg := errors.New("backend services may only be specified once.")
				Expect(err).To(Equal(msg))
			})
		})
	})
	Describe("validateStorage function is tested", func() {
		Context("", func() {
			It("", func() {
				prtSize := 20
				repFac := 2
				obj := &System{
					Spec: SystemSpec{
						Storage: &SystemStorageInfo{
							Backends: &StorageBackendList{
								{
									PartitionSize:     &prtSize,
									ReplicationFactor: &repFac,
									Type:              ceph,
									Services:          []string{"cinder", "nova"},
								},
							},
						},
					},
				}

				err := validateStorage(obj)
				Expect(err).To(BeNil())
			})
		})
	})
})

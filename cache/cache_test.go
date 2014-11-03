package cache_test

import (
	"regexp"
	"runtime"

	"github.com/DamnWidget/VenGO/cache"
)

var _ = Describe("Cache", func() {

	Describe("ExpandUser returns valid path depending on platform", func() {
		var re *regexp.Regexp
		Context("BSD and GNU/Linux", func() {
			if runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
				re = regexp.MustCompile("/home/([a-z0-9]+)/VenGO")
			}
		})

		Context("OS X", func() {
			if runtime.GOOS == "darwin" {
				re = regexp.MustCompile("/Users/([a-z0-9]+)/VenGO")
			}
		})

		Context("Windows", func() {
			if runtime.GOOS == "windows" {
				re = regexp.MustCompile("\\Users\\([a-z0-9]+)\\VenGO")
			}
		})

		It("Should Match", func() {
			Expect(re.MatchString(cache.ExpandUser("~/VenGO"))).To(BeTrue())
		})
	})

	Describe("CacheDirectory is a valid directory on each platform", func() {
		var re *regexp.Regexp
		Context("BSD and GNU/Linux", func() {
			if runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
				re = regexp.MustCompile("/home/([a-z0-9]+)/.cache/VenGO")
			}
		})

		Context("OS X", func() {
			if runtime.GOOS == "darwin" {
				re = regexp.MustCompile(
					"/Users/([a-zA-Z0-9]+)/Library/Caches/VenGO")
			}
		})

		Context("Windows", func() {
			if runtime.GOOS == "windows" {
				re = regexp.MustCompile(
					"\\Users\\([a-zA-Z0-9]+)\\AppData\\VenGO")
			}
		})

		It("Should Match", func() {
			Expect(re.MatchString(cache.CacheDirectory())).To(BeTrue())
		})
	})

})
// Package testimonials is the Testimonials content pack for togo: Testimonial + rotating Quote resources.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package testimonials

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("testimonials", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}

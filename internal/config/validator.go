package config

import (
	"fmt"
	"net/url"
	"strings"
)

// ValidationError describes a single config field violation.
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("  config.%s: %s", e.Field, e.Message)
}

// ValidationErrors is an ordered list of config violations.
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	msgs := make([]string, len(e))
	for i, err := range e {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "\n")
}

func (e ValidationErrors) HasErrors() bool {
	return len(e) > 0
}

// Validate checks that all required fields are present and well-formed.
func Validate(cfg *Config) ValidationErrors {
	var errs ValidationErrors

	if strings.TrimSpace(cfg.Name) == "" {
		errs = append(errs, ValidationError{"name", "required"})
	}
	if strings.TrimSpace(cfg.Username) == "" {
		errs = append(errs, ValidationError{"username", "required"})
	}
	if strings.TrimSpace(cfg.Role) == "" {
		errs = append(errs, ValidationError{"role", "required"})
	}
	if cfg.Website != "" {
		if _, err := url.ParseRequestURI(cfg.Website); err != nil {
			errs = append(errs, ValidationError{"website", "must be a valid URL"})
		}
	}

	for i, p := range cfg.Projects {
		if strings.TrimSpace(p.Name) == "" {
			errs = append(errs, ValidationError{
				fmt.Sprintf("projects[%d].name", i), "required",
			})
		}
	}

	for i, e := range cfg.Timeline {
		if strings.TrimSpace(e.Year) == "" {
			errs = append(errs, ValidationError{
				fmt.Sprintf("timeline[%d].year", i), "required",
			})
		}
		if strings.TrimSpace(e.Title) == "" {
			errs = append(errs, ValidationError{
				fmt.Sprintf("timeline[%d].title", i), "required",
			})
		}
	}

	return errs
}

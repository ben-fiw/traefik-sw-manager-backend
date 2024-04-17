package requests

import "encoding/xml"

type ServiceRegistrationRequest struct {
	XmlName        xml.Name `json:"-" xml:"request" yaml:"-"`
	ServiceName    string   `json:"serviceName" xml:"serviceName" yaml:"serviceName" validate:"required,min=3,max=100"`
	Host           string   `json:"host" xml:"host" yaml:"host" validate:"required,hostname|ip"`
	Port           int      `json:"port" xml:"port" yaml:"port" validate:"required,number,min=1,max=65535"`
	ResourceRating int      `json:"resourceRating" xml:"resourceRating" yaml:"resourceRating" validate:"omitempty,min=1,max=100"`
}

type ServiceRequestRequest struct {
	XmlName     xml.Name `json:"-" xml:"request" yaml:"-"`
	ServiceType string   `json:"serviceType" xml:"serviceType" yaml:"serviceType" validate:"required,min=3,max=100"`
}

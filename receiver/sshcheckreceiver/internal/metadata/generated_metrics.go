// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for sshcheckreceiver metrics.
type MetricsSettings struct {
	SshcheckDuration     MetricSettings `mapstructure:"sshcheck.duration"`
	SshcheckError        MetricSettings `mapstructure:"sshcheck.error"`
	SshcheckSftpDuration MetricSettings `mapstructure:"sshcheck.sftp_duration"`
	SshcheckSftpError    MetricSettings `mapstructure:"sshcheck.sftp_error"`
	SshcheckSftpStatus   MetricSettings `mapstructure:"sshcheck.sftp_status"`
	SshcheckStatus       MetricSettings `mapstructure:"sshcheck.status"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SshcheckDuration: MetricSettings{
			Enabled: true,
		},
		SshcheckError: MetricSettings{
			Enabled: true,
		},
		SshcheckSftpDuration: MetricSettings{
			Enabled: false,
		},
		SshcheckSftpError: MetricSettings{
			Enabled: false,
		},
		SshcheckSftpStatus: MetricSettings{
			Enabled: false,
		},
		SshcheckStatus: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular metric.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledProvidedByUser bool
}

func (ras *ResourceAttributeSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ras, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ras.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesSettings provides settings for sshcheckreceiver metrics.
type ResourceAttributesSettings struct {
	SSHEndpoint ResourceAttributeSettings `mapstructure:"ssh.endpoint"`
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{
		SSHEndpoint: ResourceAttributeSettings{
			Enabled: false,
		},
	}
}

type metricSshcheckDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.duration metric with initial data.
func (m *metricSshcheckDuration) init() {
	m.data.SetName("sshcheck.duration")
	m.data.SetDescription("Measures the duration of SSH connection.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricSshcheckDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckDuration) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckDuration(settings MetricSettings) metricSshcheckDuration {
	m := metricSshcheckDuration{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckError struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.error metric with initial data.
func (m *metricSshcheckError) init() {
	m.data.SetName("sshcheck.error")
	m.data.SetDescription("Records errors occurring during SSH check.")
	m.data.SetUnit("{error}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSshcheckError) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("error.message", errorMessageAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckError) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckError) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckError(settings MetricSettings) metricSshcheckError {
	m := metricSshcheckError{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_duration metric with initial data.
func (m *metricSshcheckSftpDuration) init() {
	m.data.SetName("sshcheck.sftp_duration")
	m.data.SetDescription("Measures SFTP request duration.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricSshcheckSftpDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpDuration) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpDuration(settings MetricSettings) metricSshcheckSftpDuration {
	m := metricSshcheckSftpDuration{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpError struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_error metric with initial data.
func (m *metricSshcheckSftpError) init() {
	m.data.SetName("sshcheck.sftp_error")
	m.data.SetDescription("Records errors occurring during SFTP check.")
	m.data.SetUnit("{error}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSshcheckSftpError) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("error.message", errorMessageAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpError) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpError) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpError(settings MetricSettings) metricSshcheckSftpError {
	m := metricSshcheckSftpError{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_status metric with initial data.
func (m *metricSshcheckSftpStatus) init() {
	m.data.SetName("sshcheck.sftp_status")
	m.data.SetDescription("1 if the SFTP server replied to request, otherwise 0.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSshcheckSftpStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpStatus) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpStatus(settings MetricSettings) metricSshcheckSftpStatus {
	m := metricSshcheckSftpStatus{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.status metric with initial data.
func (m *metricSshcheckStatus) init() {
	m.data.SetName("sshcheck.status")
	m.data.SetDescription("1 if the SSH client successfully connected, otherwise 0.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSshcheckStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckStatus) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckStatus(settings MetricSettings) metricSshcheckStatus {
	m := metricSshcheckStatus{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                  pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity            int                 // maximum observed number of metrics per resource.
	resourceCapacity           int                 // maximum observed number of resource attributes.
	metricsBuffer              pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                  component.BuildInfo // contains version information
	resourceAttributesSettings ResourceAttributesSettings
	metricSshcheckDuration     metricSshcheckDuration
	metricSshcheckError        metricSshcheckError
	metricSshcheckSftpDuration metricSshcheckSftpDuration
	metricSshcheckSftpError    metricSshcheckSftpError
	metricSshcheckSftpStatus   metricSshcheckSftpStatus
	metricSshcheckStatus       metricSshcheckStatus
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

// WithResourceAttributesSettings sets ResourceAttributeSettings on the metrics builder.
func WithResourceAttributesSettings(ras ResourceAttributesSettings) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.resourceAttributesSettings = ras
	}
}

func NewMetricsBuilder(ms MetricsSettings, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                  pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:              pmetric.NewMetrics(),
		buildInfo:                  settings.BuildInfo,
		resourceAttributesSettings: DefaultResourceAttributesSettings(),
		metricSshcheckDuration:     newMetricSshcheckDuration(ms.SshcheckDuration),
		metricSshcheckError:        newMetricSshcheckError(ms.SshcheckError),
		metricSshcheckSftpDuration: newMetricSshcheckSftpDuration(ms.SshcheckSftpDuration),
		metricSshcheckSftpError:    newMetricSshcheckSftpError(ms.SshcheckSftpError),
		metricSshcheckSftpStatus:   newMetricSshcheckSftpStatus(ms.SshcheckSftpStatus),
		metricSshcheckStatus:       newMetricSshcheckStatus(ms.SshcheckStatus),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(ResourceAttributesSettings, pmetric.ResourceMetrics)

// WithSSHEndpoint sets provided value as "ssh.endpoint" attribute for current resource.
func WithSSHEndpoint(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.SSHEndpoint.Enabled {
			rm.Resource().Attributes().PutStr("ssh.endpoint", val)
		}
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/sshcheckreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSshcheckDuration.emit(ils.Metrics())
	mb.metricSshcheckError.emit(ils.Metrics())
	mb.metricSshcheckSftpDuration.emit(ils.Metrics())
	mb.metricSshcheckSftpError.emit(ils.Metrics())
	mb.metricSshcheckSftpStatus.emit(ils.Metrics())
	mb.metricSshcheckStatus.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesSettings, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSshcheckDurationDataPoint adds a data point to sshcheck.duration metric.
func (mb *MetricsBuilder) RecordSshcheckDurationDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckDuration.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckErrorDataPoint adds a data point to sshcheck.error metric.
func (mb *MetricsBuilder) RecordSshcheckErrorDataPoint(ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	mb.metricSshcheckError.recordDataPoint(mb.startTime, ts, val, errorMessageAttributeValue)
}

// RecordSshcheckSftpDurationDataPoint adds a data point to sshcheck.sftp_duration metric.
func (mb *MetricsBuilder) RecordSshcheckSftpDurationDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckSftpDuration.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckSftpErrorDataPoint adds a data point to sshcheck.sftp_error metric.
func (mb *MetricsBuilder) RecordSshcheckSftpErrorDataPoint(ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	mb.metricSshcheckSftpError.recordDataPoint(mb.startTime, ts, val, errorMessageAttributeValue)
}

// RecordSshcheckSftpStatusDataPoint adds a data point to sshcheck.sftp_status metric.
func (mb *MetricsBuilder) RecordSshcheckSftpStatusDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckSftpStatus.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckStatusDataPoint adds a data point to sshcheck.status metric.
func (mb *MetricsBuilder) RecordSshcheckStatusDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckStatus.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

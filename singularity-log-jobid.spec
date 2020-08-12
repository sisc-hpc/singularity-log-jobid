# the minimal version of singularity that we need to run this
# In principle, a plugin only keeps on working within the same major version.
%define singularity_ver 3.6.0

Summary: Log executed commands to syslog along with username and jobid (both slurm and torque are supported)
Name: singularity-log-jobid
Version: 0.1.0
Release: %{singularity_ver}
License: BSD
Group: Applications/System
Source: singularity-log-jobid.sif
BuildArch: x86_64
Requires: singularity >= %{singularity_ver}

%description
Log executed commands to syslog along with username and jobid (both slurm and torque are supported)

%prep

%build

%install
%{__mkdir_p} %{buildroot}%{_libexecdir}/singularity
%{__install} -pm 644 %{SOURCE0} %{buildroot}%{_libexecdir}/singularity

%clean
rm -rf %{buildroot}

%post
singularity plugin install %{_libexecdir}/singularity/singularity-log-jobid.sif

%postun
singularity plugin uninstall github.com/wpoely86/singularity-log-jobid

%files
%defattr(-,root,root,-)
%{_libexecdir}/singularity/singularity-log-jobid.sif

%changelog
* Wed Aug 12 2020 Ward Poelmans <ward.poelmans@vub.be>
- First version

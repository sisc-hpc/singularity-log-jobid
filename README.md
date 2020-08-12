Singularity jobid logging
=========================

This is a simple plugin based on the log-plugin example of singularity that will log the username (instead of UID) 
and the jobid (for both slurm and torque).

Install
-------

The binary sif file is included. To start using, do:

```
singularity plugin install singularity-log-jobid.sif
```

Building
--------

Building the sif file is not for the feint hearted. You need the singularity source and a singularity binary 
built from that exact same source (both version and location need to match).

These instructions should work on CentOS 7 (with EPEL):

```
export rpmbasedir=$PWD/rpmbuild
mkdir -p $rpmbasedir/{BUILD,RPMS,SRPMS,SOURCES}
yumdownloader --source singularity
rpmbuild --define "_topdir $rpmbasedir" --rebuild --noclean singularity-3.6.1-1.el7.src.rpm
rpm -ivh $rpmbasedir/RPMS/x86_64/singularity-3.6.1-1.el7.x86_64.rpm 
singularity plugin compile .
```

You need a recent enough for of git for the cloning of the go packages to work.

RPM
---

There is a spec file included. In principle you can simply run `bdist_rpm.sh` and get a RPM. It's tied to
main version of singularity. The developer promise to keep plugins working within the same major
singularity release (3.6.x).
